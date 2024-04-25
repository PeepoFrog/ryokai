package git

import (
	"os"

	"github.com/go-git/go-git/v5" // Import go-git
)

func CloneRepo(url, path string) error {
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:           url,
		Progress:      os.Stdout,
		Depth:         1, // Depth 1 means only clone the latest history
		ReferenceName: "feature/joiner",
	})
	if err != nil {
		return err
	}
	// b, err := r.Branch("joiner")
	// if err != nil {
	// 	return err
	// }
	// log.Println(b)
	return nil
}
