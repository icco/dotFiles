package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
)

// gitCommitAll stages all tracked changes and commits, no-op if the tree is clean.
func gitCommitAll(ctx context.Context, msg string) error {
	clean, err := gitTreeClean(ctx)
	if err != nil {
		return err
	}
	if clean {
		log.Printf("git: nothing to commit for %q\n", msg)
		return nil
	}
	return runGit(ctx, "commit", "-a", "-m", msg)
}

// gitTreeClean reports whether the working tree has no staged or unstaged changes.
func gitTreeClean(ctx context.Context) (bool, error) {
	out, err := exec.CommandContext(ctx, "git", "status", "--porcelain").Output()
	if err != nil {
		return false, fmt.Errorf("git status: %w", err)
	}
	return len(out) == 0, nil
}

// runGit invokes git with the given arguments, returning a wrapped error that
// includes combined stdout+stderr on failure.
func runGit(ctx context.Context, args ...string) error {
	// #nosec G204 -- args are internal, not user-supplied; only "git" is invoked.
	out, err := exec.CommandContext(ctx, "git", args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("git %v: %s: %w", args, out, err)
	}
	return nil
}
