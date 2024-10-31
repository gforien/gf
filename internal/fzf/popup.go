package fzf

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Options []string

var defaultOpts = Options{
	"--bind=ctrl-l:replace-query",
	"--bind=ctrl-c:abort",
	"-p", "50%",
}

func Popup(args []string, addOpts *Options) string {
	cmd, stdin := initPopup(addOpts)

	go feedArgs(stdin, args)

	res, err := cmd.CombinedOutput()
	if err != nil {
		log.Default().Fatalln("Error when running fzf:", err)
	}
	return string(res)
}

// PopupFile is a duplicate of Popup but reading a file instead of []strings
func PopupFile(path string, addOpts *Options) string {
	cmd, stdin := initPopup(addOpts)

	go feedFile(stdin, path)

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

// feedArgs reads a slice of strings and writes them to stdin
func feedArgs(stdin io.WriteCloser, args []string) {
	defer stdin.Close()
	for _, s := range args {
		if _, err := io.WriteString(stdin, s+"\n"); err != nil {
			fmt.Println("Error writing to stdin:", err)
			return
		}
	}
}

// feedArgs reads a file and writes it to stdin
func feedFile(stdin io.WriteCloser, path string) {
	defer stdin.Close()

	f, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		log.Default().Fatalln("Error opening file:", err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		if _, err := io.WriteString(stdin, line+"\n"); err != nil {
			log.Default().Fatalln("Error writing to stdin:", err)
			return
		}
	}

	if err := scan.Err(); err != nil {
		log.Default().Fatalln("Error reading file:", err)
	}
}
