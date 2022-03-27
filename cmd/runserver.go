package main

import (
	"github.com/againDDM/movie-warehouse/internal/server"
	"github.com/spf13/cobra"
)

func addRunServerCommand(command *cobra.Command) {
	runServerCommand := &cobra.Command{
		Use:   "runserver",
		Short: "Run server",
		Run:   func(_ *cobra.Command, _ []string) { server.Run() },
	}
	command.AddCommand(runServerCommand)
}
