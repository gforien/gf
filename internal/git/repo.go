package git

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Repo struct to store parsed repository information
type Repo struct {
	httpUrl string // eg "https://github.com/go-git/go-git.git/"
	path    string // eg "$WORK/github.com/go-git/go-git/"
	host    string // eg "github.com"
	user    string // eg "go-git"
	name    string // eg "go-git"
}

var Work = os.Getenv("WORK")

var (
	httpRegex            = regexp.MustCompile(`^https://([-.[:alnum:]]+)/([-_.[:alnum:]]+)/([-_.[:alnum:]]+)(/pull/([[:digit:]]+))?/?.*$`)
	sshRegex             = regexp.MustCompile(`^git@([-.[:alnum:]]+):([-_.[:alnum:]]+)/([-_.[:alnum:]]+)$`)
	localPathRegex       = regexp.MustCompile(`^WORK/([^/]+)/([^/]+)/([^/#]+)(#([[:digit:]]+))?/?$`)
	hostOwnerRepoPRRegex = regexp.MustCompile(`^([^/]+)/([^/]+)/([^/#]+)(#([[:digit:]]+))?/?$`)
	ownerRepoPRRegex     = regexp.MustCompile(`^([^/]+)/([^/]+)(#([[:digit:]]+))?/?$`)
)

func FromUrl(input string) Repo {
	var repo Repo

	if httpRegex.MatchString(input) {
		matches := httpRegex.FindStringSubmatch(input)
		repo.host = matches[1]
		repo.user = matches[2]
		repo.name = strings.TrimSuffix(matches[3], ".git")
		repo.httpUrl = fmt.Sprintf("https://%s/%s/%s.git/", repo.host, repo.user, repo.name)
		repo.path = defaultPath(repo.host, repo.user, repo.name)
	}
	return repo
}

// TODO: FromAnything
func FromAnything(input string) Repo {
	var repo Repo

	// Attempt to match each pattern and extract relevant parts
	switch {
	case httpRegex.MatchString(input):
		matches := httpRegex.FindStringSubmatch(input)
		repo.httpUrl = input
		repo.host = matches[1]
		repo.user = matches[2]
		repo.name = strings.TrimSuffix(matches[3], ".git")
		repo.path = defaultPath(repo.host, repo.user, repo.name)
		if matches[5] != "" {
			repo.path = fmt.Sprintf("/pull/%s", matches[5])
		}

	case sshRegex.MatchString(input):
		matches := sshRegex.FindStringSubmatch(input)
		repo.host = matches[1]
		repo.user = matches[2]
		repo.name = strings.TrimSuffix(matches[3], ".git")
		repo.path = defaultPath(repo.host, repo.user, repo.name)

	case localPathRegex.MatchString(input):
		matches := localPathRegex.FindStringSubmatch(input)
		repo.host = matches[1]
		repo.user = matches[2]
		repo.name = matches[3]
		repo.path = defaultPath(repo.host, repo.user, repo.name)
		if matches[5] != "" {
			repo.path = fmt.Sprintf("#%s", matches[5])
		}

	case hostOwnerRepoPRRegex.MatchString(input):
		matches := hostOwnerRepoPRRegex.FindStringSubmatch(input)
		repo.host = matches[1]
		repo.user = matches[2]
		repo.name = matches[3]
		repo.path = defaultPath(repo.host, repo.user, repo.name)
		if matches[5] != "" {
			repo.path = fmt.Sprintf("#%s", matches[5])
		}

	case ownerRepoPRRegex.MatchString(input):
		matches := ownerRepoPRRegex.FindStringSubmatch(input)
		repo.host = "github.com"
		repo.user = matches[1]
		repo.name = matches[2]
		if matches[4] != "" {
			repo.path = fmt.Sprintf("#%s", matches[4])
		}
		repo.path = defaultPath(repo.host, repo.user, repo.name)

	default:
		fmt.Fprintf(os.Stderr, "unexpected error: failed to parse '%s'\n", input)
		os.Exit(1)
	}

	// Output result
	fmt.Printf("Repo Struct: %+v\n", repo)
	return repo
}

func defaultPath(host string, user string, name string) string {
	return fmt.Sprintf("%s/%s/%s/%s", Work, host, user, name)
}

func (r *Repo) Equal(other *Repo) bool {
	if r == nil && other == nil {
		return true
	}
	if r == nil {
		return *other == Repo{}
	}
	if other == nil {
		return *r == Repo{}
	}

	// Field-by-field comparison
	return r.httpUrl == other.httpUrl &&
		r.path == other.path &&
		r.host == other.host &&
		r.user == other.user &&
		r.name == other.name
}
