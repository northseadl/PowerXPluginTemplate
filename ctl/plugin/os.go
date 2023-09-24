package plugin

import (
	"os"
	"os/exec"
)

var (
	stdout = os.Stdout
	stderr = os.Stderr
)

// SetStdout 设置标准输出
func SetStdout(f *os.File) {
	stdout = f
}

// SetStderr 设置标准错误输出
func SetStderr(f *os.File) {
	stderr = f
}

// 重定向 cmd 的标准输出和标准错误输出
func initCmdOutput(cmd *exec.Cmd) {
	cmd.Stdout = stdout
	cmd.Stderr = stderr
}
