package viper

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

// UnmarshalStringEnv takes a single key, expected to contain a string.
// It differs from UnmarshalKey because the environment variables
// contained in this string are evaluated at runtime with sh -c
//
// Example:
//
//	cfgPath: $XDG_CONFIG_HOME/fzf/config.yaml
//
//	UnmarshalKey("cfgPath")       -> $XDG_CONFIG_HOME/fzf/config.yaml
//	UnmarshalStringEnv("cfgPath") -> /home/foo/.config/fzf/config.yaml
func UnmarshalStringEnv(key string) string {
	var str string
	err := viper.UnmarshalKey(key, &str)
	if err != nil {
		log.Default().Panic(err)
	}

	cmd := exec.Command(
		"sh",
		"-c",
		fmt.Sprintf("echo \"%s\"", str))
	outBytes, err := cmd.Output()
	if err != nil {
		log.Default().Panic(err)
	}

	outStr := (string)(outBytes)
	outStr = strings.TrimSpace(outStr)
	return outStr
}

// UnmarshalRegexArray retrieves an array of strings from the provided key,
// concatenates them into a single regular expression pattern, and compiles it.
//
// The strings in the array are joined with a pipe (|) to form an OR expression
// (e.g., "a|b|c"), and the resulting pattern is enclosed in parentheses to create
// a group. The final regular expression is then compiled and returned.
//
// If there is an error during unmarshalling, the function will panic.
func UnmarshalRegexArray(key string) *regexp.Regexp {
	var patterns []string
	err := viper.UnmarshalKey(key, &patterns)
	if err != nil {
		log.Default().Panic(err)
	}

	if len(patterns) == 0 {
		return regexp.MustCompile("a^") // impossible pattern
	}

	concat := "(" + strings.Join(patterns, "|")
	concat = strings.TrimSuffix(concat, "|") + ")"

	regex := regexp.MustCompile(concat)

	return regex
}
