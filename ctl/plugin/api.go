package plugin

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func buildAPI(apiPath string, target string) error {
	cmd := exec.Command("goctl", "api", "go", "-api", apiPath, "-dir", target)
	initCmdOutput(cmd)
	if err := cmd.Run(); err != nil {
		return err
	}

	serviceName, err := getServiceName(apiPath)
	if err != nil {
		return err
	}

	// remove {target}/etc/{serviceName}.yaml
	etcYamlPath := filepath.Join(target, "etc", serviceName+".yaml")
	if err := os.Remove(etcYamlPath); err != nil {
		return err
	}

	// remove {target}/{serviceName}.go
	serviceGoPath := filepath.Join(target, serviceName+".go")
	if err := os.Remove(serviceGoPath); err != nil {
		return err
	}

	return nil
}

func getServiceName(apiPath string) (string, error) {
	file, err := os.Open(apiPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	pattern := `service\s+(\w+)\s+\{`
	re := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			return matches[1], nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", errors.New("service name not found")
}
