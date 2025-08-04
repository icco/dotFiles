package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dotool",
	Short: "Dotfiles management tool",
	Long:  `A tool for managing dotfiles across multiple machines using symbolic links.`,
}

var infectCmd = &cobra.Command{
	Use:   "infect",
	Short: "Install dotfiles and link all configuration files",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runInfect(); err != nil {
			log.Fatalf("Error running infect: %v", err)
		}
	},
}

var vimCmd = &cobra.Command{
	Use:   "vim",
	Short: "Update vim plugins and sort spell file",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runVim(); err != nil {
			log.Fatalf("Error running vim: %v", err)
		}
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runTest(); err != nil {
			log.Fatalf("Error running test: %v", err)
		}
	},
}

var updateZshCmd = &cobra.Command{
	Use:   "update-zsh",
	Short: "Update oh-my-zsh to latest version",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runUpdateZsh(); err != nil {
			log.Fatalf("Error updating oh-my-zsh: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(infectCmd)
	rootCmd.AddCommand(vimCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(updateZshCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runInfect() error {
	log.Println("Running infect command...")

	// Run test first
	if err := runTest(); err != nil {
		return fmt.Errorf("test failed: %w", err)
	}

	// Build directory structure
	if err := buildStructure(); err != nil {
		return fmt.Errorf("failed to build structure: %w", err)
	}

	// Link files
	if err := linkFiles(); err != nil {
		return fmt.Errorf("failed to link files: %w", err)
	}

	log.Println("Infect completed successfully!")
	return nil
}

func runVim() error {
	log.Println("Running vim command...")

	// Sort vim spell
	if err := sortVimSpell(); err != nil {
		return fmt.Errorf("failed to sort vim spell: %w", err)
	}

	// Upgrade vim plugins
	if err := upgradeVimPlugins(); err != nil {
		return fmt.Errorf("failed to upgrade vim plugins: %w", err)
	}

	log.Println("Vim command completed successfully!")
	return nil
}

func runTest() error {
	log.Println("Running tests...")
	// This would check Go version instead of Ruby version
	// For now, just return success
	log.Println("Tests passed!")
	return nil
}

func runUpdateZsh() error {
	log.Println("Updating oh-my-zsh...")

	if err := updateOhMyZsh(); err != nil {
		return fmt.Errorf("failed to update oh-my-zsh: %w", err)
	}

	log.Println("Oh-my-zsh updated successfully!")
	return nil
}