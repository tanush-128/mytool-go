package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var setupSSLCmd = &cobra.Command{
	Use:   "setup_ssl",
	Short: "Set up SSL certificates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Setting up SSL for domain: %s with email: %s\n", serverName, email)

		// Example: Run certbot command to set up SSL
		err := exec.Command("sudo", "certbot", "--nginx", "-d", serverName, "--non-interactive", "--agree-tos", "-m", email, "--redirect").Run()
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("SSL certificate setup successful.")
		}
	},
}

func init() {
	rootCmd.AddCommand(setupSSLCmd)

	setupSSLCmd.Flags().StringVarP(&serverName, "domain", "d", "", "Domain name (required)")
	setupSSLCmd.Flags().StringVarP(&email, "email", "e", "", "Email address (required)")

	setupSSLCmd.MarkFlagRequired("domain")
	setupSSLCmd.MarkFlagRequired("email")
}
