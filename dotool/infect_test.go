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

func TestLinkSpecificFilesFollowsSourceSymlink(t *testing.T) {
	tempDir := t.TempDir()
	t.Chdir(tempDir)

	for _, dir := range []string{"specific/codex", "specific/claude", "home"} {
		if err := os.MkdirAll(dir, 0750); err != nil {
			t.Fatalf("failed to create %s: %v", dir, err)
		}
	}
	if err := os.WriteFile("specific/codex/AGENTS.md", []byte("rules"), 0600); err != nil {
		t.Fatalf("failed to write AGENTS.md: %v", err)
	}
	if err := os.Symlink("../codex/AGENTS.md", "specific/claude/CLAUDE.md"); err != nil {
		t.Fatalf("failed to symlink CLAUDE.md: %v", err)
	}

	homeDir := filepath.Join(tempDir, "home")
	if err := linkSpecificFiles(homeDir); err != nil {
		t.Fatalf("linkSpecificFiles() error = %v", err)
	}

	// #nosec G304 -- path is inside t.TempDir().
	got, err := os.ReadFile(filepath.Join(homeDir, ".claude", "CLAUDE.md"))
	if err != nil {
		t.Fatalf("failed to read linked CLAUDE.md: %v", err)
	}
	if string(got) != "rules" {
		t.Errorf("linked CLAUDE.md = %q, want %q", got, "rules")
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

func TestLinkSkipsDotEntries(t *testing.T) {
	tempDir := t.TempDir()
	t.Chdir(tempDir)

	for _, file := range []string{"link/.DS_Store", "link/.zsh_favlist", "link/zshrc", "specific/.DS_Store", "specific/config/app/.DS_Store", "specific/config/app/.envrc"} {
		if err := os.MkdirAll(filepath.Dir(file), 0750); err != nil {
			t.Fatalf("failed to create directory for %s: %v", file, err)
		}
		if err := os.WriteFile(file, []byte("x"), 0600); err != nil {
			t.Fatalf("failed to create test file %s: %v", file, err)
		}
	}

	homeDir := filepath.Join(tempDir, "home")
	if err := os.MkdirAll(homeDir, 0750); err != nil {
		t.Fatalf("failed to create home directory: %v", err)
	}
	if err := linkDirectory("link", homeDir, true); err != nil {
		t.Fatalf("linkDirectory() error = %v", err)
	}
	if err := linkSpecificFiles(homeDir); err != nil {
		t.Fatalf("linkSpecificFiles() error = %v", err)
	}

	for _, want := range []string{".zshrc", ".config/app/.envrc"} {
		if _, err := os.Lstat(filepath.Join(homeDir, want)); err != nil {
			t.Errorf("expected symlink %s: %v", want, err)
		}
	}
	for _, unwanted := range []string{"..DS_Store", "..zsh_favlist", ".config/app/.DS_Store"} {
		if _, err := os.Lstat(filepath.Join(homeDir, unwanted)); err == nil {
			t.Errorf("unexpected symlink %s", unwanted)
		}
	}
}
