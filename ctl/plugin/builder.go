package plugin

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const defaultBackendDir = "backend"
const defaultFrontendDir = "frontend"
const requiredGoVersion = "1.21"
const requiredNodeVersion = "16.0.0"

type Builder struct {
	name        string
	workDir     string
	backendPath string
}

func NewBuilder(workDir string) (*Builder, error) {
	absDir, err := filepath.Abs(workDir)
	if err != nil {
		return nil, err
	}
	return &Builder{
		workDir: absDir,
	}, nil
}

// CheckDependencies 检查依赖
func (b *Builder) CheckDependencies() error {
	// 检查当前环境 darwin/linux/windows
	slog.Info(fmt.Sprintf("os is %s", runtime.GOOS))
	return nil
}

// BuildBackend 构建后端
func (b *Builder) BuildBackend(relativeDir string) error {
	// 执行 go mod download，如果出错则返回错误
	slog.Info(fmt.Sprintf("start build backend project %s", relativeDir))
	cmd := exec.Command("go", "mod", "download")
	cmd.Dir = filepath.Join(b.workDir, relativeDir)
	cmd.Env = append(os.Environ(), "GO111MODULE=on")
	initCmdOutput(cmd)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("exec go mod download failed: %v", err)
	}

	// 提取 relativeDir/etc/etc.yaml 文件, 提取到 workDir/target/etc.yaml
	srcPath := filepath.Join(b.workDir, relativeDir, "etc/etc.yaml")
	dstPath := filepath.Join(b.workDir, "target/etc.yaml")

	err := copyFile(srcPath, dstPath)
	if err != nil {
		return err
	}

	// 使用 go build 编译 relativeDir目录, 编译以后的结果提取到 workDir/target 目录下
	buildDir := filepath.Join(b.workDir, relativeDir)
	outputPath := filepath.Join(b.workDir, "target", "backend.exe")
	cmd = exec.Command("go", "build", "-o", outputPath, ".")
	cmd.Dir = buildDir
	initCmdOutput(cmd)
	if err = cmd.Run(); err != nil {
		return err
	}
	b.backendPath = "backend.exe"
	return nil
}

// BuildFrontend 构建前端
func (b *Builder) BuildFrontend(relativeDir string) error {
	// 执行 npm install && npm run build， 如果出错则返回错误
	slog.Info(fmt.Sprintf("start build frontend project %s", relativeDir))
	cmd := exec.Command("yarn", "install")
	cmd.Dir = filepath.Join(b.workDir, relativeDir)
	initCmdOutput(cmd)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("exec npm install failed: %v", err)
	}

	cmd = exec.Command("yarn", "run", "build")
	cmd.Dir = filepath.Join(b.workDir, relativeDir)
	initCmdOutput(cmd)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("exec npm run build failed: %v", err)
	}

	slog.Info(fmt.Sprintf("build frontend project %s success", relativeDir))

	// 提取 relativeDir/src/api 目录除了 interceptor.ts 以外的所有文件, 提取到 workDir/target/api 目录下
	err := copyFilesExclude(filepath.Join(b.workDir, relativeDir, "src", "api"), filepath.Join(b.workDir, "target", "frontend", "api"), "interceptor.ts")
	if err != nil {
		return fmt.Errorf("copy api files failed: %v", err)
	}

	// 提取 relativeDir/src/router/routes/modules 目录下的第一个文件, 获取其name, 提取为 workDir/target/module.ts
	moduleName, err := getFirstModuleName(filepath.Join(b.workDir, relativeDir, "src", "router", "routes", "modules"))
	if err != nil {
		return fmt.Errorf("copy module file failed: %v", err)
	}

	err = copyFile(filepath.Join(b.workDir, relativeDir, "src", "router", "routes", "modules", moduleName+".ts"), filepath.Join(b.workDir, "target", "frontend", "module.ts"))
	if err != nil {
		return fmt.Errorf("copy module file failed: %v", err)
	}

	// 提取 relativeDir/src/views 目录下文件名为 name 的目录, 将其下的所有文件提取到 workDir/target/views 目录下
	err = copyDir(filepath.Join(b.workDir, relativeDir, "src", "views", moduleName), filepath.Join(b.workDir, "target", "frontend", "views"))
	if err != nil {
		return fmt.Errorf("copy views files failed: %v", err)
	}

	// 提取 assets 下的 images 文件夹, 提取到 workDir/target/assets 目录下
	err = copyDir(filepath.Join(b.workDir, relativeDir, "src", "assets", "images"), filepath.Join(b.workDir, "target", "frontend", "assets", "images"))

	return nil
}

func (b *Builder) BuildExtra() error {
	// 检查根目录下是否存在 plugin.yaml 文件，如果存在则读取，如果不存在则报错
	var info BuildInfo
	var route Route
	if _, err := os.Stat(filepath.Join(b.workDir, "plugin.yaml")); os.IsNotExist(err) {
		return fmt.Errorf("plugin.yaml not found")
	} else {
		content, err := os.ReadFile(filepath.Join(b.workDir, "plugin.yaml"))
		if err != nil {
			return fmt.Errorf("read plugin.yaml file failed: %v", err)
		}
		if err = yaml.Unmarshal(content, &info); err != nil {
			return fmt.Errorf("parse plugin.yaml file failed: %v", err)
		}
		if err = yaml.Unmarshal(content, &route); err != nil {
			return fmt.Errorf("parse plugin.yaml file failed: %v", err)
		}
	}

	b.name = info.Name

	// 检查 target 目录下是否已经存在 info.yaml 文件，如果存在则跳过, 如果不存在，则将 info 写入到 target/info.yaml 文件中
	if _, err := os.Stat(filepath.Join(b.workDir, "target", "info.yaml")); os.IsNotExist(err) {
		info.Backend = b.backendPath
		content, err := yaml.Marshal(info)
		if err != nil {
			return err
		}
		err = os.WriteFile(filepath.Join(b.workDir, "target", "info.yaml"), content, 0644)
		if err != nil {
			return err
		}
	}
	// 检查 target 目录下是否已经存在 route.json 文件，如果存在则跳过, 如果不存在，则将 route 写入到 target/route.json 文件中
	if _, err := os.Stat(filepath.Join(b.workDir, "target", "route.json")); os.IsNotExist(err) {
		route.Name = info.Name
		route.Path = "/" + strings.ToLower(info.Name)
		content, err := json.Marshal(route)
		if err != nil {
			return err
		}
		err = os.WriteFile(filepath.Join(b.workDir, "target", "route.json"), content, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Builder) Package() error {
	slog.Info(fmt.Sprintf("start package plugin %s", b.name))
	return compressToZip(filepath.Join(b.workDir, "target"), b.name)
}
