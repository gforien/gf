package cmd

import (
	"github.com/gforien/gf/internal/tmux"
	"github.com/spf13/cobra"
)

// tmuxCmd represents the tmux command
var tmuxCmd = &cobra.Command{
	Use: "tmux",
}

var tmuxNeww = &cobra.Command{
	Use: "neww",
	Run: func(cmd *cobra.Command, args []string) {
		tmux.Neww()
	},
}

func init() {
	tmuxCmd.AddCommand(tmuxNeww)
	rootCmd.AddCommand(tmuxCmd)
}
