package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gforien/gf/internal/viper"
	"github.com/spf13/cobra"
)

// setWallpaperCmd represents the setWallpaper command
var setWallpaperCmd = &cobra.Command{
	Use:   "setWallpaper",
	Short: "Set desktop wallpaper",
	Run:   setWallpaper,
}

func setWallpaper(cobraCmd *cobra.Command, args []string) {
	wallpaperPath := viper.UnmarshalStringEnv("macos.wallpaper")
	if wallpaperPath == "" {
		wallpaperPath = os.Getenv("HOME") + "/Library/Mobile Documents/com~apple~CloudDocs/Wallpapers/Chroma-Blue-Darker-3.jpg"
	}

	cmd := exec.Command(
		"osascript",
		"-e",
		fmt.Sprintf(
			`tell application "System Events" to tell every desktop to set picture to "%s"`,
			wallpaperPath,
		),
	)
	err := cmd.Run()
	cobra.CheckErr(err)
}

func init() {
	macosCmd.AddCommand(setWallpaperCmd)
}
