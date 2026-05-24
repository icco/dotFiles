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
	RunE:  func(cmd *cobra.Command, args []string) error { return runInfect() },
}

var vimCmd = &cobra.Command{
	Use:   "vim",
	Short: "Update vim plugins and sort spell file",
	RunE:  func(cmd *cobra.Command, args []string) error { return runVim() },
}

var omzCmd = &cobra.Command{
	Use:   "omz",
	Short: "Update oh-my-zsh to latest version",
	RunE:  func(cmd *cobra.Command, args []string) error { return updateOhMyZsh() },
}

func init() {
	rootCmd.AddCommand(infectCmd)
	rootCmd.AddCommand(vimCmd)
	rootCmd.AddCommand(omzCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runInfect() error {
	log.Println("Running infect command...")

	if err := linkFiles(); err != nil {
		return fmt.Errorf("failed to link files: %w", err)
	}

	log.Println("Infect completed successfully!")
	return nil
}

func runVim() error {
	log.Println("Running vim command...")

	if err := sortVimSpell(); err != nil {
		return fmt.Errorf("failed to sort vim spell: %w", err)
	}

	if err := upgradeVimPlugins(); err != nil {
		return fmt.Errorf("failed to upgrade vim plugins: %w", err)
	}

	log.Println("Vim command completed successfully!")
	return nil
}
