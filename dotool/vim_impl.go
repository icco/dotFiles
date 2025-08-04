package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// sortVimSpell sorts the vim spell file and commits the changes
func sortVimSpell() error {
	log.Println("Sorting vim spell...")

	spellFile := "link/vim/spell/en.utf-8.add"

	// Check if spell file exists
	if _, err := os.Stat(spellFile); os.IsNotExist(err) {
		return fmt.Errorf("spell file %s does not exist", spellFile)
	}

	// Create temporary file
	tempFile := "t"

	// Sort the spell file
	cmd := exec.Command("sh", "-c", fmt.Sprintf("cat %s | sort -if | uniq > %s", spellFile, tempFile))
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to sort spell file: %s, %w", string(output), err)
	}

	// Move temporary file back
	cmd = exec.Command("mv", tempFile, spellFile)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to move sorted file: %s, %w", string(output), err)
	}

	// Commit the changes
	cmd = exec.Command("git", "commit", "-a", "-m", "vim spell sort")
	if output, err := cmd.CombinedOutput(); err != nil {
		// Check if the error is due to no changes to commit
		if strings.Contains(string(output), "nothing to commit") {
			log.Println("No changes to commit - spell file was already sorted")
		} else {
			return fmt.Errorf("failed to commit spell changes: %s, %w", string(output), err)
		}
	}

	log.Println("Vim spell sorted and committed successfully!")
	return nil
}

// upgradeVimPlugins upgrades all vim plugins by cloning them fresh
func upgradeVimPlugins() error {
	log.Println("Upgrading vim plugins...")

	repos := []string{
		"airblade/vim-rooter",
		"craigmac/vim-mermaid",
		"dense-analysis/ale",
		"editorconfig/editorconfig-vim",
		"fatih/vim-go",
		"github/copilot.vim",
		"godlygeek/tabular",
		"google/vim-jsonnet",
		"grafana/vim-alloy",
		"hashivim/vim-terraform",
		"isobit/vim-caddyfile",
		"jparise/vim-graphql",
		"junegunn/fzf.vim",
		"kaarmu/typst.vim",
		"mhinz/vim-signify",
		"nanotee/zoxide.vim",
		"nathanielc/vim-tickscript",
		"preservim/tagbar",
		"preservim/vim-markdown",
		"tpope/vim-commentary",
		"tpope/vim-fugitive",
		"uarun/vim-protobuf",
		"wakatime/vim-wakatime",
	}

	for _, repo := range repos {
		if err := upgradePlugin(repo); err != nil {
			return fmt.Errorf("failed to upgrade plugin %s: %w", repo, err)
		}
	}

	// Commit all plugin changes
	cmd := exec.Command("git", "commit", "-a", "-m", "vim upgrades")
	if output, err := cmd.CombinedOutput(); err != nil {
		// Check if the error is due to no changes to commit
		if strings.Contains(string(output), "nothing to commit") ||
			strings.Contains(string(output), "nothing added to commit") {
			log.Println("No changes to commit - plugins were already up to date")
		} else {
			return fmt.Errorf("failed to commit plugin changes: %s, %w", string(output), err)
		}
	}

	log.Println("All vim plugins upgraded successfully!")
	return nil
}

// upgradePlugin upgrades a single vim plugin
func upgradePlugin(repo string) error {
	log.Printf("Upgrading plugin: %s\n", repo)

	// Extract plugin name from repo
	parts := strings.Split(repo, "/")
	if len(parts) != 2 {
		return fmt.Errorf("invalid repo format: %s", repo)
	}
	pluginName := parts[1]

	// Plugin directory path
	pluginDir := filepath.Join("link/vim/bundle", pluginName)

	// Remove existing plugin directory
	if err := os.RemoveAll(pluginDir); err != nil {
		return fmt.Errorf("failed to remove existing plugin directory %s: %w", pluginDir, err)
	}

	// Clone the plugin
	cmd := exec.Command("git", "clone", fmt.Sprintf("git@github.com:%s.git", repo), pluginDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to clone plugin %s: %s, %w", repo, string(output), err)
	}

	// Remove .git directory from plugin
	gitDir := filepath.Join(pluginDir, ".git")
	if err := os.RemoveAll(gitDir); err != nil {
		return fmt.Errorf("failed to remove .git directory from %s: %w", pluginDir, err)
	}

	// Add plugin to git
	cmd = exec.Command("git", "add", pluginDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to add plugin %s to git: %s, %w", pluginDir, string(output), err)
	}

	return nil
}
