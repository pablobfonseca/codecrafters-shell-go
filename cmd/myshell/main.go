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

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("%s: command not found\n", strings.TrimRight(command, "\n"))
	}
}
