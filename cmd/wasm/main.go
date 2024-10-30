//go:build js && wasm

package main

import (
	"bytes"
	"log"
	"strings"
	"syscall/js"

	"github.com/gforien/gf/internal/gf"
)

func main() {
	log.Default().Println("main() called ")
	js.Global().Set("runGoCode", js.FuncOf(mainJs))
	select {} // Keep the Go program running
}

// main Javascript function
func mainJs(this js.Value, jsArgs []js.Value) any {
	if jsArgs == nil || len(jsArgs) != 1 {
		return "unexpected error: jsArgs is nil or #jsArgs != 1"
	}
	args := strings.Split(jsArgs[0].String(), " ")
	log.Default().Printf(`mainJs() called with args[0] = %v`, args)

	// NOTE: discard first element if it's the top-level 'gf'
	if args[0] == "gf" {
		args = args[1:]
	}

	// capture and restore old output descriptor
	old := gf.RootCmd.OutOrStdout()
	defer gf.RootCmd.SetOut(old)

	var buf bytes.Buffer
	gf.RootCmd.SetOut(&buf)

	gf.RootCmd.SetArgs(args)
	gf.RootCmd.Execute()

	return buf.String()
}
