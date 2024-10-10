package viper

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

// UnmarshalStringEnv takes a single key, expected to be a string.
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
