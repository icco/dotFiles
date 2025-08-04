package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// buildStructure creates the required directory structure in the home directory
func buildStructure() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	dirs := []string{"Projects", "bin", "tmp"}
	for _, dir := range dirs {
		dirPath := filepath.Join(homeDir, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			log.Printf("Creating directory: %s\n", dirPath)
			if err := os.MkdirAll(dirPath, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
			}
		}
	}

	return nil
}

// linkFiles links all the dotfiles and builds directory structure
func linkFiles() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	// Build directory structure
	if err := buildStructure(); err != nil {
		return fmt.Errorf("failed to build directory structure: %w", err)
	}

	// Link files from link/ directory
	if err := linkDirectory("link", homeDir, true); err != nil {
		return fmt.Errorf("failed to link main files: %w", err)
	}

	// Link files from specific/ directory
	if err := linkSpecificFiles(homeDir); err != nil {
		return fmt.Errorf("failed to link specific files: %w", err)
	}

	// Link files from bin/ directory
	if err := linkBinFiles(homeDir); err != nil {
		return fmt.Errorf("failed to link bin files: %w", err)
	}

	return nil
}

// linkDirectory links all files in a directory to the home directory
func linkDirectory(sourceDir, homeDir string, addDot bool) error {
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", sourceDir, err)
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(sourceDir, entry.Name())
		var targetPath string

		if addDot {
			targetPath = filepath.Join(homeDir, "."+entry.Name())
		} else {
			targetPath = filepath.Join(homeDir, entry.Name())
		}

		if err := createSymlink(sourcePath, targetPath); err != nil {
			return fmt.Errorf("failed to link %s: %w", sourcePath, err)
		}
	}

	return nil
}

// linkSpecificFiles links files from the specific/ directory
func linkSpecificFiles(homeDir string) error {
	return filepath.Walk("specific", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// Remove "specific/" prefix and add dot
		relativePath, err := filepath.Rel("specific", path)
		if err != nil {
			return fmt.Errorf("failed to get relative path for %s: %w", path, err)
		}

		targetPath := filepath.Join(homeDir, "."+relativePath)
		targetDir := filepath.Dir(targetPath)

		// Create target directory if it doesn't exist
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			if err := os.MkdirAll(targetDir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", targetDir, err)
			}
		}

		return createSymlink(path, targetPath)
	})
}

// linkBinFiles links files from the bin/ directory
func linkBinFiles(homeDir string) error {
	entries, err := os.ReadDir("bin")
	if err != nil {
		return fmt.Errorf("failed to read bin directory: %w", err)
	}

	for _, entry := range entries {
		sourcePath := filepath.Join("bin", entry.Name())
		targetPath := filepath.Join(homeDir, sourcePath)

		if err := createSymlink(sourcePath, targetPath); err != nil {
			return fmt.Errorf("failed to link bin file %s: %w", sourcePath, err)
		}
	}

	return nil
}

// createSymlink creates a symbolic link, backing up existing files
func createSymlink(source, target string) error {
	// Check if target already exists
	if _, err := os.Lstat(target); err == nil {
		// Backup existing file
		backupPath := fmt.Sprintf("%s.%d.backup", target, time.Now().Unix())
		homeDir, _ := os.UserHomeDir()
		backupDir := filepath.Join(homeDir, "tmp")
		backupPath = filepath.Join(backupDir, filepath.Base(backupPath))

		// Ensure backup directory exists
		if err := os.MkdirAll(backupDir, 0755); err != nil {
			return fmt.Errorf("failed to create backup directory %s: %w", backupDir, err)
		}

		log.Printf("Backing up %s to %s\n", target, backupPath)
		if err := exec.Command("cp", "-r", target, backupPath).Run(); err != nil {
			return fmt.Errorf("failed to backup %s: %w", target, err)
		}

		// Remove existing file
		if err := os.RemoveAll(target); err != nil {
			return fmt.Errorf("failed to remove existing %s: %w", target, err)
		}
	}

	// Get absolute path for source
	absSource, err := filepath.Abs(source)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for %s: %w", source, err)
	}

	// Create symlink
	log.Printf("Linking %s -> %s\n", target, absSource)
	if err := os.Symlink(absSource, target); err != nil {
		return fmt.Errorf("failed to create symlink %s -> %s: %w", target, absSource, err)
	}

	return nil
}
