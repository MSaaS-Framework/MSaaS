package base

import (
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command) int {
	cmd.Execute()
	return 0
}
