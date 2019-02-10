package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
var subject string
var body string
var templatefile string
// RootCommand - this is app root command
var RootCommand = &cobra.Command{
	Use: "sm",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to mail-client..")
	},
}

func Execute() {
	fmt.Println("in execute method..")
	HelloCommand.Flags().StringVarP(&subject, "subject", "s", "Default Subject", "Notification subject")
	HelloCommand.Flags().StringVarP(&body, "body", "b", "Default Body", "Notification body")
	HelloCommand.Flags().StringVarP(&templatefile, "templatefile", "t", "default_template.html", "Notification body template")
	RootCommand.AddCommand(HelloCommand)
	RootCommand.Execute()

}
