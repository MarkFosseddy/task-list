package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

type Task struct {
    id int
    title string
    description string
}

func draw(tasks []Task, message string) {
    fmt.Print("\033c")

    fmt.Println("Task List")
    fmt.Println(tasks)
    fmt.Println()

    if len(message) > 0 {
        fmt.Println(message)
    } else {
        fmt.Println()
    }

    fmt.Print(">> ")
}

func main() {
    id := 1
    exit := false
    message := ""
    tasks := []Task{}

    input := bufio.NewScanner(os.Stdin)

    for !exit {
        message = ""

        draw(tasks, message)

        input.Scan()
        cmd := strings.Trim(input.Text(), " ")
        if len(cmd) == 0 { continue }

        switch cmd {
        case "q":
            exit = true

        case "add": {
            message = "Enter title (empty to cancel):"
            draw(tasks, message)

            input.Scan()
            title := strings.Trim(input.Text(), " ")
            if len(title) == 0 { continue }

            message = "Enter description (optional):"
            draw(tasks, message)

            input.Scan()
            desc := strings.Trim(input.Text(), " ")

            t := Task{}
            t.id = id
            t.title = title
            t.description = desc

            tasks = append(tasks, t)

            id += 1
        }

        default:
            message = fmt.Sprintf("Unknown command `%s`", cmd)
        }
    }
}
