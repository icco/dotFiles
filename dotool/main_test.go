package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildStructure(t *testing.T) {
	tempDir := t.TempDir()

	if err := buildStructure(tempDir); err != nil {
		t.Fatalf("buildStructure() error = %v", err)
	}

	for _, dir := range []string{"Projects", "bin", "tmp"} {
		if _, err := os.Stat(filepath.Join(tempDir, dir)); err != nil {
			t.Errorf("expected directory %s: %v", dir, err)
		}
	}
}

func TestCreateSymlink(t *testing.T) {
	tempDir := t.TempDir()

	sourceFile := filepath.Join(tempDir, "source.txt")
	if err := os.WriteFile(sourceFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("failed to create source file: %v", err)
	}

	countBackups := func() int {
		entries, _ := os.ReadDir(filepath.Join(tempDir, "tmp"))
		return len(entries)
	}

	t.Run("fresh target", func(t *testing.T) {
		target := filepath.Join(tempDir, "fresh.txt")
		if err := createSymlink(sourceFile, target, tempDir); err != nil {
			t.Fatalf("createSymlink() error = %v", err)
		}
		if _, err := os.Lstat(target); err != nil {
			t.Errorf("symlink not created: %v", err)
		}
	})

	t.Run("over existing symlink — no backup", func(t *testing.T) {
		target := filepath.Join(tempDir, "linked.txt")
		if err := createSymlink(sourceFile, target, tempDir); err != nil {
			t.Fatalf("setup symlink: %v", err)
		}
		before := countBackups()
		if err := createSymlink(sourceFile, target, tempDir); err != nil {
			t.Fatalf("re-link error: %v", err)
		}
		if countBackups() != before {
			t.Errorf("expected no new backup when replacing a symlink")
		}
	})

	t.Run("over existing regular file — backed up", func(t *testing.T) {
		target := filepath.Join(tempDir, "regular.txt")
		if err := os.WriteFile(target, []byte("pre-existing"), 0644); err != nil {
			t.Fatalf("setup file: %v", err)
		}
		before := countBackups()
		if err := createSymlink(sourceFile, target, tempDir); err != nil {
			t.Fatalf("link error: %v", err)
		}
		if countBackups() != before+1 {
			t.Errorf("expected one new backup when replacing a regular file")
		}
	})
}
