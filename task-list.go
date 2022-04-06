package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

type State struct {
    message string
    tasks []Task
}

type Task struct {
    id int
    title string
    description string
}

func draw(s *State) {
    fmt.Print("\033c")

    fmt.Println("Task List")
    fmt.Println(s.tasks)
    fmt.Println()

    if len(s.message) > 0 {
        fmt.Println(s.message)
    } else {
        fmt.Println()
    }

    fmt.Print(">> ")
}

func main() {
    state := &State{}
    state.message = ""
    state.tasks = []Task{}

    id := 1
    exit := false

    input := bufio.NewScanner(os.Stdin)

    for !exit {
        state.message = ""
        draw(state)

        input.Scan()
        cmd := strings.Trim(input.Text(), " ")
        if len(cmd) == 0 { continue }

        switch cmd {
        case "q":
            exit = true

        case "add": {
            state.message = "Enter title (empty to cancel):"
            draw(state)

            input.Scan()
            title := strings.Trim(input.Text(), " ")
            if len(title) == 0 { continue }

            state.message = "Enter description (optional):"
            draw(state)

            input.Scan()
            desc := strings.Trim(input.Text(), " ")

            t := Task{}
            t.id = id
            t.title = title
            t.description = desc

            state.tasks = append(state.tasks, t)

            id += 1
        }

        default:
            state.message = fmt.Sprintf("Unknown command `%s`", cmd)
        }
    }
}
