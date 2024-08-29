package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Root command for the CLI
var rootCmd = &cobra.Command{
	Use:   "mytool",
	Short: "MyTool is a CLI for managing Nginx configurations and SSL certificates",
	Long: `MyTool is a command-line tool for managing Nginx configurations
and setting up SSL certificates using Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default action when no subcommand is provided
		fmt.Println("Use one of the commands: manage_server, setup_nginx, setup_ssl")
	},
}

// Execute runs the root command, called by main()
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you can define global flags or configuration settings
}
