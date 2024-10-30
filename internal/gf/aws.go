package gf

import (
	"github.com/spf13/cobra"
)

// awsCmd represents the aws command
var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "AWS-related commands",
}

func init() {
	RootCmd.AddCommand(awsCmd)
}
