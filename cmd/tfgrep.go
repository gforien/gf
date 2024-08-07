package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

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
func processLines(inputChan <-chan string, outputChan chan<- string) {
	re := regexp.MustCompile(`Reading...|Read complete|Refreshing state...`)
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

func TfGrep(cmd *cobra.Command, args []string) {
	inputChan := make(chan string)
	outputChan := make(chan string)

	go readLines(inputChan)
	go processLines(inputChan, outputChan)
	writeLines(outputChan)
}

func init() {
	rootCmd.AddCommand(tfgrepCmd)
}
