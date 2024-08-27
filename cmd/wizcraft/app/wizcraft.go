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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Wizcraft API Server")
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
