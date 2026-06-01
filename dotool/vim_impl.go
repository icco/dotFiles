package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

var vimPlugins = []string{
	"airblade/vim-rooter",
	"craigmac/vim-mermaid",
	"dense-analysis/ale",
	"editorconfig/editorconfig-vim",
	"fatih/vim-go",
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

func sortVimSpell(ctx context.Context) error {
	log.Println("Sorting vim spell...")

	spellFile := "link/vim/spell/en.utf-8.add"
	if _, err := os.Stat(spellFile); err != nil {
		return fmt.Errorf("spell file %s: %w", spellFile, err)
	}

	if err := sortSpellFile(spellFile); err != nil {
		return fmt.Errorf("failed to sort spell file: %w", err)
	}

	if err := gitCommitAll(ctx, "vim spell sort"); err != nil {
		return err
	}

	log.Println("Vim spell sorted and committed successfully!")
	return nil
}

func upgradeVimPlugins(ctx context.Context) error {
	log.Println("Upgrading vim plugins...")

	for _, repo := range vimPlugins {
		if err := upgradePlugin(ctx, repo); err != nil {
			return fmt.Errorf("failed to upgrade plugin %s: %w", repo, err)
		}
	}

	if err := gitCommitAll(ctx, "vim upgrades"); err != nil {
		return err
	}

	log.Println("All vim plugins upgraded successfully!")
	return nil
}

func upgradePlugin(ctx context.Context, repo string) error {
	log.Printf("Upgrading plugin: %s\n", repo)

	owner, name, ok := strings.Cut(repo, "/")
	if !ok || owner == "" || name == "" {
		return fmt.Errorf("invalid repo format: %s", repo)
	}

	pluginDir := filepath.Join("link/vim/bundle", name)

	if err := os.RemoveAll(pluginDir); err != nil {
		return fmt.Errorf("failed to remove existing plugin directory %s: %w", pluginDir, err)
	}

	cloneURL := fmt.Sprintf("git@github.com:%s.git", repo)
	// #nosec G204 -- repo comes from the program's hardcoded vimPlugins list, not user input.
	if out, err := exec.CommandContext(ctx, "git", "clone", cloneURL, pluginDir).CombinedOutput(); err != nil {
		return fmt.Errorf("failed to clone plugin %s: %s: %w", repo, string(out), err)
	}

	if err := os.RemoveAll(filepath.Join(pluginDir, ".git")); err != nil {
		return fmt.Errorf("failed to remove .git directory from %s: %w", pluginDir, err)
	}

	// Strip .terraform/ — vendored by hashivim/vim-terraform fixtures, shouldn't be committed.
	walkErr := filepath.WalkDir(pluginDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && d.Name() == ".terraform" {
			// #nosec G122 -- pluginDir is program-controlled, not user-supplied; symlink TOCTOU not a concern for a local dev tool.
			if err := os.RemoveAll(path); err != nil {
				return err
			}
			return filepath.SkipDir
		}
		return nil
	})
	if walkErr != nil {
		return fmt.Errorf("failed to scrub .terraform from %s: %w", pluginDir, walkErr)
	}

	if err := runGit(ctx, "add", pluginDir); err != nil {
		return fmt.Errorf("failed to add plugin %s to git: %w", pluginDir, err)
	}

	return nil
}

func sortSpellFile(filename string) error {
	// #nosec G304 -- filename is a program-controlled path under link/vim/spell.
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	seen := make(map[string]bool)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		key := strings.ToLower(line)
		if seen[key] {
			continue
		}
		seen[key] = true
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	slices.SortFunc(lines, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})

	// #nosec G306 -- vim spell file is a user-facing config; 0644 is the standard mode.
	return os.WriteFile(filename, []byte(strings.Join(lines, "\n")+"\n"), 0644)
}
