package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	docType string
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation",
	Long: `Generates project documentation in various formats:
- Markdown
- OpenAPI/Swagger`,
	Run: func(cmd *cobra.Command, args []string) {
		generateDocumentation()
	},
}

func generateDocumentation() {
	switch strings.ToLower(docType) {
	case "markdown":
		generateMarkdownDocs()
	case "openapi":
		generateOpenAPIDocs()
	default:
		fmt.Println("Unsupported documentation type. Use --type [markdown|openapi]")
		os.Exit(1)
	}
}

func generateMarkdownDocs() {
	// Generate markdown documentation
	cmd := exec.Command("godoc2md", "./...")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error generating markdown docs: %v\n", err)
		os.Exit(1)
	}

	// Write to README.md
	if err := os.WriteFile("README.md", output, 0644); err != nil {
		fmt.Printf("Error writing markdown docs: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Markdown documentation generated in README.md")
}

func generateOpenAPIDocs() {
	// Generate OpenAPI documentation
	cmd := exec.Command("swag", "init")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error generating OpenAPI docs: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("OpenAPI documentation generated in docs/")
}

func InitDocs(rootCmd *cobra.Command) {
	rootCmd.AddCommand(docsCmd)
	docsCmd.Flags().StringVarP(&docType, "type", "t", "markdown", "Documentation type (markdown|openapi)")
}
