package viper

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalStringEnv(t *testing.T) {
	viperKey := "cfgPath"

	// Table of test cases
	tests := []struct {
		given       string
		expected    string
		description string
	}{
		{
			given:    "$HOME/.config/fzf/cfg.yaml",
			expected: os.Getenv("HOME") + "/.config/fzf/cfg.yaml",
		},
		{
			description: "$UNDEFINED_VAR in string",
			given:       "$UNDEFINED_VAR/.config/fzf/cfg.yaml",
			expected:    os.Getenv("UNDEFINED_VAR") + "/.config/fzf/cfg.yaml",
		},
		{
			description: "Multiple $VAR in string",
			given:       "Running $SHELL with $LANG",
			expected:    fmt.Sprintf("Running %s with %s", os.Getenv("SHELL"), os.Getenv("LANG")),
		},
		{
			description: "No $ENV_VAR in string",
			given:       "/usr/local/etc/cfg.yaml",
			expected:    "/usr/local/etc/cfg.yaml",
		},
	}

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.given, func(t *testing.T) {
			// Set up Viper configuration for the current test case
			viper.Set(viperKey, tt.given)

			// Test UnmarshalStringEnv
			result := UnmarshalStringEnv(viperKey)

			// Check that the result matches the expected value
			assert.Equal(t, tt.expected, result, "The evaluated path should match the expected result.")
		})
	}
}
