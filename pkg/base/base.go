package base

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Run 함수는 주어진 Cobra 명령어를 실행하고 종료 코드를 반환합니다.
func Run(command *cobra.Command) int {
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}
	return 0
}
