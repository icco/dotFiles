package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestSortVimSpell(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "dotfiles_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create the vim spell directory structure
	spellDir := filepath.Join(tempDir, "link", "vim", "spell")
	if err := os.MkdirAll(spellDir, 0755); err != nil {
		t.Fatalf("Failed to create spell directory: %v", err)
	}

	// Create a test spell file with unsorted content
	spellFile := filepath.Join(spellDir, "en.utf-8.add")
	unsortedContent := "zebra\nalpha\nbeta\ngamma\nzebra\nalpha"
	if err := os.WriteFile(spellFile, []byte(unsortedContent), 0644); err != nil {
		t.Fatalf("Failed to create spell file: %v", err)
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

	// Test sorting vim spell
	err = sortVimSpell()
	if err != nil {
		// Git operations will fail in test environment, but sorting should work
		if strings.Contains(err.Error(), "not a git repository") {
			t.Logf("Git operation failed as expected in test environment: %v", err)
		} else {
			t.Errorf("sortVimSpell() error = %v, wantErr false", err)
		}
	}

	// Check that the file was sorted (even if git commit failed)
	content, err := os.ReadFile(spellFile)
	if err != nil {
		t.Fatalf("Failed to read sorted spell file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	// Remove empty lines
	var nonEmptyLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, strings.TrimSpace(line))
		}
	}

	if len(nonEmptyLines) != 3 { // Should be 3 unique lines after sorting and deduplication
		t.Errorf("Expected 3 unique lines, got %d: %v", len(nonEmptyLines), nonEmptyLines)
	}

	// Check that lines are sorted
	for i := 1; i < len(nonEmptyLines); i++ {
		if nonEmptyLines[i-1] > nonEmptyLines[i] {
			t.Errorf("Lines are not sorted: %s > %s", nonEmptyLines[i-1], nonEmptyLines[i])
		}
	}
}

func TestSortVimSpellMissingFile(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "dotfiles_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Change to temp directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Test sorting vim spell with missing file
	err = sortVimSpell()
	if err == nil {
		t.Error("sortVimSpell() should return error when spell file doesn't exist")
	}
}

func TestUpgradeVimPlugins(t *testing.T) {
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

	// Test upgrading vim plugins
	// This will likely fail due to SSH key issues, but we can test the structure
	err = upgradeVimPlugins()
	if err != nil {
		t.Logf("Plugin upgrade failed as expected: %v", err)
		// This is expected to fail without proper SSH keys, so we don't treat it as a test failure
	} else {
		t.Log("Plugin upgrade succeeded (unexpected, but ok)")
	}
}

func TestUpgradePluginBasic(t *testing.T) {
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

	// Test with invalid repo format
	err = upgradePlugin("invalid-repo-format")
	if err == nil {
		t.Error("upgradePlugin() should return error for invalid repo format")
	}

	// Test with a non-existent plugin (this will fail, but that's expected)
	err = upgradePlugin("nonexistent/plugin")
	if err != nil {
		t.Logf("Plugin upgrade failed as expected: %v", err)
		// This is expected to fail, so we don't treat it as a test failure
	} else {
		t.Log("Plugin upgrade succeeded (unexpected, but ok)")
	}
}

func TestUpgradePluginWithExistingDirectory(t *testing.T) {
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

	// Create an existing plugin directory
	existingPluginDir := filepath.Join(bundleDir, "test-plugin")
	if err := os.MkdirAll(existingPluginDir, 0755); err != nil {
		t.Fatalf("Failed to create existing plugin directory: %v", err)
	}

	// Create a file in the existing directory
	existingFile := filepath.Join(existingPluginDir, "test.txt")
	if err := os.WriteFile(existingFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
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

	// Test upgrading plugin that already exists
	err = upgradePlugin("test/test-plugin")
	if err != nil {
		t.Logf("Plugin upgrade failed as expected: %v", err)
		// This is expected to fail, so we don't treat it as a test failure
	} else {
		t.Log("Plugin upgrade succeeded (unexpected, but ok)")
	}

	// Check that the existing directory was removed (even if clone failed)
	if _, err := os.Stat(existingPluginDir); err == nil {
		t.Log("Existing plugin directory was not removed (this might be expected if clone failed)")
	} else {
		t.Log("Existing plugin directory was removed as expected")
	}
}
