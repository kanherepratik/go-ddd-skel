package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var monorepoCmd = &cobra.Command{
	Use:   "monorepo",
	Short: "Setup monorepo structure",
	Long: `Creates a monorepo-compatible folder structure with:
- Shared packages
- Service-specific directories
- Common configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		setupMonorepo()
	},
}

func setupMonorepo() {
	// Create base directories
	dirs := []string{
		"services/service1",
		"services/service2",
		"packages/shared",
		"packages/config",
		"scripts",
		"deploy",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			os.Exit(1)
		}
	}

	// Create shared package
	sharedPkg := `package shared

// Add shared utilities here
`
	if err := os.WriteFile(filepath.Join("packages/shared/shared.go"), []byte(sharedPkg), 0644); err != nil {
		fmt.Printf("Error creating shared package: %v\n", err)
		os.Exit(1)
	}

	// Create config package
	configPkg := `package config

// Add common configuration here
`
	if err := os.WriteFile(filepath.Join("packages/config/config.go"), []byte(configPkg), 0644); err != nil {
		fmt.Printf("Error creating config package: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Monorepo structure created successfully")
}

func InitSetupMonorepo(rootCmd *cobra.Command) {
	rootCmd.AddCommand(monorepoCmd)
}
