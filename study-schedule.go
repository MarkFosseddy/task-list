package main

import (
    "fmt"
    "os/exec"
    "os"
    "strings"
    "time"
    "bufio"
)

type Task struct {
    id int
    title string
    desc string
}

var exit bool = false
var isDeleting bool = false
var message string = ""
var tasks []Task = []Task{
    Task{id: 0, title: "Hello, World!", desc: "How are you?"},
    Task{id: 1, title: "Second Title", desc: ""},
}

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func update() {
    message = ""

    scanner.Scan()
    cmd := strings.ToLower(strings.Trim(scanner.Text(), " "))
    if len(cmd) == 0 { return }

    switch cmd {
    case "add": {
        message = "Enter task title (empty for cancel):\n"
        draw()
        scanner.Scan()
        title := strings.Trim(scanner.Text(), " ")

        if len(title) > 0 {
            message = "Enter task description (optional):\n"
            draw()
            scanner.Scan()
            desc := strings.Trim(scanner.Text(), " ")

            id := time.Now().Nanosecond()
            newTask := Task{id, title, desc}
            tasks = append(tasks, newTask)
        }

        message = ""
    }

    case "delete":
        message = "TODO: Delete Item\n"

    case "q", "quit", "exit":
        exit = true
        clearScreen()

    default:
        message = fmt.Sprintf("Unknown command: `%s`\n", cmd)
    }
}

func draw() {
    clearScreen()

    drawText("┏━━ Study Schedule ━━━\n")
    drawEmptyLine()

    if len(tasks) > 0 {
        drawText("  List:\n")
        drawList(len(tasks), func (i int) {
            item := tasks[i]

            if isDeleting {
                drawText(fmt.Sprintf("    [%v] %v\n", i + 1, item.title))
            } else {
                drawText(fmt.Sprintf("    %v\n", item.title))
            }

            if (len(item.desc) > 0) {
                drawText("      " + item.desc + "\n")
            }
            drawEmptyLine()
        })
    }

    if len(message) > 0 {
        drawText(message)
    } else {
        drawEmptyLine()
    }

    drawText(">> ")
}

// UI
func clearScreen() {
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
}

func drawText(s string) {
    fmt.Print(s)
}

func drawEmptyLine() {
    drawText("\n")
}

func drawList(listLen int, builder func(i int)) {
    for i := 0; i < listLen; i += 1 {
        builder(i)
    }
}

func main() {
    for !exit {
        draw()
        update()
    }
}
