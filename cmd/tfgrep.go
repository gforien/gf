package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/gforien/gf/internal/viper"
	"github.com/spf13/cobra"
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
func processLines(inputChan <-chan string, outputChan chan<- string, dotRegex *regexp.Regexp, hideRegex *regexp.Regexp) {
	isDot := false
	isEmpty := false

	for line := range inputChan {
		if hideRegex.MatchString(line) {
			isEmpty = true
			continue
		} else if isEmpty {
			isEmpty = false
			continue
		} else if dotRegex.MatchString(line) {
			outputChan <- "."
			isDot = true
		} else if isDot {
			outputChan <- "\n" + line + "\n"
			isDot = false
		} else {
			outputChan <- line + "\n"
		}
		isEmpty = false
	}
	close(outputChan)
}

// Function to write lines from the output channel to stdout
func writeLines(outputChan <-chan string) {
	for line := range outputChan {
		fmt.Print(line)
	}
}

func TfGrep(cmd *cobra.Command, args []string) {
	dotRegex := viper.UnmarshalRegexArray("tfgrep.dot_patterns")
	hideRegex := viper.UnmarshalRegexArray("tfgrep.hide_patterns")
	inputChan := make(chan string)
	outputChan := make(chan string)

	go readLines(inputChan)
	go processLines(inputChan, outputChan, dotRegex, hideRegex)
	writeLines(outputChan)
}

func init() {
	rootCmd.AddCommand(tfgrepCmd)
}
