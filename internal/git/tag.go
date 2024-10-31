package git

import (
	"log"

	"github.com/gforien/gf/internal/semver"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// GetVersion opens repository at current path, and return the latest tag
func GetVersion() semver.Version {
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("Error opening repository: %v", err)
	}

	tags, err := repo.Tags()
	if err != nil {
		log.Fatalf("Error fetching tags: %v", err)
	}

	var latestTag string
	err = tags.ForEach(func(ref *plumbing.Reference) error {
		latestTag = ref.Name().Short()
		return nil
	})
	if err != nil {
		log.Fatalf("Error iterating tags: %v", err)
	}

	s, err := semver.FromString(latestTag)
	if err != nil {
		log.Fatalf("Error parsing latest tag: %v", err)
	}

	return s
}
