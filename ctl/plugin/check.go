package plugin

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

func checkCode(workDir string) error {
	// Parse {workDir}/plugin.yaml to struct
	var pluginInfo Info
	pluginYamlPath := filepath.Join(workDir, "plugin.yaml")
	pluginYamlFile, err := os.Open(pluginYamlPath)
	if err != nil {
		return err
	}
	defer pluginYamlFile.Close()

	pluginYamlData, err := io.ReadAll(pluginYamlFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(pluginYamlData, &pluginInfo)
	if err != nil {
		return err
	}

	// {Info.Name} must be in camel case
	if !isHyphenCase(pluginInfo.Name) {
		return errors.New("Info.Name must be in hyphen case")
	}

	// {workDir}/etc/etc.yaml must exist, and is valid yaml
	etcYamlPath := filepath.Join(workDir, "backend/etc/etc.yaml")
	_, err = os.Stat(etcYamlPath)
	if os.IsNotExist(err) {
		return errors.New("etc/etc.yaml must exist")
	}

	etcYamlFile, err := os.Open(etcYamlPath)
	if err != nil {
		return err
	}
	defer etcYamlFile.Close()

	etcYamlData, err := io.ReadAll(etcYamlFile)
	if err != nil {
		return err
	}

	var etcYaml any
	err = yaml.Unmarshal(etcYamlData, &etcYaml)
	if err != nil {
		return errors.New("etc/etc.yaml must be valid yaml")
	}

	// {workDir}/frontend/api/*.ts, only interceptor.ts or {Info.Name}.ts is allowed
	apiDir := filepath.Join(workDir, "frontend/src/api")
	err = filepath.Walk(apiDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".ts" {
			fileName := filepath.Base(path)
			if fileName != "interceptor.ts" && fileName != fmt.Sprintf("%s.ts", pluginInfo.Name) {
				return fmt.Errorf("only interceptor.ts or %s.ts is allowed in frontend/src/api", pluginInfo.Name)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	// {workDir}/frontend/module.ts, only {Info.Name}.ts is allowed
	moduleDir := filepath.Join(workDir, "frontend/src/router/routes/modules")
	err = filepath.Walk(moduleDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".ts" {
			fileName := filepath.Base(path)
			if fileName != fmt.Sprintf("%s.ts", pluginInfo.Name) {
				return fmt.Errorf("only %s.ts is allowed in frontend/src/router/routes/modules", pluginInfo.Name)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	// {workDir}/frontend/views/{Info.Name} dir must exist
	viewsDir := filepath.Join(workDir, "frontend/src/views", pluginInfo.Name)
	_, err = os.Stat(viewsDir)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s dir must exist", viewsDir)
	}

	return nil
}

// 是否是连缀符格式
func isHyphenCase(s string) bool {
	hyphenPattern := `^[a-z]+(?:-[a-z]+)*$`
	match, _ := regexp.MatchString(hyphenPattern, s)
	return match
}
