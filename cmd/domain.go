package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

var domainCmd = &cobra.Command{
	Use:   "domain [name]",
	Short: "Generate a new domain entity",
	Long: `Creates a new domain entity with:
- Entity struct
- Repository interface
- Value objects (optional)`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domainName := args[0]
		createDomainStructure(domainName)
	},
}

func createDomainStructure(domainName string) {
	// Create domain directory
	domainPath := filepath.Join("internal/core", domainName)
	if err := os.MkdirAll(domainPath, 0755); err != nil {
		fmt.Printf("Error creating domain directory: %v\n", err)
		os.Exit(1)
	}

	// Generate entity file
	entityTemplate := `package {{.Domain}}

type {{.Entity}} struct {
	ID string
	// Add domain-specific fields here
}
`
	generateFile(filepath.Join(domainPath, "entity.go"), entityTemplate, map[string]string{
		"Domain": domainName,
		"Entity": domainName,
	})

	// Generate repository interface
	repoTemplate := `package {{.Domain}}

type {{.Repository}} interface {
	Save(entity *{{.Entity}}) error
	FindByID(id string) (*{{.Entity}}, error)
	// Add additional repository methods here
}
`
	generateFile(filepath.Join(domainPath, "repository.go"), repoTemplate, map[string]string{
		"Domain":     domainName,
		"Repository": domainName + "Repository",
		"Entity":     domainName,
	})

	fmt.Printf("Successfully created domain %s in %s\n", domainName, domainPath)
}

func generateFile(path string, tmpl string, data map[string]string) {
	t := template.Must(template.New("").Parse(tmpl))
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", path, err)
		os.Exit(1)
	}
	defer f.Close()

	if err := t.Execute(f, data); err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		os.Exit(1)
	}
}

func InitGenDomain(rootCmd *cobra.Command) {
	rootCmd.AddCommand(domainCmd)
}
