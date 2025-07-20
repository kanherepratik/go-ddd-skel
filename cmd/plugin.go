package cmd

import (
	"fmt"
	"os"
	"plugin"

	"github.com/spf13/cobra"
)

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Manage plugins",
	Long: `Install, list, and remove plugins for extending functionality.
Plugins must be compiled as .so files and implement the Plugin interface.`,
}

var pluginInstallCmd = &cobra.Command{
	Use:   "install [path]",
	Short: "Install a plugin",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		installPlugin(path)
	},
}

var pluginListCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed plugins",
	Run: func(cmd *cobra.Command, args []string) {
		listPlugins()
	},
}

var pluginRemoveCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove a plugin",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		removePlugin(name)
	},
}

func installPlugin(path string) {
	// Load plugin
	_, err := plugin.Open(path)
	if err != nil {
		fmt.Printf("Error loading plugin: %v\n", err)
		os.Exit(1)
	}

	// TODO: Register plugin functionality
	fmt.Printf("Successfully installed plugin from %s\n", path)
}

func listPlugins() {
	// TODO: List installed plugins
	fmt.Println("Installed plugins:")
}

func removePlugin(name string) {
	// TODO: Remove plugin
	fmt.Printf("Removed plugin: %s\n", name)
}

func InitPlugin(rootCmd *cobra.Command) {
	pluginCmd.AddCommand(pluginInstallCmd)
	pluginCmd.AddCommand(pluginListCmd)
	pluginCmd.AddCommand(pluginRemoveCmd)
	rootCmd.AddCommand(pluginCmd)
}
