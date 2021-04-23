package cmd

import (
    "log"
    "os/exec"
    "github.com/spf13/cobra"
)

func Sound() {
    cmd := exec.Command("afplay", "/System/Library/Sounds/Blow.aiff")
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}

var rootCmd = &cobra.Command {
        Use: "sound",
        Run: func(cmd *cobra.Command, args []string) {
                Sound()
            },
}

func Execute() {
    rootCmd.Execute()
}
