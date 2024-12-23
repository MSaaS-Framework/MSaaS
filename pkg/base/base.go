package base

import (
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent"
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
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

func GetDBClientFromContext(c *gin.Context) (*ent.Client, error) {
	dbClient, exists := c.Get("dbClient")
	if !exists {
		return nil, errors.New("database connection not available")
	}

	// DBClient를 ent.Client 타입으로 캐스팅
	client, ok := dbClient.(*ent.Client)
	if !ok {
		return nil, errors.New("invalid database connection")
	}
	return client, nil
}
