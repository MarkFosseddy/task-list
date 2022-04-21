package main

import (
	"bufio"
	"errors"
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

type task struct {
	title string
	desc  string
}

type storage struct {
	file string
}

type application struct {
	input    *cmdInput
	tasks    []task
	message  string
	editing  bool
	deleting bool
}

func (i *cmdInput) scan() bool {
	return i.s.Scan()
}

func (i *cmdInput) text() string {
	return strings.TrimSpace(i.s.Text())
}

func (s *storage) read() []task {
	f, err := os.OpenFile(s.file, os.O_RDONLY|os.O_CREATE, 0644)
	assert(err == nil, err)
	defer f.Close()

	bytes, err := io.ReadAll(f)
	assert(err == nil, err)

	data := strings.TrimSpace(string(bytes))
	items := []task{}

	if len(data) > 0 {
		err := ""
		for _, line := range strings.Split(data, "\n") {
			if title, desc, found := strings.Cut(line, " <-$-> "); found {
				if desc == "-$-" {
					desc = ""
				}
				items = append(items, task{title, desc})
			} else {
				err += fmt.Sprintf("Error parsing task: %s\n", line)
			}
		}
		assert(err == "", errors.New(err))
	}

	return items
}

func (s *storage) write(data []task) {
	f, err := os.OpenFile(s.file, os.O_WRONLY|os.O_TRUNC, 0644)
	assert(err == nil, err)
	defer f.Close()

	buf := ""
	for _, item := range data {
		desc := item.desc
		if desc == "" {
			desc = "-$-"
		}
		buf += fmt.Sprintf("%s <-$-> %s\n", item.title, desc)
	}

	_, err = f.WriteString(buf)
	assert(err == nil, err)
}

func assert(cond bool, err error) {
	if !cond {
		log.Fatal(err)
	}
}

var store storage = storage{"/home/fosseddy/.task-list"}

var app application = application{
	input:    &cmdInput{bufio.NewScanner(os.Stdin)},
	tasks:    []task{},
	message:  "",
	editing:  false,
	deleting: false,
}

func draw() {
	// @NOTE(art): clear screen
	fmt.Print("\033c")

	fmt.Println("Task List")
	fmt.Println()

	for i, t := range app.tasks {
		if app.deleting || app.editing {
			fmt.Printf(
				"  [%v] title: %v\n  description: %v\n\n",
				i+1, t.title, t.desc,
			)
		} else {
			fmt.Printf("  title: %v\n  description: %v\n\n", t.title, t.desc)
		}
	}

	if app.message != "" {
		fmt.Println(app.message)
	} else {
		fmt.Println()
	}

	fmt.Print(">> ")
}

func main() {
	app.tasks = store.read()
	exit := false

	for !exit {
		draw()

		app.message = ""
		app.input.scan()
		cmd := app.input.text()
		if cmd == "" {
			continue
		}

		switch cmd {
		case "add":
			app.message = "Enter title (empty to cancel):"
			draw()
			app.input.scan()
			title := app.input.text()

			if title != "" {
				app.message = "Enter description (optional):"
				draw()
				app.input.scan()
				desc := app.input.text()
				app.tasks = append(app.tasks, task{title, desc})
			}

			store.write(app.tasks)

			app.message = ""
		case "delete":
			if len(app.tasks) == 0 {
				app.message = "You have no tasks. Use `add` command to create one"
				break
			}

			app.message = "Choose task to delete (empty to cancel):"
			app.deleting = true
			id := -1

			for {
				draw()
				app.input.scan()
				text := app.input.text()

				if text == "" {
					break
				}

				if val, err := strconv.Atoi(text); err == nil {
					if val >= 1 && val <= len(app.tasks) {
						// @NOTE(art): in ui we draw ids starting from 1
						id = val - 1
						break
					}
				}
			}

			if id != -1 {
				var tmp []task
				for i, t := range app.tasks {
					if i != id {
						tmp = append(tmp, t)
					}
				}
				app.tasks = tmp
				store.write(app.tasks)
			}

			app.deleting = false
			app.message = ""
		case "edit":
			if len(app.tasks) == 0 {
				app.message = "You have no tasks. Use `add` command to create one"
				break
			}

			app.message = "Choose task to edit (empty to cancel):"
			app.editing = true
			id := -1

			for {
				draw()
				app.input.scan()
				text := app.input.text()

				if text == "" {
					break
				}

				if val, err := strconv.Atoi(text); err == nil {
					if val >= 1 && val <= len(app.tasks) {
						// @NOTE(art): in ui we draw ids starting from 1
						id = val - 1
						break
					}
				}
			}

			if id != -1 {
				app.message = "Enter new title (empty to skip):"
				draw()
				app.input.scan()
				title := app.input.text()

				app.message = "Enter new description (empty to skip):"
				draw()
				app.input.scan()
				desc := app.input.text()

				changed := false

				if title != "" {
					app.tasks[id].title = title
					changed = true
				}

				if desc != "" {
					app.tasks[id].desc = desc
					changed = true
				}

				if changed {
					store.write(app.tasks)
				}
			}

			app.editing = false
			app.message = ""
		case "exit":
			exit = true
		case "help":
			app.message = "add\ndelete\nedit\nhelp\nexit\n"
		default:
			app.message = fmt.Sprintf(
				"Unknown command `%s`. Type `help` to see commands",
				cmd,
			)
		}
	}
}
