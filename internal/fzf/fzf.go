package fzf

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

type Options []string

var defaultOpts = Options{"-p", "50%"}

func Popup(args []string, addOpts *Options) string {
	cmd, stdin := initPopup(addOpts)

	go feedArgs(stdin, args)

	res, err := cmd.CombinedOutput()
	if err != nil {
		log.Default().Fatalln("Error when running fzf:", err)
	}
	return string(res)
}

func initPopup(opts *Options) (*exec.Cmd, io.WriteCloser) {
	if opts == nil {
		log.Default().Fatalln("Error: opts is nil")
	}
	for _, o := range *opts {
		defaultOpts = append(defaultOpts, o)
	}
	cmd := exec.Command("fzf-tmux", defaultOpts...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Default().Fatalln("Error creating stdin pipe: ", err)
		return nil, nil
	}

	return cmd, stdin
}

func feedArgs(stdin io.WriteCloser, args []string) {
	defer stdin.Close()
	for _, s := range args {
		if _, err := io.WriteString(stdin, s+"\n"); err != nil {
			fmt.Println("Error writing to stdin:", err)
			return
		}
	}
}
