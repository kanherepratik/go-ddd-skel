package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var archCmd = &cobra.Command{
	Use:   "arch",
	Short: "Visualize architecture",
	Long: `Generates a visualization of the project's architecture using go-callvis.
Creates a dependency graph showing relationships between packages and components.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateArchitectureGraph()
	},
}

func generateArchitectureGraph() {
	// Check if go-callvis is installed
	if _, err := exec.LookPath("go-callvis"); err != nil {
		fmt.Println("go-callvis not found. Please install it first:")
		fmt.Println("go install github.com/ofabry/go-callvis@latest")
		os.Exit(1)
	}

	// Generate architecture graph
	cmd := exec.Command("go-callvis", "-focus", ".", "-group", "pkg,type", "-http", ":7878")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error generating architecture graph: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Architecture visualization available at http://localhost:7878")
}

func InitGraphArch(rootCmd *cobra.Command) {
	rootCmd.AddCommand(archCmd)
}
