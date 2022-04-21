package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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
	title string
	desc  string
}

type storage struct {
	file string
}

func (s *storage) read() []task {
	fd, err := os.OpenFile(s.file, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	bytes, err := io.ReadAll(fd)
	if err != nil {
		log.Fatal(err)
	}

	data := strings.TrimSpace(string(bytes))
	items := []task{}

	if len(data) > 0 {
		for _, line := range strings.Split(data, "\n") {
			if title, desc, found := strings.Cut(line, " <-$-> "); found {
				if desc == "-$-" {
					desc = ""
				}
				items = append(items, task{title, desc})
			} else {
				message += fmt.Sprintf("Error parsing task: %s\n", line)
			}
		}
	}

	return items
}

func (s *storage) write(data []task) {
	fd, err := os.OpenFile(s.file, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	buf := ""
	for _, item := range data {
		desc := item.desc
		if len(desc) == 0 {
			desc = "-$-"
		}
		buf += fmt.Sprintf("%s <-$-> %s\n", item.title, desc)
	}

	_, err = fd.WriteString(buf)
	if err != nil {
		log.Fatal(err)
	}
}

var message string = ""
var tasks []task = []task{}
var deleting bool = false
var updating bool = false

func draw() {
	// @NOTE(art): clear screen
	fmt.Print("\033c")

	fmt.Println("Task List")
	fmt.Println()

	for i, t := range tasks {
		if deleting || updating {
			fmt.Printf(
				"  [%v] title: %v\n  description: %v\n\n",
				i+1, t.title, t.desc,
			)
		} else {
			fmt.Printf("  title: %v\n  description: %v\n\n", t.title, t.desc)
		}
	}

	if len(message) > 0 {
		fmt.Println(message)
	} else {
		fmt.Println()
	}

	fmt.Print(">> ")
}

func main() {
	store := storage{"/home/fosseddy/.task-list"}
	tasks = store.read()

	exit := false
	input := cmdInput{bufio.NewScanner(os.Stdin)}

	for !exit {
		draw()

		message = ""
		input.scan()
		cmd := input.text()
		if len(cmd) == 0 {
			continue
		}

		switch cmd {
		case "add":
			message = "Enter title (empty to cancel):"
			draw()
			input.scan()
			title := input.text()

			if len(title) > 0 {
				message = "Enter description (optional):"
				draw()
				input.scan()
				desc := input.text()
				tasks = append(tasks, task{title, desc})
			}

			store.write(tasks)

			message = ""
		case "delete":
			if len(tasks) == 0 {
				message = "You have no tasks. Use `add` command to create one"
				break
			}

			message = "Choose task to delete (empty to cancel):"
			deleting = true
			id := -1

			for {
				draw()
				input.scan()
				text := input.text()

				if len(text) == 0 {
					break
				}

				if val, err := strconv.Atoi(text); err == nil {
					if val >= 1 && val <= len(tasks) {
						// @NOTE(art): in ui we draw ids starting from 1
						id = val - 1
						break
					}
				}
			}

			if id != -1 {
				var tmp []task
				for i, t := range tasks {
					if i != id {
						tmp = append(tmp, t)
					}
				}
				tasks = tmp
				store.write(tasks)
			}

			deleting = false
			message = ""
		case "update":
			if len(tasks) == 0 {
				message = "You have no tasks. Use `add` command to create one"
				break
			}

			message = "Choose task to update (empty to cancel):"
			updating = true
			id := -1

			for {
				draw()
				input.scan()
				text := input.text()

				if len(text) == 0 {
					break
				}

				if val, err := strconv.Atoi(text); err == nil {
					if val >= 1 && val <= len(tasks) {
						// @NOTE(art): in ui we draw ids starting from 1
						id = val - 1
						break
					}
				}
			}

			if id != -1 {
				message = "Enter new title (empty to skip):"
				draw()
				input.scan()
				title := input.text()

				message = "Enter new description (empty to skip):"
				draw()
				input.scan()
				desc := input.text()

				changed := false

				if len(title) > 0 {
					tasks[id].title = title
					changed = true
				}

				if len(desc) > 0 {
					tasks[id].desc = desc
					changed = true
				}

				if changed {
					store.write(tasks)
				}
			}

			updating = false
			message = ""
		case "exit":
			exit = true
		case "help":
			message = "add\ndelete\nupdate\nhelp\nexit\n"
		default:
			message = fmt.Sprintf(
				"Unknown command `%s`. Type `help` to see commands",
				cmd,
			)
		}
	}
}
