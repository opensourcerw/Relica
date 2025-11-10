package git

import (
	"github.com/go-git/go-git/v6"
)

// IsGitRepo checks whether the current directory is a valid Git repo.
func IsGitRepo(path string) bool {
	_, err := git.PlainOpen(path)
	return err == nil
}
