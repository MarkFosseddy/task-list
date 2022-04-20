package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type task struct {
	id    string
	title string
	desc  string
}

var message string = ""
var tasks []task = []task{}

func draw() {
	// @NOTE(art): clear screen
	fmt.Print("\033c")

	fmt.Println("Task List")
	fmt.Println()

	for _, t := range tasks {
		fmt.Printf(
			"  id: %v\n  title: %v\n  description: %v\n\n",
			t.id, t.title, t.desc,
		)
	}

	if len(message) > 0 {
		fmt.Println(message)
	} else {
		fmt.Println()
	}

	fmt.Print(">> ")
	// @NOTE(art): kinda hacky to reset message here, but otherwise you have to
	// do it in every command handler
	message = ""
}

func main() {
	exit := false
	input := bufio.NewScanner(os.Stdin)

	for !exit {
		draw()

		input.Scan()
		cmd := strings.Trim(input.Text(), " ")
		if len(cmd) == 0 {
			continue
		}

		switch cmd {
		case "add":
			message = "Enter title(empty to cancel):"
			draw()
			input.Scan()
			title := strings.Trim(input.Text(), " ")
			if len(title) == 0 {
				break
			}

			message = "Enter description(optional):"
			draw()
			input.Scan()
			desc := strings.Trim(input.Text(), " ")

			tasks = append(tasks, task{
				id:    strconv.FormatUint(rand.Uint64(), 36),
				title: title,
				desc:  desc,
			})
		case "q":
			exit = true

		default:
			message = fmt.Sprintf("Unknown command `%s`", cmd)
		}
	}
}
