package plugin

import (
	"archive/zip"
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	projectUrl = "https://github.com/northseadl/PowerXPluginTemplate/archive/refs/heads/master.zip"
)

type Starter struct {
	project     string
	description string
	menuName    string
	icon        string
}

func NewStarter() *Starter {
	return &Starter{}
}

func (s *Starter) Start() error {
	// scan project name
	for {
		color.Green("[Step1]: Please input project name, only support hyphen format, must be lowercase, e.g. powerx-plugin")
		if _, err := fmt.Scanln(&s.project); err != nil {
			return err
		}
		if err := validateProjectName(s.project); err != nil {
			color.Red(err.Error())
			continue
		}
		break
	}

	// scan description
	color.Green("[Step2]: Please input project description, e.g. PowerX plugin")
	if _, err := fmt.Scanln(&s.description); err != nil {
		return err
	}

	// scan menu name
	color.Green("[Step3]: Please input menu name, e.g. PowerX plugin")
	if _, err := fmt.Scanln(&s.menuName); err != nil {
		return err
	}

	// scan icon
	color.Green("[Step4]: Please input icon, e.g. icon-plus, refer to https://arco.design/vue/component/icon")
	if _, err := fmt.Scanln(&s.icon); err != nil {
		return err
	}

	// clone template
	color.Green("[Step5]: Cloning template...")
	if err := downloadProject(projectUrl, "."); err != nil {
		return err
	}
	if err := os.Rename("PowerXPluginTemplate-master", s.project); err != nil {
		return err
	}

	color.Green("[Step6]: Preparing...")

	// remove useless files
	removeDir := []string{"ctl"}
	for _, dir := range removeDir {
		_ = os.RemoveAll(filepath.Join(s.project, dir))
	}
	removeFiles := []string{"LICENSE", "README.md", "go.mod", "go.sum", ".gitignore"}
	for _, file := range removeFiles {
		_ = os.Remove(filepath.Join(s.project, file))
	}
	info := Info{
		Name:        s.project,
		Description: s.description,
		Version:     "v1.0.0",
		Metadata: struct {
			Icon   string `yaml:"icon"`
			Locale string `yaml:"locale"`
		}{
			Icon:   s.icon,
			Locale: s.menuName,
		},
	}

	// replace backend module name
	if err := replaceBackendModuleName(filepath.Join(s.project, "backend"), s.project); err != nil {
		return err
	}

	// replace frontend module name
	if err := replaceFrontendModuleName(filepath.Join(s.project, "frontend"), s.project); err != nil {
		return err
	}

	// write plugin.yaml
	data, err := yaml.Marshal(info)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(s.project, "plugin.yaml"), data, 0755); err != nil {
		return err
	}

	color.Blue("Done!")

	return nil
}

func validateProjectName(project string) error {
	// 最少12个字符，最多40个字符
	reg := regexp.MustCompile(`^[a-z][a-z0-9-]{10,38}[a-z0-9]$`)
	if !reg.MatchString(project) {
		return fmt.Errorf("project name must be lowercase and hyphen format, and length must be between 12 and 40")
	}
	return nil
}

func downloadProject(url string, asName string) error {
	// download template
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// save as zip
	out, err := os.Create("template.zip")
	if err != nil {
		return err
	}
	defer os.Remove("template.zip")
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}

	// unzip to project dir
	if err := unzipToDir("template.zip", asName); err != nil {
		return err
	}
	return nil
}

func unzipToDir(zipFile string, dir string) error {
	// open zip
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	// create dir
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// unzip
	for _, file := range zipReader.Reader.File {
		// open file in zip
		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		// create file
		targetDir := filepath.Join(dir, file.Name)
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(targetDir, file.Mode()); err != nil {
				return err
			}
		} else {
			if err := os.MkdirAll(filepath.Dir(targetDir), file.Mode()); err != nil {
				return err
			}

			targetFile, err := os.OpenFile(targetDir, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer targetFile.Close()

			if _, err := io.Copy(targetFile, fileReader); err != nil {
				return err
			}
		}
	}
	return nil
}

func replaceBackendModuleName(dir string, project string) error {
	// replace backend go.mod
	backendGoMod := filepath.Join(dir, "go.mod")
	data, err := os.ReadFile(backendGoMod)
	if err != nil {
		return err
	}
	content := string(data)
	content = strings.ReplaceAll(content, "PluginTemplate", project)
	if err := os.WriteFile(backendGoMod, []byte(content), 0755); err != nil {
		return err
	}

	// replace other xx.go/xx.api
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		content := string(data)
		content = strings.ReplaceAll(content, "PluginTemplate", project)
		if err := os.WriteFile(path, []byte(content), 0755); err != nil {
			return err
		}
		return nil
	})
	return nil
}

func replaceFrontendModuleName(dir string, project string) error {
	var oldPaths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Index(path, "plugin-example") != -1 && info.IsDir() {
			oldPaths = append(oldPaths, path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	for _, oldPath := range oldPaths {
		newPath := strings.ReplaceAll(oldPath, "plugin-example", project)
		if err := os.Rename(oldPath, newPath); err != nil {
			return err
		}
	}
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Index(path, "plugin-example") != -1 {
			if !info.IsDir() {
				content, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				newContent := strings.ReplaceAll(string(content), "plugin-example", project)
				newPath := strings.ReplaceAll(path, "plugin-example", project)
				if err := os.WriteFile(newPath, []byte(newContent), 0755); err != nil {
					return err
				}
				if err := os.Remove(path); err != nil {
					return err
				}
				return nil
			}
		}
		return nil
	})
}
