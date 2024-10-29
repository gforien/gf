package tmux

import (
	"os/exec"

	"github.com/spf13/cobra"
)

func Neww() {
	cmd := exec.Command("tmux", "neww")
	err := cmd.Run()
	cobra.CheckErr(err)
}
