package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Command-line flags for server management
	serverName := flag.String("n", "", "Server name (required)")
	serverPort := flag.String("p", "", "Server port (required)")
	protocol := flag.String("t", "http", "Protocol: http or https (default: http)")
	sslDomain := flag.String("d", "", "Domain name for SSL (required for SSL setup)")
	email := flag.String("e", "tanuedu128@gmail.com", "Email address for SSL certificate (default: tanuedu128@gmail.com)")

	flag.Parse()

	// Ensure required arguments for server management are provided
	if *serverName == "" || *serverPort == "" {
		fmt.Println("Error: Missing required arguments.")
		fmt.Println("Usage: -n <server_name> -p <server_port> [-t http|https] [-d <domain_name> for SSL setup] [-e <email_address> for SSL setup]")
		os.Exit(1)
	}

	// Manage server configuration
	err := manageServer(*serverName, *serverPort, *protocol)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// If the protocol is HTTPS, or if SSL setup is required
	if *protocol == "https" || *sslDomain != "" {
		if *sslDomain == "" {
			*sslDomain = *serverName
		}
		// Set up SSL certificate
		err := setupSSL(*sslDomain, *email)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
}
