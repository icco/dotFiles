package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

func buildStructure(homeDir string) error {
	for _, dir := range []string{"Projects", "bin", "tmp"} {
		dirPath := filepath.Join(homeDir, dir)
		if err := os.MkdirAll(dirPath, 0750); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
		}
	}
	return nil
}

func linkFiles() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	if err := buildStructure(homeDir); err != nil {
		return fmt.Errorf("failed to build directory structure: %w", err)
	}

	if err := linkDirectory("link", homeDir, true); err != nil {
		return fmt.Errorf("failed to link main files: %w", err)
	}

	if err := linkSpecificFiles(homeDir); err != nil {
		return fmt.Errorf("failed to link specific files: %w", err)
	}

	if err := linkBinFiles(homeDir); err != nil {
		return fmt.Errorf("failed to link bin files: %w", err)
	}

	return nil
}

func linkDirectory(sourceDir, homeDir string, addDot bool) error {
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", sourceDir, err)
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(sourceDir, entry.Name())
		name := entry.Name()
		if addDot {
			name = "." + name
		}
		targetPath := filepath.Join(homeDir, name)

		if err := createSymlink(sourcePath, targetPath, homeDir); err != nil {
			return fmt.Errorf("failed to link %s: %w", sourcePath, err)
		}
	}

	return nil
}

func linkSpecificFiles(homeDir string) error {
	return filepath.WalkDir("specific", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel("specific", path)
		if err != nil {
			return fmt.Errorf("failed to get relative path for %s: %w", path, err)
		}

		targetPath := filepath.Join(homeDir, "."+relativePath)
		if err := os.MkdirAll(filepath.Dir(targetPath), 0750); err != nil {
			return fmt.Errorf("failed to create directory for %s: %w", targetPath, err)
		}

		return createSymlink(path, targetPath, homeDir)
	})
}

func linkBinFiles(homeDir string) error {
	entries, err := os.ReadDir("bin")
	if err != nil {
		return fmt.Errorf("failed to read bin directory: %w", err)
	}

	for _, entry := range entries {
		sourcePath := filepath.Join("bin", entry.Name())
		targetPath := filepath.Join(homeDir, sourcePath)

		if err := createSymlink(sourcePath, targetPath, homeDir); err != nil {
			return fmt.Errorf("failed to link bin file %s: %w", sourcePath, err)
		}
	}

	return nil
}

// createSymlink replaces target with a symlink, backing the existing file up to ~/tmp.
func createSymlink(source, target, homeDir string) error {
	if info, err := os.Lstat(target); err == nil {
		// A pre-existing symlink has no data to preserve — just drop it.
		if info.Mode()&os.ModeSymlink == 0 {
			backupDir := filepath.Join(homeDir, "tmp")
			if err := os.MkdirAll(backupDir, 0750); err != nil {
				return fmt.Errorf("failed to create backup directory %s: %w", backupDir, err)
			}
			backupPath := filepath.Join(backupDir, fmt.Sprintf("%s.%d.backup", filepath.Base(target), time.Now().Unix()))

			log.Printf("Backing up %s to %s\n", target, backupPath)
			if err := backupFile(target, backupPath); err != nil {
				return fmt.Errorf("failed to backup %s: %w", target, err)
			}
		}

		if err := os.RemoveAll(target); err != nil {
			return fmt.Errorf("failed to remove existing %s: %w", target, err)
		}
	}

	absSource, err := filepath.Abs(source)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for %s: %w", source, err)
	}

	log.Printf("Linking %s -> %s\n", target, absSource)
	if err := os.Symlink(absSource, target); err != nil {
		return fmt.Errorf("failed to create symlink %s -> %s: %w", target, absSource, err)
	}

	return nil
}

func backupFile(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	if srcInfo.IsDir() {
		return copyDir(src, dst)
	}
	return copyFile(src, dst)
}

func copyFile(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// #nosec G304 -- src is a program-controlled path inside link/ or a temp clone dir.
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// #nosec G306,G703 -- preserves srcInfo.Mode() to faithfully mirror the source file; dst is program-controlled.
	return os.WriteFile(dst, data, srcInfo.Mode())
}

func copyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		if relPath == "." {
			return nil
		}

		dstPath := filepath.Join(dst, relPath)
		if d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			return os.MkdirAll(dstPath, info.Mode())
		}
		return copyFile(path, dstPath)
	})
}
