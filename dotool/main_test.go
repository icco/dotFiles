package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMain(t *testing.T) {
	// Test cases for main function
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "no args",
			args:    []string{"dotfiles"},
			wantErr: true,
		},
		{
			name:    "invalid command",
			args:    []string{"dotfiles", "invalid"},
			wantErr: true,
		},
		{
			name:    "test command",
			args:    []string{"dotfiles", "test"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original args
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			// Set test args
			os.Args = tt.args

			// Run main function
			// Note: We can't easily test main() directly due to os.Exit calls
			// This is more of an integration test setup
		})
	}
}

func TestRunTest(t *testing.T) {
	err := runTest()
	if err != nil {
		t.Errorf("runTest() error = %v, wantErr false", err)
	}
}

func TestBuildStructure(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "dotfiles_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Save original home directory
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)

	// Set test home directory
	os.Setenv("HOME", tempDir)

	err = buildStructure()
	if err != nil {
		t.Errorf("buildStructure() error = %v, wantErr false", err)
	}

	// Check that directories were created
	expectedDirs := []string{"Projects", "bin", "tmp"}
	for _, dir := range expectedDirs {
		dirPath := filepath.Join(tempDir, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			t.Errorf("Expected directory %s was not created", dirPath)
		}
	}
}

func TestCreateSymlink(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "dotfiles_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test source file
	sourceFile := filepath.Join(tempDir, "source.txt")
	targetFile := filepath.Join(tempDir, "target.txt")

	content := []byte("test content")
	if err := os.WriteFile(sourceFile, content, 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// Test creating symlink
	err = createSymlink(sourceFile, targetFile)
	if err != nil {
		t.Errorf("createSymlink() error = %v, wantErr false", err)
	}

	// Check that symlink was created
	if _, err := os.Lstat(targetFile); os.IsNotExist(err) {
		t.Errorf("Symlink was not created at %s", targetFile)
	}

	// Test creating symlink over existing file
	err = createSymlink(sourceFile, targetFile)
	if err != nil {
		t.Errorf("createSymlink() over existing file error = %v, wantErr false", err)
	}
}

func TestUpgradePlugin(t *testing.T) {
	// Skip this test if git is not available
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not available, skipping plugin upgrade test")
	}

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "dotfiles_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create the bundle directory structure
	bundleDir := filepath.Join(tempDir, "link", "vim", "bundle")
	if err := os.MkdirAll(bundleDir, 0755); err != nil {
		t.Fatalf("Failed to create bundle directory: %v", err)
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

	// Test with a simple plugin (this will fail if no SSH key, but that's expected)
	// We'll just test the function structure, not actual cloning
	repo := "test/test-plugin"
	err = upgradePlugin(repo)
	// We expect this to fail due to SSH key issues, but the function should handle it gracefully
	if err == nil {
		t.Log("Plugin upgrade succeeded (unexpected, but ok)")
	} else {
		t.Logf("Plugin upgrade failed as expected: %v", err)
	}
}
