package cmd

import (
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
		log.Panic(err)
	}
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

	err = tpls.Execute(os.Stdout, map[string]any{
		"Date": time.Now(),
	})
	checkErr(err)

	// outDir := unmarshalStringEnv("gomn.output_dir")
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
