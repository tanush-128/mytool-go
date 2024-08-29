package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	serverName string
	serverPort string
	protocol   string
	email      string
)

var manageServerCmd = &cobra.Command{
	Use:   "manage_server",
	Short: "Manage server configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Managing server: %s on port: %s with protocol: %s\n", serverName, serverPort, protocol)

		// Example: Run a command similar to your bash script
		err := exec.Command("sudo", "nginx", "-t").Run()
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("Nginx configuration test passed.")
		}
	},
}

func init() {
	rootCmd.AddCommand(manageServerCmd)

	manageServerCmd.Flags().StringVarP(&serverName, "name", "n", "", "Server name (required)")
	manageServerCmd.Flags().StringVarP(&serverPort, "port", "p", "", "Server port (required)")
	manageServerCmd.Flags().StringVarP(&protocol, "protocol", "t", "http", "Protocol: http or https (default: http)")
	manageServerCmd.Flags().StringVarP(&email, "email", "e", "", "Email address (required if using https)")

	manageServerCmd.MarkFlagRequired("name")
	manageServerCmd.MarkFlagRequired("port")
}
