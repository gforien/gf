package gf

import (
	"fmt"

	"github.com/gforien/gf/internal/fzf"
	"github.com/spf13/cobra"
)

var fzfCmd = &cobra.Command{
	Use: "fzf",
}

var fzfPlanets = &cobra.Command{
	Use: "planets",
	Run: func(cmd *cobra.Command, args []string) {
		res := fzf.Popup(
			[]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"},
			&fzf.Options{})
		fmt.Print(res)
	},
}

func init() {
	fzfCmd.AddCommand(fzfPlanets)
	RootCmd.AddCommand(fzfCmd)
}
