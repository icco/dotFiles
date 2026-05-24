package main

import (
	"fmt"
	"log"
	"os/exec"
)

// gitCommitAll stages all tracked changes and commits, no-op if the tree is clean.
func gitCommitAll(msg string) error {
	clean, err := gitTreeClean()
	if err != nil {
		return err
	}
	if clean {
		log.Printf("git: nothing to commit for %q\n", msg)
		return nil
	}
	return runGit("commit", "-a", "-m", msg)
}

// gitTreeClean reports whether the working tree has no staged or unstaged changes.
func gitTreeClean() (bool, error) {
	out, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return false, fmt.Errorf("git status: %w", err)
	}
	return len(out) == 0, nil
}

func runGit(args ...string) error {
	out, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("git %v: %s: %w", args, out, err)
	}
	return nil
}
