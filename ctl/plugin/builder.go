package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

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

// Check checks dependencies
func (b *Builder) Check() error {
	// Check the current environment: darwin/linux/windows
	slog.Info(fmt.Sprintf("os is %s", runtime.GOOS))
	// Check code
	if err := checkCode(b.workDir); err != nil {
		return err
	}
	slog.Info("check code success")
	return nil
}

// BuildBackend builds the backend
func (b *Builder) BuildBackend(relativeDir string) error {
	// Execute go mod download, return an error if it fails
	slog.Info(fmt.Sprintf("start build backend project %s", relativeDir))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS signals to ensure graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		cancel()
	}()

	cmd := exec.CommandContext(ctx, "go", "mod", "download")
	cmd.Dir = filepath.Join(b.workDir, relativeDir)
	cmd.Env = append(os.Environ(), "GO111MODULE=on")
	initCmdOutput(cmd)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("exec go mod download failed: %v", err)
	}

	// Extract etc/etc.yaml file from relativeDir, and save it to workDir/target/etc.yaml
	srcPath := filepath.Join(b.workDir, relativeDir, "etc/etc.yaml")
	dstPath := filepath.Join(b.workDir, "target/etc.yaml")

	err := copyFile(srcPath, dstPath)
	if err != nil {
		return err
	}

	// Use go build to compile the relativeDir directory, and save the compiled result to workDir/target directory
	buildDir := filepath.Join(b.workDir, relativeDir)
	var outputPath string
	if runtime.GOOS == "windows" {
		outputPath = filepath.Join(b.workDir, "target/backend.exe")
	} else {
		outputPath = filepath.Join(b.workDir, "target/backend")
	}
	cmd = exec.Command("go", "build", "-o", outputPath, ".")
	cmd.Dir = buildDir
	initCmdOutput(cmd)
	if err = cmd.Run(); err != nil {
		return err
	}
	b.backendPath = filepath.Base(outputPath)
	return nil
}

// BuildFrontend builds the frontend
func (b *Builder) BuildFrontend(relativeDir string) error {
	// Execute npm install && npm run build, return an error if it fails
	slog.Info(fmt.Sprintf("start build frontend project %s", relativeDir))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS signals to ensure graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		cancel()
	}()

	cmd := exec.CommandContext(ctx, "yarn", "install")
	cmd.Dir = filepath.Join(b.workDir, relativeDir)
	initCmdOutput(cmd)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("exec npm install failed: %v", err)
	}

	cmd = exec.CommandContext(ctx, "yarn", "run", "prebuild")
	cmd.Dir = filepath.Join(b.workDir, relativeDir)
	initCmdOutput(cmd)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("exec npm run build failed: %v", err)
	}

	slog.Info(fmt.Sprintf("build frontend project %s success", relativeDir))

	// Extract all files except interceptor.ts from relativeDir/src/api directory, and save them to workDir/target/api directory
	err := copyFilesExclude(filepath.Join(b.workDir, relativeDir, "src", "api"), filepath.Join(b.workDir, "target", "frontend", "api"), "interceptor.ts")
	if err != nil {
		return fmt.Errorf("copy api files failed: %v", err)
	}

	// Extract the first file from relativeDir/src/router/routes/modules directory, get its name, and save it as workDir/target/module.ts
	moduleName, err := getFirstModuleName(filepath.Join(b.workDir, relativeDir, "src", "router", "routes", "modules"))
	if err != nil {
		return fmt.Errorf("copy module file failed: %v", err)
	}

	err = copyFile(filepath.Join(b.workDir, relativeDir, "src", "router", "routes", "modules", moduleName+".ts"), filepath.Join(b.workDir, "target", "frontend", "module.ts"))
	if err != nil {
		return fmt.Errorf("copy module file failed: %v", err)
	}

	// Extract the directory with the name "name" from relativeDir/src/views, and save all its files to workDir/target/views directory
	err = copyDir(filepath.Join(b.workDir, relativeDir, "src", "views", moduleName), filepath.Join(b.workDir, "target", "frontend", "views"))
	if err != nil {
		return fmt.Errorf("copy views files failed: %v", err)
	}

	// Extract the images folder from assets, and save it to workDir/target/assets directory
	err = copyDir(filepath.Join(b.workDir, relativeDir, "src", "assets", "images"), filepath.Join(b.workDir, "target", "frontend", "assets", "images"))

	return nil
}

func (b *Builder) BuildExtra() error {
	// Check if there is a plugin.yaml file in the root directory, read it if it exists, and return an error if it doesn't
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

	// Check if there is an info.yaml file in the target directory, if it exists, skip it, if it doesn't, write info to target/info.yaml file
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

	// Check if there is a route.json file in the target directory, if it exists, skip it, if it doesn't, write route to target/route.json file
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
