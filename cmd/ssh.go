/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH configuration",
	Long:  `SSH will be used to connect to the container.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Erro ao carregar arquivo .env: %v", err)
		}

		user := os.Getenv("SSH_USER")
		password := os.Getenv("SSH_PASSWORD")
		host := os.Getenv("SSH_HOST")
		port := os.Getenv("SSH_PORT")

		sshConfig := &ssh.ClientConfig{
			// User: cmd.Flag("SSH").Value.String(),
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		//connect
		address := fmt.Sprintf("%s:%s", host, port)
		client, err := ssh.Dial("tcp", address, sshConfig)
		if err != nil {
			log.Fatalf("Falha ao conectar ao servidor SSH: %v", err)
		}
		defer client.Close()

		//session
		session, err := client.NewSession()
		if err != nil {
			log.Fatalf("Falha ao criar sessão SSH: %v", err)
		}
		defer session.Close()

		// Run remote cmd
		output, err := session.CombinedOutput("docker ps")
		// output, err := session.CombinedOutput(cmd.Flag("ID_CONTAINER").Value.String())
		if err != nil {
			log.Fatalf("Falha ao executar comando: %v", err)
		}

		// Print output
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
