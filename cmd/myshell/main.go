package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	ExitCmd = "exit"
	EchoCmd = "echo"
	TypeCmd = "type"
)

func main() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		command = strings.TrimRight(command, "\n")
		cmdWithArgs := strings.Split(command, " ")

		command = cmdWithArgs[0]
		args := cmdWithArgs[1:]

		switch command {
		case ExitCmd:
			os.Exit(0)
		case EchoCmd:
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
			break
		case TypeCmd:
			typeCmd := args[0]
			if typeCmd == ExitCmd || typeCmd == EchoCmd || typeCmd == TypeCmd {
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
			fmt.Fprintf(os.Stdout, "%s: command not found\n", strings.TrimRight(command, "\n"))
		}
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
