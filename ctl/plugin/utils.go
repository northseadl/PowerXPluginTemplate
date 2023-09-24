package plugin

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var ErrSkipDir = errors.New("skip dir")

func compareVersions(v1, v2 string) int {
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	for i := 0; i < len(parts1) && i < len(parts2); i++ {
		num1, _ := strconv.Atoi(parts1[i])
		num2, _ := strconv.Atoi(parts2[i])

		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}

	if len(parts1) > len(parts2) {
		return 1
	} else if len(parts1) < len(parts2) {
		return -1
	}

	return 0
}

func copyFilesExclude(srcDir, dstDir, excludeFile string) error {
	err := filepath.Walk(srcDir, func(srcPath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() != excludeFile {
			relPath, err := filepath.Rel(srcDir, srcPath)
			if err != nil {
				return err
			}

			dstPath := filepath.Join(dstDir, relPath)
			err = os.MkdirAll(filepath.Dir(dstPath), 0755)
			if err != nil {
				return err
			}

			srcFile, err := os.Open(srcPath)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			dstFile, err := os.Create(dstPath)
			if err != nil {
				return err
			}
			defer dstFile.Close()

			_, err = io.Copy(dstFile, srcFile)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func getFirstModuleName(dir string) (string, error) {
	var moduleName string

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".ts") {
			moduleName = strings.TrimSuffix(info.Name(), ".ts")
			return ErrSkipDir
		}

		return nil
	})

	if err != nil && !errors.Is(err, ErrSkipDir) {
		return "", err
	}

	return moduleName, nil
}

func copyDir(srcDir, dstDir string) error {
	err := filepath.Walk(srcDir, func(srcPath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, srcPath)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dstDir, relPath)

		if info.IsDir() {
			err = os.MkdirAll(dstPath, info.Mode())
			if err != nil {
				return err
			}
		} else {
			srcFile, err := os.Open(srcPath)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			dstFile, err := os.Create(dstPath)
			if err != nil {
				return err
			}
			defer dstFile.Close()

			_, err = io.Copy(dstFile, srcFile)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func copyFile(srcPath, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("打开源文件失败: %v", err)
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(dstPath), 0755)
	if err != nil {
		return fmt.Errorf("创建目标目录失败: %v", err)
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %v", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("复制文件失败: %v", err)
	}

	return nil
}

func compressToZip(dir string, name string) error {
	zipFile, err := os.Create(name + ".zip")
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name, _ = filepath.Rel(dir, path)

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
