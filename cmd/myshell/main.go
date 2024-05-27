package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	ExitCmd = "exit"
	EchoCmd = "echo"
	TypeCmd = "type"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		in, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
		}

		command := strings.TrimRight(in, "\n")
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
			} else {
				fmt.Fprintf(os.Stdout, "%s not found\n", typeCmd)
				break
			}
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", strings.TrimRight(command, "\n"))
		}
	}
}
