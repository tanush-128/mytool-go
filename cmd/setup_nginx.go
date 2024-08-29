package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setupNginxCmd = &cobra.Command{
	Use:   "setup_nginx",
	Short: "Set up Nginx configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Setting up Nginx for server: %s on port: %s with protocol: %s\n", serverName, serverPort, protocol)

		// Implementation of setting up Nginx configuration
		// Example: Modify Nginx configuration file, reload Nginx, etc.
	},
}

func init() {
	rootCmd.AddCommand(setupNginxCmd)

	setupNginxCmd.Flags().StringVarP(&serverName, "name", "n", "", "Server name (required)")
	setupNginxCmd.Flags().StringVarP(&serverPort, "port", "p", "", "Server port (required)")
	setupNginxCmd.Flags().StringVarP(&protocol, "protocol", "t", "http", "Protocol: http or https (default: http)")

	setupNginxCmd.MarkFlagRequired("name")
	setupNginxCmd.MarkFlagRequired("port")
}
