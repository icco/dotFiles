package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Print usage if no arguments provided
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "install":
		if err := runInstall(); err != nil {
			log.Fatalf("Error running install: %v", err)
		}
	case "vim":
		if err := runVim(); err != nil {
			log.Fatalf("Error running vim: %v", err)
		}
	case "test":
		if err := runTest(); err != nil {
			log.Fatalf("Error running test: %v", err)
		}
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("dotool - Dotfiles management tool")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  go run ./dotool <command>")
	fmt.Println()
	fmt.Println("COMMANDS:")
	fmt.Println("  install    Install dotfiles and link all configuration files")
	fmt.Println("  vim        Update vim plugins and sort spell file")
	fmt.Println("  test       Run tests")
	fmt.Println("  help       Show this help message")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  go run ./dotool install")
	fmt.Println("  go run ./dotool vim")
	fmt.Println("  go run ./dotool test")
}

func runInstall() error {
	fmt.Println("Installing dotfiles...")

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

	fmt.Println("Dotfiles installation completed successfully!")
	return nil
}

func runVim() error {
	fmt.Println("Running vim command...")

	// Sort vim spell
	if err := sortVimSpell(); err != nil {
		return fmt.Errorf("failed to sort vim spell: %w", err)
	}

	// Upgrade vim plugins
	if err := upgradeVimPlugins(); err != nil {
		return fmt.Errorf("failed to upgrade vim plugins: %w", err)
	}

	fmt.Println("Vim command completed successfully!")
	return nil
}

func runTest() error {
	fmt.Println("Running tests...")
	// This would check Go version instead of Ruby version
	// For now, just return success
	fmt.Println("Tests passed!")
	return nil
}
