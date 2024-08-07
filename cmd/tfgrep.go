package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tfgrepCmd represents the tfgrep command
var tfgrepCmd = &cobra.Command{
	Use:   "tfgrep",
	Short: "gf tfgrep",
	Long:  `Usage: gf tfgrep`,
	Run:   TfGrep,
}

// Function to read lines from stdin and send them to a channel
func readLines(inputChan chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputChan <- scanner.Text()
	}
	close(inputChan)
}

// Function to process lines from the input channel and send results to the output channel
func processLines(inputChan <-chan string, outputChan chan<- string, re *regexp.Regexp) {
	isDot := false

	for line := range inputChan {
		if re.MatchString(line) {
			outputChan <- "."
			isDot = true
		} else if isDot {
			outputChan <- "\n" + line + "\n"
			isDot = false
		} else {
			outputChan <- line + "\n"
		}
	}
	close(outputChan)
}

// Function to write lines from the output channel to stdout
func writeLines(outputChan <-chan string) {
	for line := range outputChan {
		fmt.Print(line)
	}
}

// readConfig reads config file $HOME/.gf.yaml and returns
// a regex matching any of the defined patterns.
func readConfig() *regexp.Regexp {
	patterns := []string{"Reading...", "Read complete"}
	err := viper.UnmarshalKey("tfgrep.patterns", &patterns)
	cobra.CheckErr(err)

	pattern := strings.Join(patterns, "|")
	re := regexp.MustCompile(pattern)

	return re
}

func TfGrep(cmd *cobra.Command, args []string) {
	re := readConfig()
	inputChan := make(chan string)
	outputChan := make(chan string)

	go readLines(inputChan)
	go processLines(inputChan, outputChan, re)
	writeLines(outputChan)
}

func init() {
	rootCmd.AddCommand(tfgrepCmd)
}
