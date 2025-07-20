package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var dxCmd = &cobra.Command{
	Use:   "dx",
	Short: "Developer experience tools",
	Long: `Setup and manage developer experience features:
- Linting
- Live reload
- Telemetry`,
}

var setupLintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Setup linting",
	Run: func(cmd *cobra.Command, args []string) {
		setupLinting()
	},
}

var setupAirCmd = &cobra.Command{
	Use:   "air",
	Short: "Setup live reload",
	Run: func(cmd *cobra.Command, args []string) {
		setupAir()
	},
}

var setupTelemetryCmd = &cobra.Command{
	Use:   "telemetry",
	Short: "Setup telemetry",
	Run: func(cmd *cobra.Command, args []string) {
		setupTelemetry()
	},
}

func setupLinting() {
	// Install golangci-lint
	cmd := exec.Command("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error installing golangci-lint: %v\n", err)
		os.Exit(1)
	}

	// Create config file
	config := `run:
  timeout: 5m
  modules-download-mode: readonly

linters:
  enable:
    - gocritic
    - govet
    - staticcheck
`
	if err := os.WriteFile(".golangci.yml", []byte(config), 0644); err != nil {
		fmt.Printf("Error creating lint config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Linting setup complete. Use 'golangci-lint run' to lint your code.")
}

func setupAir() {
	// Install air
	cmd := exec.Command("go", "install", "github.com/cosmtrek/air@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error installing air: %v\n", err)
		os.Exit(1)
	}

	// Create air config
	config := `[build]
  cmd = "go build -o ./tmp/main ."
  bin = "./tmp/main"
  log = "build-errors.log"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor"]
  include_ext = ["go", "tpl", "tmpl", "html"]
`
	if err := os.WriteFile("air.toml", []byte(config), 0644); err != nil {
		fmt.Printf("Error creating air config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Air setup complete. Use 'air' to start live reload.")
}

func setupTelemetry() {
	// Install OpenTelemetry dependencies
	cmd := exec.Command("go", "get", "go.opentelemetry.io/otel")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error installing OpenTelemetry: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Telemetry setup complete. Add instrumentation to your code.")
}

func InitSetupDX(rootCmd *cobra.Command) {
	dxCmd.AddCommand(setupLintCmd)
	dxCmd.AddCommand(setupAirCmd)
	dxCmd.AddCommand(setupTelemetryCmd)
	rootCmd.AddCommand(dxCmd)
}
