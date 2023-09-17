package internal

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
)

// Checks if a git repository exists at the path, if not, it will initialize a new one.
func initGitRepo(path string) error {
	_, err := git.PlainOpen(path)
	switch {
	case errors.Is(err, git.ErrRepositoryNotExists):
		fmt.Println("Initialising a new git repository...")
		_, err = git.PlainInit(path, false)
		if err != nil {
			return fmt.Errorf("Failed to initialize git repository: %w", err)
		}
		fmt.Println("Git repository initialized")
	case err == nil:
		fmt.Println("Git repository already exists")
	default:
		return fmt.Errorf("Failed to open git repository: %w", err)
	}
	return nil
}
