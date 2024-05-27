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
			fmt.Println(err.Error())
		}

		command := strings.TrimRight(in, "\n")
		cmd_with_args := strings.Split(command, " ")

		command = cmd_with_args[0]
		args := cmd_with_args[1:]

		switch command {
		case ExitCmd:
			if args[0] == "0" {
				os.Exit(0)
			}
		case EchoCmd:
			fmt.Printf("%s\n", strings.Join(args, " "))
			break
		case TypeCmd:
			type_cmd := args[0]
			if type_cmd == ExitCmd || type_cmd == EchoCmd || type_cmd == TypeCmd {
				fmt.Printf("%s is a shell builtin\n", type_cmd)
				break
			} else {
				fmt.Printf("%s not found\n", type_cmd)
				break
			}
		default:
			fmt.Printf("%s: command not found\n", strings.TrimRight(command, "\n"))
		}
	}
}
