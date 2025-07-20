package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

type ProjectConfig struct {
	Router   string
	Logger   string
	Database string
	Cache    string
	UseRedis bool
	UseKafka bool
	UseGRPC  bool
}

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
		config := interactiveSetup()
		createProjectStructure(projectName, &config)
	},
}

func interactiveSetup() ProjectConfig {
	var config ProjectConfig

	// Router selection
	survey.AskOne(&survey.Select{
		Message: "Choose your router:",
		Options: []string{"net/http", "gin", "echo", "chi"},
	}, &config.Router)

	// Logger selection
	survey.AskOne(&survey.Select{
		Message: "Choose your logger:",
		Options: []string{"log", "logrus", "zap", "zerolog"},
	}, &config.Logger)

	// Database selection
	survey.AskOne(&survey.Select{
		Message: "Choose your database:",
		Options: []string{"none", "postgres", "mysql", "mongodb"},
	}, &config.Database)

	// Cache selection
	survey.AskOne(&survey.Select{
		Message: "Choose your cache:",
		Options: []string{"none", "in-memory", "redis"},
	}, &config.Cache)

	// Additional services
	survey.AskOne(&survey.Confirm{
		Message: "Use Redis?",
	}, &config.UseRedis)

	survey.AskOne(&survey.Confirm{
		Message: "Use Kafka?",
	}, &config.UseKafka)

	survey.AskOne(&survey.Confirm{
		Message: "Use gRPC?",
	}, &config.UseGRPC)

	return config
}

func createProjectStructure(projectName string, config *ProjectConfig) {
	// Create basic project structure
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

	// Initialize Go module
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = projectName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error initializing Go module: %v\n", err)
		os.Exit(1)
	}

	// Create main.go file based on selected router
	var mainContent string
	switch config.Router {
	case "gin":
		mainContent = `package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	r.Run()
}
`
	case "echo":
		mainContent = `package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World!")
	})
	e.Start(":8080")
}
`
	case "chi":
		mainContent = `package main

import (
	"net/http"
	
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":8080", r)
}
`
	default: // net/http
		mainContent = `package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from %s!", r.URL.Path[1:])
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
`
	}

	mainPath := filepath.Join(projectName, "main.go")
	if err := os.WriteFile(mainPath, []byte(mainContent), 0644); err != nil {
		fmt.Printf("Error creating main.go: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully created DDD project structure in %s/ with Go module initialized and main.go created\n", projectName)
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(initCmd)
}
