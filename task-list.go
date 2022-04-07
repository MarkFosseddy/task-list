package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	exit := false
	msg := ""
	input := bufio.NewScanner(os.Stdin)

	for !exit {
		// draw
		fmt.Print("\033c")

		fmt.Println("Task List")
		fmt.Println()

		if len(msg) > 0 {
			fmt.Println(msg)
		} else {
			fmt.Println()
		}

		fmt.Print(">> ")

		// update
		msg = ""

		input.Scan()
		cmd := strings.Trim(input.Text(), " ")
		if len(cmd) == 0 {
			continue
		}

		switch cmd {
		case "q":
			exit = true

		default:
			msg = fmt.Sprintf("Unknown command `%s`", cmd)
		}
	}
}
