// 封装git命令及其执行抽象
package git

import (
	"bytes"
	"os/exec"
)

// git命令
const gitCommand = "git"

// 执行git命令并响应执行结果
// @param args git命令行参数
func runGit(args ...string) (string, error) {
	buffer := new(bytes.Buffer)
	cmd := exec.Command(gitCommand, args...)
	cmd.Stdout = buffer
	cmd.Stderr = buffer
	err := cmd.Run()
	return buffer.String(), err
}

// 执行 git version 命令
func RunGitVersion() (string, error) {
	return runGit("version")
}

// 执行 git alias 命令
// @param alias 别名
// @param cmd 原命令名
func RunGitAlias(alias string, cmd string) error {
	_, err := runGit("config", "alias."+alias, cmd)
	return err
}
