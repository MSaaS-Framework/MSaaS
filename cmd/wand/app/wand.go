package app

import (
	"github.com/spf13/cobra"
)

func NewWandCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "wand",
		Long: `

		`,
		Short: "Wand is a CLI tool for managing microservices",
	}

	// cmd.AddCommand(CreateProject())

	return cmd
}
