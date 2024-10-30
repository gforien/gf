package gf

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var ecyCmd = &cobra.Command{
	Use:   "ecy",
	Short: "Exa configuration generator from YAML",
	Long:  `Usage: gf ecy`,
	Run:   Ecy,
}

var EcyPath = fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".dotfiles/common/config/eza/colors.yaml")

type Data struct {
	Entries []Entry `yaml:"exacolors"`
}

type Entry struct {
	Ansi        string   `yaml:"ansi"`
	Identifiers []string `yaml:"iden"`
}

func Ecy(cmd *cobra.Command, args []string) {
	var ecyFile []byte
	ecyFile, err := os.ReadFile(EcyPath)
	if err != nil {
		log.Fatalf("error '%s' when reading ecy config file '%s'", err, EcyPath)
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

func init() {
	RootCmd.AddCommand(ecyCmd)
}
