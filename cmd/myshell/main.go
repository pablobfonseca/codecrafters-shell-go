package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/utils"
)

var builtInCommands = []string{"exit", "echo", "type", "pwd", "cd"}

func main() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		command, args := utils.ParseUserInput(userInput)

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
			} else if path, ok := utils.FindExecutable(typeCmd, paths); ok {
				fmt.Fprintf(os.Stdout, "%s is %s\n", typeCmd, path)
				break
			} else {
				fmt.Fprintf(os.Stdout, "%s: not found\n", typeCmd)
				break
			}
		case "pwd":
			path, _ := os.Getwd()
			fmt.Fprintln(os.Stdout, path)
		case "cd":
			path := args[0]
			home := os.Getenv("HOME")
			if path == "~" {
				path = home
			}
			if strings.HasPrefix(path, "~") {
				path, _ = strings.CutPrefix(path, "~/")
				path = fmt.Sprintf("%s/%s", home, path)
			}
			err := os.Chdir(path)
			if err != nil {
				fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", path)
			}
		default:
			utils.ExecCommand(command, args)
		}
	}
}
