package gf

import (
	"fmt"
	"os"

	"github.com/gforien/gf/internal/fzf"
	"github.com/spf13/cobra"
)

var fzfCmd = &cobra.Command{
	Use: "fzf",
}

var fzfBrewPackages = &cobra.Command{
	Use:   "brew",
	Short: "Homebrew packages fuzzy-finder",
	Run: func(cmd *cobra.Command, args []string) {
		res := fzf.PopupFile(os.Getenv("FPOPUP_CACHE")+"/fbiu", &fzf.Options{})
		fmt.Print(res)
	},
}

var fzfManPages = &cobra.Command{
	Use:   "man",
	Short: "Man pages fuzzy-finder",
	Run: func(cmd *cobra.Command, args []string) {
		res := fzf.PopupFile(os.Getenv("FPOPUP_CACHE")+"/fman", &fzf.Options{})
		fmt.Print(res)
	},
}

var fzfPlanets = &cobra.Command{
	Use:   "planets",
	Short: "Planets fuzzy-finder, mostly for debugging/testing",
	Run: func(cmd *cobra.Command, args []string) {
		res := fzf.Popup(
			[]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"},
			&fzf.Options{})
		fmt.Print(res)
	},
}

func init() {
	fzfCmd.AddCommand(fzfBrewPackages)
	fzfCmd.AddCommand(fzfManPages)
	fzfCmd.AddCommand(fzfPlanets)
	RootCmd.AddCommand(fzfCmd)
}
