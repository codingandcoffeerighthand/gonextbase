package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	version    string
	commitHash string
)

const (
	flagConfigFilePath = "config-file-path"
)

func server() *cobra.Command {
	command := &cobra.Command{
		Use:   "server",
		Short: "Starts the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
	return command
}

func main() {
	rootCommand := &cobra.Command{
		Version: fmt.Sprintf("%s-%s", version, commitHash),
	}
	rootCommand.AddCommand(
		server(),
	)
	if err := rootCommand.Execute(); err != nil {
		log.Panic(err)
	}
}
