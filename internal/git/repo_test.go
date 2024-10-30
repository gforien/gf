package git

import (
	"fmt"
	"os"
	"testing"
)

// TestFromUrl tests the FromUrl function for various input cases.
func TestFromUrl(t *testing.T) {
	// Set WORK environment variable for testing
	os.Setenv("WORK", "/mock/work")

	tests := []struct {
		name     string
		input    string
		expected Repo
	}{
		{
			name:  "Smoke test",
			input: "https://github.com/user/repo.git",
			expected: Repo{
				httpUrl: "https://github.com/user/repo.git/",
				host:    "github.com",
				user:    "user",
				name:    "repo",
				path:    Work + "/github.com/user/repo",
			},
		},
		{
			name:  "URL with trailing slash",
			input: "https://github.com/user/repo.git/",
			expected: Repo{
				httpUrl: "https://github.com/user/repo.git/",
				host:    "github.com",
				user:    "user",
				name:    "repo",
				path:    Work + "/github.com/user/repo",
			},
		},
		{
			name:  "URL without ending .git",
			input: "https://github.com/user/repo",
			expected: Repo{
				httpUrl: "https://github.com/user/repo.git/",
				host:    "github.com",
				user:    "user",
				name:    "repo",
				path:    Work + "/github.com/user/repo",
			},
		},
		{
			name:  "URL without ending .git, with trailing slash",
			input: "https://github.com/user/repo/",
			expected: Repo{
				httpUrl: "https://github.com/user/repo.git/",
				host:    "github.com",
				user:    "user",
				name:    "repo",
				path:    Work + "/github.com/user/repo",
			},
		},
		{
			name:     "Invalid URL",
			input:    "invalid-url",
			expected: Repo{}, // Expected to return a default Repo with empty fields
		},
		{
			name:     "Empty URL",
			input:    "",     // Edge case: empty string
			expected: Repo{}, // Expected to return a default Repo with empty fields
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FromUrl(test.input)

			// Compare the expected Repo with the actual result
			if !result.Equal(&test.expected) {
				t.Errorf("For input %q, expected %+v, but got %+v", test.input, test.expected, result)
			}
		})
	}
}

func TestFromAnything(t *testing.T) {
	// TODO: tests
	t.Skip("TODO")
	tests := []struct {
		input    string
		expected Repo
	}{
		// HTTPS URL with pull request
		{
			input: "https://github.com/user/repo/",
			expected: Repo{
				httpUrl: "https://github.com/user/repo/",
				host:    "github.com",
				user:    "user",
				name:    "repo",
				path:    fmt.Sprintf("%s/github.com/user/repo", Work),
			},
		},
		// HTTPS URL without pull request
		{
			input: "https://github.com/user/repo.git",
			expected: Repo{
				httpUrl: "https://github.com/user/repo.git",
				host:    "github.com",
				user:    "user",
				name:    "repo",
				path:    fmt.Sprintf("%s/github.com/user/repo", Work),
			},
		},
		// SSH URL
		{
			input: "git@github.com:user/repo.git",
			expected: Repo{
				host: "github.com",
				user: "user",
				name: "repo",
				path: fmt.Sprintf("%s/github.com/user/repo", Work),
			},
		},
		// Local path
		{
			input: "WORK/host/user/repo",
			expected: Repo{
				host: "host",
				user: "user",
				name: "repo",
				path: fmt.Sprintf("%s/github.com/user/repo", Work),
			},
		},
		// Host/Owner/Repo#PR
		{
			input: "host/user/repo",
			expected: Repo{
				host: "host",
				user: "user",
				name: "repo",
				path: fmt.Sprintf("%s/github.com/user/repo", Work),
			},
		},
		// Owner/Repo#PR with default host based on hostname
		{
			input: "user/repo#1011",
			expected: Repo{
				host: "github.com",
				user: "user",
				name: "repo",
				path: "#1011",
			},
		},
	}

	// Run each test
	for _, test := range tests {
		result := FromAnything(test.input)

		// Compare expected and actual results
		if !result.Equal(&test.expected) {
			t.Errorf("For input '%s', expected %+v, but got %+v", test.input, test.expected, result)
		}
	}
}
