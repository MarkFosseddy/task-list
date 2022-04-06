package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    exit := false
    message := ""

    input := bufio.NewScanner(os.Stdin)

    for !exit {
        fmt.Print("\033c")

        fmt.Println("Task List")
        fmt.Println()

        if len(message) > 0 {
            fmt.Println(message)
        } else {
            fmt.Println()
        }

        fmt.Print(">> ")

        message = ""

        input.Scan()
        cmd := strings.Trim(input.Text(), " ")

        if len(cmd) == 0 { continue }

        switch cmd {
        case "q":
            exit = true

        default:
            message = fmt.Sprintf("Unknown command `%s`", cmd)
        }
    }
}
