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
	targetFile := filepath.Join(tempDir, "target.txt")
	if err := os.WriteFile(sourceFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("failed to create source file: %v", err)
	}

	if err := createSymlink(sourceFile, targetFile, tempDir); err != nil {
		t.Errorf("createSymlink() error = %v", err)
	}
	if _, err := os.Lstat(targetFile); err != nil {
		t.Errorf("symlink was not created at %s: %v", targetFile, err)
	}

	// Re-linking over an existing target should succeed (and back the old one up).
	if err := createSymlink(sourceFile, targetFile, tempDir); err != nil {
		t.Errorf("createSymlink() over existing file error = %v", err)
	}
}
