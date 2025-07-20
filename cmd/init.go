package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new DDD project",
	Long: `Creates a new Go project with Domain-Driven Design structure including:
- cmd/ for command implementations
- internal/ for core domain logic
- pkg/ for shared utilities
- config/ for configuration
- migrations/ for database migrations`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		createProjectStructure(projectName)
	},
}

func createProjectStructure(projectName string) {
	dirs := []string{
		"cmd",
		"internal/adapters/external",
		"internal/adapters/persistence",
		"internal/adapters/ports",
		"internal/config",
		"internal/core",
		"internal/interfaces",
		"internal/usecase",
		"migrations",
		"pkg",
		"scripts",
		"sql",
		"static",
	}

	for _, dir := range dirs {
		fullPath := filepath.Join(projectName, dir)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", fullPath, err)
			os.Exit(1)
		}
	}

	fmt.Printf("Successfully created DDD project structure in %s/\n", projectName)
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(initCmd)
}
