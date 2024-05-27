package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

		if command == "exit" {
			if args[0] == "0" {
				os.Exit(0)
			}
		}

		fmt.Printf("%s: command not found\n", strings.TrimRight(command, "\n"))
	}
}
