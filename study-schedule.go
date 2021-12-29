package main

import (
    "fmt"
    "os/exec"
    "os"
    "strings"
)

var exit bool = false
var message string = ""
var tasks []Task = []Task{
    Task{id: 0, title: "Hello, World!", desc: "How are you?"},
    Task{id: 1, title: "Second Title", desc: ""},
}

func main() {
    for !exit {
        draw()
        update()
    }
}

type Task struct {
    id int
    title string
    desc string
}

func update() {
    message = ""

    cmd := ""
    fmt.Scanln(&cmd)
    if len(cmd) == 0 {
        return
    }
    cmd = strings.ToLower(cmd)

    switch cmd {
    case "add", "ad", "a":
        message = "TODO: Add new Item\n"

    case "delete", "del", "d":
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
            drawText("    " + item.title + "\n")
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
