package main

import (
    "fmt"
    "os/exec"
    "os"
)

func main() {
    exit := false
    message := ""

    for !exit {
        // draw
        clearScreen()

        fmt.Println("┏━━ Study Schedule ━━━");
        fmt.Println();

        fmt.Println(message);

        fmt.Print(">> ");

        // update
        message = "";

        cmd := ""
        fmt.Scanln(&cmd)
        if len(cmd) == 0 {
            continue
        }

        switch cmd {
        case "add","ad","a":
            message = "TODO: Add new Item"

        case "delete", "del", "d":
            message = "TODO: Delete Item"

        case "q", "quit", "exit":
            exit = true
            clearScreen()

        default:
            message = fmt.Sprintf("Unknown command: `%s`", cmd);
        }
    }
}

func clearScreen() {
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
}
