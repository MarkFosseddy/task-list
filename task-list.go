package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type cmdInput struct {
	s *bufio.Scanner
}

func (i *cmdInput) scan() bool {
	return i.s.Scan()
}

func (i *cmdInput) text() string {
	return strings.TrimSpace(i.s.Text())
}

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
	input := cmdInput{bufio.NewScanner(os.Stdin)}

	for !exit {
		draw()

		input.scan()
		cmd := input.text()
		if len(cmd) == 0 {
			continue
		}

		switch cmd {
		case "add":
			message = "Enter title(empty to cancel):"
			draw()
			input.scan()
			title := input.text()
			if len(title) == 0 {
				break
			}

			message = "Enter description(optional):"
			draw()
			input.scan()
			desc := input.text()

			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			tasks = append(tasks, task{
				id:    strconv.FormatUint(r.Uint64(), 36),
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
