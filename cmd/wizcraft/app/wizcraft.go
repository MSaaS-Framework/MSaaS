package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// NewWizcraftCommand는 루트 명령어를 생성합니다.
func NewWizcraftCommand() *cobra.Command {
	// 루트 명령어 정의
	cmd := &cobra.Command{
		Use:   "wizcraft",
		Short: "Wizcraft is a backend server for handling MSaaS API",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please provide a subcommand")
		},
	}

	// start 서브 명령어 추가
	cmd.AddCommand(newStartCommand())

	return cmd
}

// newStartCommand는 서버를 시작하는 명령어를 생성합니다.
func newStartCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the Wizcraft server",
		Run: func(cmd *cobra.Command, args []string) {
			StartServer()
		},
	}
}

// StartServer는 웹 서버를 시작하는 함수입니다.
func StartServer() {

	RegisterRoutes()

	// 등록된 모든 핸들러 적용
	for _, handler := range HttpHandlers {
		// 모든 메서드를 처리하되, 핸들러 내에서 메서드 구분
		http.HandleFunc(handler.Path, func(w http.ResponseWriter, r *http.Request) {
			if r.Method == handler.Method {
				handler.Function(w, r)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		})
	}

	log.Println("서버가 8080 포트에서 실행 중입니다.")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("서버 시작 실패: %v", err)
	}
}
