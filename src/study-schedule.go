package main

import (
    "fmt"
    "os/exec"
    "os"
)

func main() {
    exit := false

    for !exit {
        clearScreen()
        fmt.Println("hello, world")
    }
}

func clearScreen() {
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
}
