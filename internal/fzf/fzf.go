package fzf

import (
	"fmt"
	"io"
	"os/exec"
)

type Options []string

func Popup(args []string) {
	outputChan := make(chan string)
	go func() {
		for s := range outputChan {
			fmt.Println("Got: " + s)
		}
	}()

	opts := Options{"--multi", "--reverse", "--border", "-p", "50%"}
	cmd := exec.Command("fzf-tmux", opts...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error creating stdin pipe:", err)
		return
	}

	// Goroutine to read from inputChan and write to stdin
	go func() {
		defer stdin.Close()
		for _, s := range args {
			if _, err := io.WriteString(stdin, s+"\n"); err != nil {
				fmt.Println("Error writing to stdin:", err)
				return
			}
		}
	}()

	// Start the command
	if err := cmd.Run(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}
}
