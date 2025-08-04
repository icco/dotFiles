package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// updateOhMyZsh clones the latest oh-my-zsh and copies everything except .git and custom
func updateOhMyZsh() error {
	// Create temporary directory for cloning
	tmpDir, err := os.MkdirTemp("", "ohmyzsh-update-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	cloneDir := filepath.Join(tmpDir, "ohmyzsh")
	log.Printf("Cloning oh-my-zsh to %s...\n", cloneDir)

	// Clone oh-my-zsh repository
	cmd := exec.Command("git", "clone", "https://github.com/ohmyzsh/ohmyzsh.git", cloneDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to clone oh-my-zsh: %s, %w", string(output), err)
	}

	// Target directory
	targetDir := "link/oh-my-zsh"

	// Backup custom directory if it exists
	customDir := filepath.Join(targetDir, "custom")
	customBackup := ""
	if _, err := os.Stat(customDir); err == nil {
		customBackup = filepath.Join(tmpDir, "custom-backup")
		log.Printf("Backing up custom directory to %s...\n", customBackup)
		if err := copyDir(customDir, customBackup); err != nil {
			return fmt.Errorf("failed to backup custom directory: %w", err)
		}
	}

	// Remove existing oh-my-zsh directory (except custom)
	if err := os.RemoveAll(targetDir); err != nil {
		return fmt.Errorf("failed to remove existing oh-my-zsh directory: %w", err)
	}

	// Create target directory
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	// Copy everything from cloned repo except .git and custom
	entries, err := os.ReadDir(cloneDir)
	if err != nil {
		return fmt.Errorf("failed to read cloned directory: %w", err)
	}

	for _, entry := range entries {
		if entry.Name() == ".git" || entry.Name() == "custom" {
			log.Printf("Skipping %s...\n", entry.Name())
			continue
		}

		srcPath := filepath.Join(cloneDir, entry.Name())
		dstPath := filepath.Join(targetDir, entry.Name())

		log.Printf("Copying %s...\n", entry.Name())
		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return fmt.Errorf("failed to copy directory %s: %w", entry.Name(), err)
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return fmt.Errorf("failed to copy file %s: %w", entry.Name(), err)
			}
		}
	}

	// Restore custom directory if it was backed up
	if customBackup != "" {
		log.Printf("Restoring custom directory from backup...\n")
		if err := copyDir(customBackup, customDir); err != nil {
			return fmt.Errorf("failed to restore custom directory: %w", err)
		}
	}

	// Commit the changes
	cmd = exec.Command("git", "add", targetDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to add oh-my-zsh to git: %s, %w", string(output), err)
	}

	cmd = exec.Command("git", "commit", "-m", "oh-my-zsh update")
	if output, err := cmd.CombinedOutput(); err != nil {
		// Check if the error is due to no changes to commit
		if strings.Contains(string(output), "nothing to commit") ||
			strings.Contains(string(output), "nothing added to commit") {
			log.Println("No changes to commit - oh-my-zsh was already up to date")
		} else {
			return fmt.Errorf("failed to commit oh-my-zsh changes: %s, %w", string(output), err)
		}
	}

	return nil
}

// copyFile copies a single file from src to dst
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Get source file info for permissions
	srcInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	// Set the same permissions as source
	return os.Chmod(dst, srcInfo.Mode())
}

// copyDir recursively copies a directory from src to dst
func copyDir(src, dst string) error {
	// Get source directory info
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Create destination directory
	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}