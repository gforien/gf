package cmd

import (
	"github.com/gforien/gf/internal/fzf"
	"github.com/spf13/cobra"
)

var fzfCmd = &cobra.Command{
	Use: "fzf",
}

var fzfPlanets = &cobra.Command{
	Use: "planets",
	Run: func(cmd *cobra.Command, args []string) {
		fzf.Popup([]string{"a", "b", "c", "d"})
	},
}

func init() {
	fzfCmd.AddCommand(fzfPlanets)
	rootCmd.AddCommand(fzfCmd)
}
