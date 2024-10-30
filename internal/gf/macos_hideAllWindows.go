package gf

import (
	"os/exec"

	"github.com/spf13/cobra"
)

// hideAllWindowsCmd represents the hideAllWindows command
var hideAllWindowsCmd = &cobra.Command{
	Use:   "hideAllWindows",
	Short: "Hide all windows except kitty",
	Run:   hideAllWindows,
}

func hideAllWindows(cobraCmd *cobra.Command, args []string) {
	cmd := exec.Command(
		"osascript",
		"-e",
		`tell application "System Events" to set visible of every process whose visible is true and name is not "Kitty" to false`,
	)
	err := cmd.Run()
	cobra.CheckErr(err)
}

func init() {
	macosCmd.AddCommand(hideAllWindowsCmd)
}
