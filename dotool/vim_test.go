package main

import (
	"context"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestSortSpellFile(t *testing.T) {
	tempDir := t.TempDir()
	spellFile := filepath.Join(tempDir, "en.utf-8.add")

	input := "zebra\nalpha\nBeta\nzebra\nalpha\n\nbeta\n"
	if err := os.WriteFile(spellFile, []byte(input), 0600); err != nil {
		t.Fatalf("failed to create spell file: %v", err)
	}

	if err := sortSpellFile(spellFile); err != nil {
		t.Fatalf("sortSpellFile() error = %v", err)
	}

	// #nosec G304 -- spellFile is a t.TempDir() path in test scope.
	data, err := os.ReadFile(spellFile)
	if err != nil {
		t.Fatalf("failed to read spell file: %v", err)
	}

	got := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	want := []string{"alpha", "Beta", "zebra"}
	if !slices.Equal(got, want) {
		t.Errorf("sortSpellFile() = %v, want %v", got, want)
	}
}

func TestSortVimSpellMissingFile(t *testing.T) {
	t.Chdir(t.TempDir())
	if err := sortVimSpell(context.Background()); err == nil {
		t.Error("sortVimSpell() should return error when spell file doesn't exist")
	}
}

func TestUpgradePluginInvalidFormat(t *testing.T) {
	cases := []string{"invalid-repo-format", "", "/missing-owner", "missing-name/"}
	for _, repo := range cases {
		t.Run(repo, func(t *testing.T) {
			if err := upgradePlugin(context.Background(), repo); err == nil {
				t.Errorf("upgradePlugin(%q) expected error", repo)
			}
		})
	}
}
