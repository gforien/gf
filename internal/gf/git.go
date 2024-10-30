package gf

import (
	"github.com/gforien/gf/internal/git"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "A brief description of your command",
}

var gitClone = &cobra.Command{
	Use:   "clone",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		git.Clone(args[0])
	},
}

func init() {
	gitCmd.AddCommand(gitClone)
	RootCmd.AddCommand(gitCmd)
}
