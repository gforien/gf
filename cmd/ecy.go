package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// ecyCmd represents the ecy command
var ecyCmd = &cobra.Command{
	Use:   "ecy",
	Short: "Exa configuration generator from YAML",
	Long: `Usage: gf ecy [command]

Available commands:
  debug     print details about ecy configuration
  read      parse configuration
  generate  generate exa configuration`,
	Args:      cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
	ValidArgs: []string{"debug", "read", "generate"},
	Run:       run,
}

var ecyPath = fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".dotfiles/common/config/exa/colors.yaml")

type Data struct {
	Entries []Entry `yaml:"exacolors"`
}
type Entry struct {
	Identifiers []string `yaml:"iden"`
	Ansi       string   `yaml:"ansi"`
}

func run(cmd *cobra.Command, args []string) {
	var ecyFile []byte
	ecyFile, err := os.ReadFile(ecyPath)
	if err != nil {
		log.Fatalf("error '%s' when reading file '%s'", err, ecyPath)
	}

	var out Data
	if err := yaml.Unmarshal(ecyFile, &out); err != nil {
		log.Fatal("error unmarshalling data: ", err)
	}
	for _, entry := range out.Entries {
		for _, identifier := range entry.Identifiers {
			fmt.Printf("%s=%s:", identifier, entry.Ansi)
		}
	}
}

// func read(in string, out []byte) error {
//   var err error
//   outFile, err := os.ReadFile(in)
//   out = outFile
//   return err
// }
//
// func parse(in byte[], out interface{}) error {
//   return yaml.Unmarshal(in, out)
// }

func init() {
	rootCmd.AddCommand(ecyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ecyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ecyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
