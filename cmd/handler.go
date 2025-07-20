package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var handlerCmd = &cobra.Command{
	Use:   "handler [name]",
	Short: "Generate a new handler",
	Long: `Creates a new handler with:
- HTTP/GRPC handler implementation
- Route/Endpoint registration
- Request/Response mapping`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handlerName := args[0]
		createHandlerStructure(handlerName)
	},
}

func createHandlerStructure(handlerName string) {
	// Create handler directory
	handlerPath := filepath.Join("internal/interfaces", handlerName)
	if err := os.MkdirAll(handlerPath, 0755); err != nil {
		fmt.Printf("Error creating handler directory: %v\n", err)
		os.Exit(1)
	}

	// Generate HTTP handler
	httpTemplate := `package {{.Handler}}

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type {{.HTTPHandler}} struct {
	// Add dependencies here
}

func NewHTTPHandler() *{{.HTTPHandler}} {
	return &{{.HTTPHandler}}{}
}

func (h *{{.HTTPHandler}}) RegisterRoutes(router *gin.Engine) {
	router.POST("/{{.Route}}", h.handle)
}

func (h *{{.HTTPHandler}}) handle(c *gin.Context) {
	// Implement handler logic here
	c.JSON(http.StatusOK, gin.H{"message": "Hello from {{.Handler}}"})
}
`
	generateFile(filepath.Join(handlerPath, "http_handler.go"), httpTemplate, map[string]string{
		"Handler":     handlerName,
		"HTTPHandler": handlerName + "HTTPHandler",
		"Route":       handlerName,
	})

	// Generate GRPC handler
	grpcTemplate := `package {{.Handler}}

import (
	"context"

	"google.golang.org/grpc"
)

type {{.GRPCHandler}} struct {
	// Add dependencies here
}

func NewGRPCHandler() *{{.GRPCHandler}} {
	return &{{.GRPCHandler}}{}
}

func (h *{{.GRPCHandler}}) RegisterService(server *grpc.Server) {
	// Register GRPC service here
}

func (h *{{.GRPCHandler}}) Handle(ctx context.Context, req *Request) (*Response, error) {
	// Implement handler logic here
	return &Response{}, nil
}
`
	generateFile(filepath.Join(handlerPath, "grpc_handler.go"), grpcTemplate, map[string]string{
		"Handler":     handlerName,
		"GRPCHandler": handlerName + "GRPCHandler",
	})

	fmt.Printf("Successfully created handler %s in %s\n", handlerName, handlerPath)
}

func InitGenHandler(rootCmd *cobra.Command) {
	rootCmd.AddCommand(handlerCmd)
}
