/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var input_ssh string
var input_id_container string

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Input SSH",
	Long:  "SSH will be used to connect to the container.",
	Run: func(cmd *cobra.Command, args []string) {
		ssh := cmd.Flag("SSH").Value.String()
		fmt.Print(ssh)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&input_ssh, "SSH", "s", "solus_web@192.168.1.147", "SSH to connect to the container")
	cliCmd.Flags().StringVarP(&input_id_container, "ID_CONTAINER", "i", "container_id", "Container ID")
}
