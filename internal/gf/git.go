package gf

import (
	"fmt"

	"github.com/gforien/gf/internal/git"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "git-related commands",
}

var gitClone = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		git.Clone(args[0])
	},
}

var gitGetVersion = &cobra.Command{
	Use:   "getVersion",
	Short: "Get the current version, based on git tags",
	Run: func(cmd *cobra.Command, args []string) {
		current := git.GetVersion()
		fmt.Println(current.String())
	},
}

var gitReleaseMajor = &cobra.Command{
	Use:   "releaseMajor",
	Short: "Compute the next major release, based on git tags",
	Run: func(cmd *cobra.Command, args []string) {
		current := git.GetVersion()
		next := current.ReleaseMajor()
		fmt.Println(next.String())
	},
}

var gitReleaseMinor = &cobra.Command{
	Use:   "releaseMinor",
	Short: "Compute the next minor release, based on git tags",
	Run: func(cmd *cobra.Command, args []string) {
		current := git.GetVersion()
		next := current.ReleaseMinor()
		fmt.Println(next.String())
	},
}

var gitReleasePatch = &cobra.Command{
	Use:   "releasePatch",
	Short: "Compute the next patch release, based on git tags",
	Run: func(cmd *cobra.Command, args []string) {
		current := git.GetVersion()
		next := current.ReleasePatch()
		fmt.Println(next.String())
	},
}

func init() {
	gitCmd.AddCommand(gitClone)
	gitCmd.AddCommand(gitGetVersion)
	gitCmd.AddCommand(gitReleaseMajor)
	gitCmd.AddCommand(gitReleaseMinor)
	gitCmd.AddCommand(gitReleasePatch)
	RootCmd.AddCommand(gitCmd)
}
