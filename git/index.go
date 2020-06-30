package git

import (
	"bytes"
	"os/exec"
)

const gitCommand = "git"

func runGit(args ...string) (string, error) {
	buffer := new(bytes.Buffer)
	cmd := exec.Command(gitCommand, args...)
	cmd.Stdout = buffer
	cmd.Stderr = buffer
	err := cmd.Run()
	return buffer.String(), err
}

func RunGitVersion() (string, error) {
	return runGit("version")
}

func RunGitAlias(alias string, cmd string) error {
	_, err := runGit("config", "alias."+alias, cmd)
	return err
}
