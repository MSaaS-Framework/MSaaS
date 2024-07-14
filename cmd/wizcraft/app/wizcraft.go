package app

import (
	"github.com/spf13/cobra"
)

func NewWizcraftCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "wizcraft",
		Long: `

		`,
		Short: "Wizcraft is a CLI tool for managing microservices",
	}

	// cmd.AddCommand(CreateProject())

	return cmd
}
