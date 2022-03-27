package main

import (
	"github.com/againDDM/movie-warehouse/internal/server"
	"github.com/spf13/cobra"
)

func addPingServerCommand(command *cobra.Command) {
	pingServerCommand := &cobra.Command{
		Use:   "ping",
		Short: "Ping server",
		RunE:  func(_ *cobra.Command, _ []string) error { return server.Ping() },
	}
	command.AddCommand(pingServerCommand)
}
