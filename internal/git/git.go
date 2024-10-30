package git

import (
	"log"
	"os"

	git "github.com/go-git/go-git/v5"
)

func Clone(url string) {
	// Clone the given repository to the given directory
	log.Default().Printf("git clone %s", url)

	r := FromAnything(url)

	_, err := git.PlainClone(r.path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Default().Fatal(err)
	}
}
