package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// updateOhMyZsh clones the latest oh-my-zsh and copies everything except .git and custom
func updateOhMyZsh() error {
	tmpDir, err := os.MkdirTemp("", "ohmyzsh-update-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	cloneDir := filepath.Join(tmpDir, "ohmyzsh")
	log.Printf("Cloning oh-my-zsh to %s...\n", cloneDir)

	if out, err := exec.Command("git", "clone", "https://github.com/ohmyzsh/ohmyzsh.git", cloneDir).CombinedOutput(); err != nil {
		return fmt.Errorf("failed to clone oh-my-zsh: %s: %w", string(out), err)
	}

	targetDir := "link/oh-my-zsh"

	// Preserve the local custom/ directory across the wipe-and-recopy below.
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
	if err := os.MkdirAll(targetDir, 0755); err != nil {
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

	// Strip the upstream `custom/` rule from .gitignore — we commit custom/.
	if err := stripCustomFromGitignore(filepath.Join(targetDir, ".gitignore")); err != nil {
		return fmt.Errorf("failed to strip custom from .gitignore: %w", err)
	}

	if out, err := exec.Command("git", "add", targetDir).CombinedOutput(); err != nil {
		return fmt.Errorf("failed to add oh-my-zsh to git: %s: %w", string(out), err)
	}

	return runGit("commit", "-m", "chore: oh-my-zsh update")
}

func stripCustomFromGitignore(path string) error {
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

	return os.WriteFile(path, []byte(strings.Join(out, "\n")), 0644)
}
