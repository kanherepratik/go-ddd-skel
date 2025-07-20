package main

import (
	"os"

	"github.com/kanherepratik/go-ddd-skel/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-ddd-skel",
	Short: "Go DDD project scaffolding tool",
	Long: `A command-line tool for generating Go projects following 
Domain-Driven Design principles with opinionated structure.`,
}

func main() {
	cmd.Init(rootCmd)
	cmd.InitDomain(rootCmd)
	cmd.InitUsecase(rootCmd)
	cmd.InitHandler(rootCmd)
	cmd.InitTests(rootCmd)
	cmd.InitArch(rootCmd)
	cmd.InitDocs(rootCmd)
	cmd.InitPlugin(rootCmd)
	cmd.InitDX(rootCmd)
	cmd.InitMonorepo(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
