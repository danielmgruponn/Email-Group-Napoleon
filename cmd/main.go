package main

import (
	"napoleon-email/cmd/server"
	"napoleon-email/src/config/app"
	"napoleon-email/src/pkg/logger"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	app.LoadEnv()
	if err := app.ValidateEnvironment(); err != nil {
		logger.LogError("environment validation failed", err, logger.LogStruct{Action: "environment_validation_failed", User: 0})
		os.Exit(1)
	}
	rootCmd := newRootCmd()
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new-napoleon",
		Short: "Back Sender Email Grupo Napoleon",
		Long:  "A backend server that is used to send emails to users and forward those emails to the intended recipients.",
	}
	cmd.AddCommand(server.RunServerCmd)
	return cmd
}
