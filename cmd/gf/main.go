package main

import (
	"os"

	"github.com/gforien/gf/internal/gf"
)

func main() {
	err := gf.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
