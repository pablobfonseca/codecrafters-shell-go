package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

var builtInCommands = []string{"exit", "echo", "type"}

func main() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		command := strings.TrimRight(userInput, "\n")
		cmdWithArgs := strings.Split(command, " ")

		command = cmdWithArgs[0]
		args := cmdWithArgs[1:]

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
			break
		case "type":
			typeCmd := args[0]
			if slices.Contains(builtInCommands, typeCmd) {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", typeCmd)
				break
			} else if path, ok := findExecutable(typeCmd, paths); ok {
				fmt.Fprintf(os.Stdout, "%s is %s\n", typeCmd, path)
				break
			} else {
				fmt.Fprintf(os.Stdout, "%s not found\n", typeCmd)
				break
			}
		default:
			execCommand(command, args)
		}
	}
}

func execCommand(command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}

func findExecutable(name string, paths []string) (string, bool) {
	for _, path := range paths {
		fullpath := filepath.Join(path, name)

		if _, err := os.Stat(fullpath); err == nil {
			return fullpath, true
		}
	}

	return "", false
}
