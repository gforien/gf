package gf

import (
	"errors"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/gforien/gf/internal/viper"
	"github.com/spf13/cobra"
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

	tplPath := viper.UnmarshalStringEnv("gomn.template_path")

	fileBytes, err := os.ReadFile(tplPath)
	checkErr(err)
	fileStr := (string)(fileBytes)

	tpls, err := template.New("titleTest").Funcs(funcMap).Parse(fileStr)
	checkErr(err)

	outDir := viper.UnmarshalStringEnv("gomn.output_dir")

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
	if err != nil {
		return &ErrFileExists{fileName}
	}
	defer file.Close()

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

func init() {
	RootCmd.AddCommand(gomnCmd)
}
