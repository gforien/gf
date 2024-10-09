package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var gomnCmd = &cobra.Command{
	Use:   "gomn",
	Short: "Generate One Month of Notes",
	Long:  `Generate`,
	Run:   Gomn,
}

func checkErr(err error) {
	if err != nil {
		log.Default().Panic(err)
	}
}

type ErrFileExists struct {
	FileName string
}

func (e *ErrFileExists) Error() string {
	return fmt.Sprintf("Skipping '%s'", e.FileName)
}

func Gomn(cmd *cobra.Command, args []string) {
	fmt.Println("gomn called")

	funcMap := template.FuncMap{
		"datefmt": func(layout string, t time.Time) string {
			return t.Format(layout)
		},
		"yesterday": func(t time.Time) time.Time {
			return t.AddDate(0, 0, -1)
		},
		"tomorrow": func(t time.Time) time.Time {
			return t.AddDate(0, 0, 1)
		},
		"tolink": func(s1 string, s2 string) string {
			return fmt.Sprintf("[[%s|%s]]", s1, s2)
		},
	}

	tplPath := unmarshalStringEnv("gomn.template_path")

	fileBytes, err := os.ReadFile(tplPath)
	checkErr(err)
	fileStr := (string)(fileBytes)

	tpls, err := template.New("titleTest").Funcs(funcMap).Parse(fileStr)
	checkErr(err)

	outDir := unmarshalStringEnv("gomn.output_dir")

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	location := now.Location()
	startOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, location)

	// Iterate over all the days of the current month
	for day := startOfMonth; day.Month() == currentMonth; day = day.AddDate(0, 0, 1) {
		err := writeNote(tpls, day, outDir)
		var fileExists *ErrFileExists
		if errors.As(err, &fileExists) {
			log.Default().Println(err)
			continue
		} else if err != nil {
			log.Default().Printf("Unexpected error: %s", err)
			continue
		}
	}
}

func writeNote(tpls *template.Template, day time.Time, outDir string) error {
	fileName := day.Format("060102") + ".md"
	filePath := outDir + fileName

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	defer file.Close()
	if err != nil {
		return &ErrFileExists{fileName}
	}

	// Write content to the file (if necessary)
	err = tpls.Execute(file, map[string]any{
		"Date": day,
	})
	if err != nil {
		return err
	}
	log.Default().Printf("Templating '%s'", fileName)

	return nil
}

// unmarshalStringEnv
// Unmarshal a config string containing environment variables,
// and evaluate them.
func unmarshalStringEnv(key string) string {
	var str string
	err := viper.UnmarshalKey(key, &str)
	checkErr(err)

	cmd := exec.Command(
		"bash",
		"-c",
		fmt.Sprintf("echo %s", str))
	outBytes, err := cmd.Output()
	checkErr(err)

	outStr := (string)(outBytes)
	outStr = strings.TrimSpace(outStr)
	return outStr
}

func init() {
	rootCmd.AddCommand(gomnCmd)
}
