package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ExecCommand(command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

func FindExecutable(name string, paths []string) (string, bool) {
	for _, path := range paths {
		fullpath := filepath.Join(path, name)

		if _, err := os.Stat(fullpath); err == nil {
			return fullpath, true
		}
	}

	return "", false
}

func ParseUserInput(input string) (command string, args []string) {
	command = strings.TrimRight(input, "\n")
	cmdWithArgs := strings.Split(command, " ")

	command = cmdWithArgs[0]
	args = cmdWithArgs[1:]

	return
}
