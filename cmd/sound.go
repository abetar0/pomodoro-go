package main

import (
    "log"
    "os/exec"
)

func main() {
    cmd := exec.Command("afplay", "/System/Library/Sounds/Blow.aiff")
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}
