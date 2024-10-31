package semver

import (
	"fmt"
	"regexp"
	"strconv"
)

type Version struct {
	major int
	minor int
	patch int
}

var semver = regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)`)

// We expect exactly 3 matches to be found when parsing
// a provided string with the semver regexp
// If the provided string does not meet this criteria, ErrCannotParse is returned.
type ErrCannotParse struct {
	message string
}

func (e *ErrCannotParse) Error() string {
	return e.message
}

func FromString(s string) (Version, error) {
	// Regular expression to capture version numbers
	matches := semver.FindStringSubmatch(s)
	if len(matches) != 4 {
		err := &ErrCannotParse{
			message: fmt.Sprintf("Cannot parse '%s' into semver (found %d matches)", s, len(matches)),
		}
		return Version{}, err
	}

	// Convert matches to integers
	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	patch, _ := strconv.Atoi(matches[3])
	return Version{major: major, minor: minor, patch: patch}, nil
}

func (s *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch)
}

func (s *Version) ReleaseMajor() Version {
	return Version{
		major: s.major + 1,
		minor: 0,
		patch: 0,
	}
}

func (s *Version) ReleaseMinor() Version {
	return Version{
		major: s.major,
		minor: s.minor + 1,
		patch: 0,
	}
}

func (s *Version) ReleasePatch() Version {
	return Version{
		major: s.major,
		minor: s.minor,
		patch: s.patch + 1,
	}
}
