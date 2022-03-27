package main

import (
	"github.com/spf13/cobra"
	"os"
)

func main() {
	rootCommand := &cobra.Command{
		Short: "movie warehouse",
	}
	addRunServerCommand(rootCommand)
	addPingServerCommand(rootCommand)
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
