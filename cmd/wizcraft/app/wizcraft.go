package app

import (
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent"
	"MSaaS-Framework/MSaaS/pkg/crub"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

func init() {
	// .env 파일 로드
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println(".env 파일을 로드할 수 없습니다. 기본 환경 변수를 사용합니다.")
	}
	fmt.Println("환경 변수 로드 완료")
}

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

	// crud 서브 명령어 추가
	cmd.AddCommand(newCrudCommand())

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

// newCrudCommand는 리소스에 대한 CRUD 생성 제너레이터 입니다
func newCrudCommand() *cobra.Command {
	// 리소스 이름과 파일 경로를 담을 변수를 선언
	var resource, filePath string

	// cobra.Command를 생성
	cmd := &cobra.Command{
		Use:   "crud",
		Short: "Create CRUD handlers",
		Run: func(cmd *cobra.Command, args []string) {
			// 플래그로부터 값을 받아 crub.Run 함수에 전달
			if resource == "" || filePath == "" {
				fmt.Println("Both resource (-r) and file path (-f) flags are required.")
				return
			}
			crub.Run(resource, filePath)
		},
	}

	// -r (resource) 플래그 추가
	cmd.Flags().StringVarP(&resource, "resource", "r", "", "Name of the resource to generate CRUD handlers for")
	// -f (filePath) 플래그 추가
	cmd.Flags().StringVarP(&filePath, "filepath", "f", "", "File path to save the generated handlers")

	return cmd
}

// StartServer는 웹 서버를 시작하는 함수입니다.
func StartServer() {

	// 환경 변수에서 설정값을 가져옵니다.
	host := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	webPort := os.Getenv("WIZCRAFT_PORT")

	println(host, dbPort, user, dbname, password, webPort)

	// ent 초기화 - 추후에 env를 사용하여 데이터를 가져오도록 수정
	// PostgreSQL 연결 문자열을 구성합니다.
	postgresDSN := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, dbPort, user, dbname, password)

	// ent 초기화
	client, err := ent.Open("postgres", postgresDSN)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Gin의 새로운 라우터 생성
	router := gin.New()

	// 미들웨어 추가 (필요에 따라 로깅, 복구 등 추가 가능)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 라우터에 CRUD 경로를 등록
	RegisterRoutes(router)

	// 웹 서버 시작
	log.Printf("반갑습니다. 서버가 포트 %s에서 실행 중입니다.", webPort)
	if err := router.Run(":" + webPort); err != nil {
		log.Fatalf("서버 시작 실패: %v", err)
	}
}
