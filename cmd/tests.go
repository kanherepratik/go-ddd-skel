package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	withMocks bool
)

var testsCmd = &cobra.Command{
	Use:   "tests [name]",
	Short: "Generate test stubs",
	Long: `Creates test stubs for:
- Domain entities
- Use cases
- Handlers
- Repositories`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		componentName := args[0]
		createTestStructure(componentName)

		if withMocks {
			generateMocks(componentName)
		}
	},
}

func generateMocks(componentName string) {
	// Generate mock files using mockery
	mockTemplate := `package {{.Component}}_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type Mock{{.Component}} struct {
	mock.Mock
}

// Add mock methods here
`
	generateFile(filepath.Join("internal", "mocks", componentName+"_mock.go"), mockTemplate, map[string]string{
		"Component": componentName,
	})

	fmt.Printf("Generated mocks for %s\n", componentName)
}

func createTestStructure(componentName string) {
	// Determine test file path based on component type
	var testPath string
	switch {
	case isDomain(componentName):
		testPath = filepath.Join("internal/core", componentName)
	case isUsecase(componentName):
		testPath = filepath.Join("internal/usecase", componentName)
	case isHandler(componentName):
		testPath = filepath.Join("internal/interfaces", componentName)
	default:
		fmt.Printf("Unknown component type: %s\n", componentName)
		os.Exit(1)
	}

	// Generate test file
	testTemplate := `package {{.Component}}

import (
	"testing"
)

func Test{{.Component}}(t *testing.T) {
	// Add test cases here
}
`
	generateFile(filepath.Join(testPath, componentName+"_test.go"), testTemplate, map[string]string{
		"Component": componentName,
	})

	fmt.Printf("Successfully created test stubs for %s in %s\n", componentName, testPath)
}

func isDomain(name string) bool {
	// Check if component is a domain
	_, err := os.Stat(filepath.Join("internal/core", name))
	return err == nil
}

func isUsecase(name string) bool {
	// Check if component is a usecase
	_, err := os.Stat(filepath.Join("internal/usecase", name))
	return err == nil
}

func isHandler(name string) bool {
	// Check if component is a handler
	_, err := os.Stat(filepath.Join("internal/interfaces", name))
	return err == nil
}

func InitTests(rootCmd *cobra.Command) {
	rootCmd.AddCommand(testsCmd)
	testsCmd.Flags().BoolVarP(&withMocks, "with-mocks", "m", false, "Generate mock implementations")
}
