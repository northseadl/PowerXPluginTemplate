package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
)

var buildCmd = &cli.Command{
	Name:  "build",
	Usage: "build plugin",
	Action: func(c *cli.Context) error {
		return nil
	},
}

const defaultBackendDir = "backend"
const defaultFrontendDir = "frontend"
const requiredGoVersion = "1.21"
const requiredNodeVersion = "16.0.0"

type Builder struct {
	workDir string
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
	// 检查是否存在 golang 且版本号 大于等于 1.21
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("无法检测 Golang 版本: %v", err)
	}

	versionRegexp := regexp.MustCompile(`go(\d+\.\d+(?:\.\d+)?)`)
	matches := versionRegexp.FindStringSubmatch(string(output))
	if len(matches) < 2 {
		return fmt.Errorf("无法解析 Golang 版本")
	}

	currentVersion := matches[1]
	if compareVersions(currentVersion, requiredGoVersion) < 0 {
		return fmt.Errorf("golang 版本过低: 当前版本 %s, 要求版本 >= %s", currentVersion, requiredGoVersion)
	}

	// 检查当前环境 darwin/linux/windows
	slog.Info(fmt.Sprintf("当前构建环境为 %s", runtime.GOOS))

	// 检查是否存在 nodejs 且版本号 大于等于 16.0.0
	cmd = exec.Command("node", "-v")
	output, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("无法检测 Node.js 版本: %v", err)
	}

	versionRegexp = regexp.MustCompile(`v(\d+\.\d+\.\d+)`)
	matches = versionRegexp.FindStringSubmatch(string(output))
	if len(matches) < 2 {
		return fmt.Errorf("无法解析 Node.js 版本")
	}

	currentVersion = matches[1]
	if compareVersions(currentVersion, requiredNodeVersion) < 0 {
		return fmt.Errorf("node 版本过低: 当前版本 %s, 要求版本 >= %s", currentVersion, requiredNodeVersion)
	}
	return nil
}

func (b *Builder) BuildBackend(relativeDir string) error {
	// 提取 relativeDir/etc/etc.yaml 文件, 提取到 workDir/target/etc.yaml
	srcPath := filepath.Join(b.workDir, relativeDir, "etc", "etc.yaml")
	dstPath := filepath.Join(b.workDir, "target", "etc.yaml")

	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(dstPath), 0755)
	if err != nil {
		return err
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// 使用 go build 编译 relativeDir目录, 编译以后的结果提取到 workDir/target 目录下
	cmd := exec.Command("go", "build", "-o", filepath.Join(b.workDir, "target", "backend"), filepath.Join(b.workDir, relativeDir))
	cmd.Env = append(os.Environ(), "GO111MODULE=on")
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (b *Builder) BuildFrontend(relativeDir string) error {
	// 执行 npm install && npm run build， 如果出错则返回错误
	slog.Info(fmt.Sprintf("开始构建前端项目 %s", relativeDir))
	cmd := exec.Command("npm", "install")
	cmd.Dir = filepath.Join(b.workDir, relativeDir)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("执行 npm install 失败: %v", err)
	}

	cmd = exec.Command("npm", "run", "build")
	cmd.Dir = filepath.Join(b.workDir, relativeDir)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("执行 npm run build 失败: %v", err)
	}

	// 提取 relativeDir/src/api 目录除了 interceptor.ts 以外的所有文件, 提取到 workDir/target/api 目录下
	err = copyFilesExclude(filepath.Join(b.workDir, relativeDir, "src", "api"), filepath.Join(b.workDir, "target", "api"), "interceptor.ts")
	if err != nil {
		return fmt.Errorf("复制 API 文件失败: %v", err)
	}

	// 提取 relativeDir/src/router/routes/modules 目录下的第一个文件, 获取其name, 提取为 workDir/target/module.ts
	moduleName, err := getFirstModuleName(filepath.Join(b.workDir, relativeDir, "src", "router", "routes", "modules"))
	if err != nil {
		return fmt.Errorf("获取模块名称失败: %v", err)
	}

	err = copyFile(filepath.Join(b.workDir, relativeDir, "src", "router", "routes", "modules", moduleName+".ts"), filepath.Join(b.workDir, "target", "module.ts"))
	if err != nil {
		return fmt.Errorf("复制模块文件失败: %v", err)
	}

	// 提取 relativeDir/src/views 目录下文件名为 name 的目录, 将其下的所有文件提取到 workDir/target/views 目录下
	err = copyDir(filepath.Join(b.workDir, relativeDir, "src", "views", moduleName), filepath.Join(b.workDir, "target", "views"))
	if err != nil {
		return fmt.Errorf("复制视图文件失败: %v", err)
	}

	return nil
}
