package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var usecaseCmd = &cobra.Command{
	Use:   "usecase [name]",
	Short: "Generate a new use case",
	Long: `Creates a new use case with:
- Service interface
- Service implementation
- Request/Response models`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		usecaseName := args[0]
		createUsecaseStructure(usecaseName)
	},
}

func createUsecaseStructure(usecaseName string) {
	// Create usecase directory
	usecasePath := filepath.Join("internal/usecase", usecaseName)
	if err := os.MkdirAll(usecasePath, 0755); err != nil {
		fmt.Printf("Error creating usecase directory: %v\n", err)
		os.Exit(1)
	}

	// Generate service interface
	serviceTemplate := `package {{.Usecase}}

type {{.Service}} interface {
	Execute(req *Request) (*Response, error)
}
`
	generateFile(filepath.Join(usecasePath, "service.go"), serviceTemplate, map[string]string{
		"Usecase": usecaseName,
		"Service": usecaseName + "Service",
	})

	// Generate service implementation
	implTemplate := `package {{.Usecase}}

type service struct {
	// Add dependencies here
}

func NewService() {{.Service}} {
	return &service{}
}

func (s *service) Execute(req *Request) (*Response, error) {
	// Implement use case logic here
	return &Response{}, nil
}
`
	generateFile(filepath.Join(usecasePath, "service_impl.go"), implTemplate, map[string]string{
		"Usecase": usecaseName,
		"Service": usecaseName + "Service",
	})

	// Generate request/response models
	modelsTemplate := `package {{.Usecase}}

type Request struct {
	// Add request fields here
}

type Response struct {
	// Add response fields here
}
`
	generateFile(filepath.Join(usecasePath, "models.go"), modelsTemplate, map[string]string{
		"Usecase": usecaseName,
	})

	fmt.Printf("Successfully created usecase %s in %s\n", usecaseName, usecasePath)
}

func InitUsecase(rootCmd *cobra.Command) {
	rootCmd.AddCommand(usecaseCmd)
}
