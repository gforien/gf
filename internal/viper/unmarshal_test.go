package viper

import (
	"fmt"
	"os"
	"regexp"
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

func TestUnmarshalRegexArray(t *testing.T) {
	type TestMatch struct {
		given    string
		expected bool
	}

	// Define test cases
	tests := []struct {
		name          string
		viperKey      string
		viperValue    interface{}
		expectedRegex *regexp.Regexp
		testMatch     []TestMatch
	}{
		{
			name:          "Success",
			viperKey:      "test_key",
			viperValue:    []string{"abc", "def", "ghi"},
			expectedRegex: regexp.MustCompile("(abc|def|ghi)"),
			testMatch: []TestMatch{
				{"abc", true},
				{"def", true},
				{"ghi", true},
				{"xyz", false},
			},
		},
		{
			name:          "Empty array",
			viperKey:      "empty_key",
			viperValue:    []string{},
			expectedRegex: regexp.MustCompile("a^"),
			testMatch: []TestMatch{
				{"any input", false},
			},
		},
	}

	// Run each test case
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Set Viper configuration
			if test.viperValue != nil {
				viper.Set(test.viperKey, test.viperValue)
			}

			// Call the function and validate the results
			regex := UnmarshalRegexArray(test.viperKey)

			// Validate that the compiled regex matches as expected
			for _, matchTest := range test.testMatch {
				if result := regex.MatchString(matchTest.given); result != matchTest.expected {
					t.Errorf("Test '%s': expected regex match for '%s' to be %v, but got %v", test.name, matchTest.given, matchTest.expected, result)
				}
			}
		})
	}
}
