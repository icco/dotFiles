package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLinkDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "dotfiles_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test source directory
	sourceDir := filepath.Join(tempDir, "link")
	if err := os.MkdirAll(sourceDir, 0755); err != nil {
		t.Fatalf("Failed to create source directory: %v", err)
	}

	// Create test files in source directory
	testFiles := []string{"file1.txt", "file2.txt", "config.conf"}
	for _, file := range testFiles {
		filePath := filepath.Join(sourceDir, file)
		content := []byte("test content for " + file)
		if err := os.WriteFile(filePath, content, 0644); err != nil {
			t.Fatalf("Failed to create test file %s: %v", file, err)
		}
	}

	// Create test home directory
	homeDir := filepath.Join(tempDir, "home")
	if err := os.MkdirAll(homeDir, 0755); err != nil {
		t.Fatalf("Failed to create home directory: %v", err)
	}

	// Test linking directory with dot prefix
	err = linkDirectory(sourceDir, homeDir, true)
	if err != nil {
		t.Errorf("linkDirectory() error = %v, wantErr false", err)
	}

	// Check that symlinks were created with dot prefix
	for _, file := range testFiles {
		targetPath := filepath.Join(homeDir, "."+file)
		if _, err := os.Lstat(targetPath); os.IsNotExist(err) {
			t.Errorf("Expected symlink %s was not created", targetPath)
		}
	}

	// Test linking directory without dot prefix
	homeDir2 := filepath.Join(tempDir, "home2")
	if err := os.MkdirAll(homeDir2, 0755); err != nil {
		t.Fatalf("Failed to create second home directory: %v", err)
	}

	err = linkDirectory(sourceDir, homeDir2, false)
	if err != nil {
		t.Errorf("linkDirectory() without dot error = %v, wantErr false", err)
	}

	// Check that symlinks were created without dot prefix
	for _, file := range testFiles {
		targetPath := filepath.Join(homeDir2, file)
		if _, err := os.Lstat(targetPath); os.IsNotExist(err) {
			t.Errorf("Expected symlink %s was not created", targetPath)
		}
	}
}

func TestLinkSpecificFiles(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "dotfiles_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test specific directory structure
	specificDir := filepath.Join(tempDir, "specific")
	if err := os.MkdirAll(specificDir, 0755); err != nil {
		t.Fatalf("Failed to create specific directory: %v", err)
	}

	// Create nested directory structure
	nestedDir := filepath.Join(specificDir, "config", "app")
	if err := os.MkdirAll(nestedDir, 0755); err != nil {
		t.Fatalf("Failed to create nested directory: %v", err)
	}

	// Create test files
	testFiles := map[string]string{
		"config.conf":                    "config content",
		"config/app/settings.json":       "settings content",
		"config/app/another-config.yaml": "yaml content",
	}

	for filePath, content := range testFiles {
		fullPath := filepath.Join(specificDir, filePath)
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create directory for %s: %v", filePath, err)
		}
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			t.Fatalf("Failed to create test file %s: %v", filePath, err)
		}
	}

	// Create test home directory
	homeDir := filepath.Join(tempDir, "home")
	if err := os.MkdirAll(homeDir, 0755); err != nil {
		t.Fatalf("Failed to create home directory: %v", err)
	}

	// Change to temp directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Test linking specific files
	err = linkSpecificFiles(homeDir)
	if err != nil {
		t.Errorf("linkSpecificFiles() error = %v, wantErr false", err)
	}

	// Check that symlinks were created with dot prefix
	for filePath := range testFiles {
		targetPath := filepath.Join(homeDir, "."+filePath)
		if _, err := os.Lstat(targetPath); os.IsNotExist(err) {
			t.Errorf("Expected symlink %s was not created", targetPath)
		}
	}
}

func TestLinkBinFiles(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "dotfiles_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test bin directory
	binDir := filepath.Join(tempDir, "bin")
	if err := os.MkdirAll(binDir, 0755); err != nil {
		t.Fatalf("Failed to create bin directory: %v", err)
	}

	// Create test files in bin directory
	testFiles := []string{"script1.sh", "script2.py", "tool"}
	for _, file := range testFiles {
		filePath := filepath.Join(binDir, file)
		content := []byte("#!/bin/bash\necho 'test script'")
		if err := os.WriteFile(filePath, content, 0755); err != nil {
			t.Fatalf("Failed to create test file %s: %v", file, err)
		}
	}

	// Create test home directory
	homeDir := filepath.Join(tempDir, "home")
	if err := os.MkdirAll(homeDir, 0755); err != nil {
		t.Fatalf("Failed to create home directory: %v", err)
	}

	// Change to temp directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Create the bin directory in home first
	homeBinDir := filepath.Join(homeDir, "bin")
	if err := os.MkdirAll(homeBinDir, 0755); err != nil {
		t.Fatalf("Failed to create home bin directory: %v", err)
	}

	// Test linking bin files
	err = linkBinFiles(homeDir)
	if err != nil {
		t.Errorf("linkBinFiles() error = %v, wantErr false", err)
	}

	// Check that symlinks were created
	for _, file := range testFiles {
		targetPath := filepath.Join(homeDir, "bin", file)
		if _, err := os.Lstat(targetPath); os.IsNotExist(err) {
			t.Errorf("Expected symlink %s was not created", targetPath)
		}
	}
}
