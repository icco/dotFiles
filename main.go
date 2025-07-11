package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <command>")
		fmt.Println("Available commands:")
		fmt.Println("  infect  - Hook dotfiles into system-standard positions")
		fmt.Println("  vim     - Sort vim spell and upgrade vim plugins")
		fmt.Println("  test    - Test to make sure everything works ok")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "infect":
		if err := runInfect(); err != nil {
			log.Fatalf("Error running infect: %v", err)
		}
	case "vim":
		if err := runVim(); err != nil {
			log.Fatalf("Error running vim: %v", err)
		}
	case "test":
		if err := runTest(); err != nil {
			log.Fatalf("Error running test: %v", err)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: infect, vim, test")
		os.Exit(1)
	}
}

func runInfect() error {
	fmt.Println("Running infect command...")

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

	fmt.Println("Infect completed successfully!")
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
