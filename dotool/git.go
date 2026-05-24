package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// gitCommitAll runs `git commit -a -m <msg>`. A "nothing to commit" result is
// not an error — the caller wanted an idempotent sync.
func gitCommitAll(msg string) error {
	return runGit("commit", "-a", "-m", msg)
}

// runGit runs a git command, treating "nothing to commit" output as success.
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
