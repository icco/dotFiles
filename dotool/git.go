package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func gitCommitAll(msg string) error {
	return runGit("commit", "-a", "-m", msg)
}

// runGit treats "nothing to commit" as success so callers can be idempotent.
func runGit(args ...string) error {
	out, err := exec.Command("git", args...).CombinedOutput()
	if err == nil {
		return nil
	}
	s := string(out)
	if strings.Contains(s, "nothing to commit") || strings.Contains(s, "nothing added to commit") {
		log.Printf("git %s: nothing to commit\n", strings.Join(args, " "))
		return nil
	}
	return fmt.Errorf("git %s failed: %s: %w", strings.Join(args, " "), s, err)
}
