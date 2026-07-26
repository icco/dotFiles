package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func updateOhMyZsh(ctx context.Context) error {
	tmpDir, err := os.MkdirTemp("", "ohmyzsh-update-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	cloneDir := filepath.Join(tmpDir, "ohmyzsh")
	log.Printf("Cloning oh-my-zsh to %s...\n", cloneDir)

	// #nosec G204 -- cloneDir is a program-controlled temp path; clone URL is hardcoded.
	if out, err := exec.CommandContext(ctx, "git", "clone", "https://github.com/ohmyzsh/ohmyzsh.git", cloneDir).CombinedOutput(); err != nil {
		return fmt.Errorf("failed to clone oh-my-zsh: %s: %w", string(out), err)
	}

	targetDir := "link/oh-my-zsh"

	// custom/ holds local overrides — preserve across the wipe below.
	customDir := filepath.Join(targetDir, "custom")
	customBackup := ""
	if _, err := os.Stat(customDir); err == nil {
		customBackup = filepath.Join(tmpDir, "custom-backup")
		log.Printf("Backing up custom directory to %s...\n", customBackup)
		if err := copyDir(customDir, customBackup); err != nil {
			return fmt.Errorf("failed to backup custom directory: %w", err)
		}
	}

	if err := os.RemoveAll(targetDir); err != nil {
		return fmt.Errorf("failed to remove existing oh-my-zsh directory: %w", err)
	}
	if err := os.MkdirAll(targetDir, 0750); err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	entries, err := os.ReadDir(cloneDir)
	if err != nil {
		return fmt.Errorf("failed to read cloned directory: %w", err)
	}
	for _, entry := range entries {
		name := entry.Name()
		if name == ".git" || name == "custom" {
			log.Printf("Skipping %s...\n", name)
			continue
		}

		srcPath := filepath.Join(cloneDir, name)
		dstPath := filepath.Join(targetDir, name)

		log.Printf("Copying %s...\n", name)
		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return fmt.Errorf("failed to copy directory %s: %w", name, err)
			}
			continue
		}
		if err := copyFile(srcPath, dstPath); err != nil {
			return fmt.Errorf("failed to copy file %s: %w", name, err)
		}
	}

	if customBackup != "" {
		log.Printf("Restoring custom directory from backup...\n")
		if err := copyDir(customBackup, customDir); err != nil {
			return fmt.Errorf("failed to restore custom directory: %w", err)
		}
	}

	// Upstream ignores custom/; we commit it, so strip the rule.
	if err := stripCustomFromGitignore(filepath.Join(targetDir, ".gitignore")); err != nil {
		return fmt.Errorf("failed to strip custom from .gitignore: %w", err)
	}

	if err := runGit(ctx, "add", targetDir); err != nil {
		return fmt.Errorf("failed to add oh-my-zsh to git: %w", err)
	}

	return gitCommitAll(ctx, "chore: oh-my-zsh update")
}

func stripCustomFromGitignore(path string) error {
	// #nosec G304 -- path is a program-controlled location inside link/oh-my-zsh.
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	lines := strings.Split(string(data), "\n")
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "# custom files" || trimmed == "custom/" {
			continue
		}
		out = append(out, line)
	}

	// #nosec G306,G703 -- .gitignore is a user-facing config at a program-controlled path; 0644 is the standard mode.
	return os.WriteFile(path, []byte(strings.Join(out, "\n")), 0644)
}
