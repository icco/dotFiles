package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLinkDirectory(t *testing.T) {
	tempDir := t.TempDir()

	sourceDir := filepath.Join(tempDir, "link")
	if err := os.MkdirAll(sourceDir, 0750); err != nil {
		t.Fatalf("failed to create source directory: %v", err)
	}

	testFiles := []string{"file1.txt", "file2.txt", "config.conf"}
	for _, file := range testFiles {
		if err := os.WriteFile(filepath.Join(sourceDir, file), []byte("test content for "+file), 0600); err != nil {
			t.Fatalf("failed to create test file %s: %v", file, err)
		}
	}

	t.Run("with dot prefix", func(t *testing.T) {
		homeDir := t.TempDir()
		if err := linkDirectory(sourceDir, homeDir, true); err != nil {
			t.Errorf("linkDirectory() error = %v", err)
		}
		for _, file := range testFiles {
			if _, err := os.Lstat(filepath.Join(homeDir, "."+file)); err != nil {
				t.Errorf("expected dot-prefixed symlink for %s: %v", file, err)
			}
		}
	})

	t.Run("without dot prefix", func(t *testing.T) {
		homeDir := t.TempDir()
		if err := linkDirectory(sourceDir, homeDir, false); err != nil {
			t.Errorf("linkDirectory() error = %v", err)
		}
		for _, file := range testFiles {
			if _, err := os.Lstat(filepath.Join(homeDir, file)); err != nil {
				t.Errorf("expected symlink for %s: %v", file, err)
			}
		}
	})
}

func TestLinkSpecificFiles(t *testing.T) {
	tempDir := t.TempDir()
	t.Chdir(tempDir)

	testFiles := map[string]string{
		"config.conf":                    "config content",
		"config/app/settings.json":       "settings content",
		"config/app/another-config.yaml": "yaml content",
	}
	for filePath, content := range testFiles {
		fullPath := filepath.Join(tempDir, "specific", filePath)
		if err := os.MkdirAll(filepath.Dir(fullPath), 0750); err != nil {
			t.Fatalf("failed to create directory for %s: %v", filePath, err)
		}
		if err := os.WriteFile(fullPath, []byte(content), 0600); err != nil {
			t.Fatalf("failed to create test file %s: %v", filePath, err)
		}
	}

	homeDir := filepath.Join(tempDir, "home")
	if err := os.MkdirAll(homeDir, 0750); err != nil {
		t.Fatalf("failed to create home directory: %v", err)
	}

	if err := linkSpecificFiles(homeDir); err != nil {
		t.Errorf("linkSpecificFiles() error = %v", err)
	}

	for filePath := range testFiles {
		if _, err := os.Lstat(filepath.Join(homeDir, "."+filePath)); err != nil {
			t.Errorf("expected symlink %s: %v", filePath, err)
		}
	}
}

func TestLinkBinFiles(t *testing.T) {
	tempDir := t.TempDir()
	t.Chdir(tempDir)

	if err := os.MkdirAll(filepath.Join(tempDir, "bin"), 0750); err != nil {
		t.Fatalf("failed to create bin directory: %v", err)
	}
	testFiles := []string{"script1.sh", "script2.py", "tool"}
	for _, file := range testFiles {
		if err := os.WriteFile(filepath.Join(tempDir, "bin", file), []byte("#!/bin/bash\necho test"), 0600); err != nil {
			t.Fatalf("failed to create test file %s: %v", file, err)
		}
	}

	homeDir := filepath.Join(tempDir, "home")
	if err := os.MkdirAll(filepath.Join(homeDir, "bin"), 0750); err != nil {
		t.Fatalf("failed to create home bin directory: %v", err)
	}

	if err := linkBinFiles(homeDir); err != nil {
		t.Errorf("linkBinFiles() error = %v", err)
	}

	for _, file := range testFiles {
		if _, err := os.Lstat(filepath.Join(homeDir, "bin", file)); err != nil {
			t.Errorf("expected symlink %s: %v", file, err)
		}
	}
}
