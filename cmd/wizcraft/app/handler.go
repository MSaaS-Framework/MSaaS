package app

import (
	"net/http"
)

// HttpHandlers는 여러 핸들러를 관리하는 슬라이스입니다.
var HttpHandlers []HttpHandler

// HttpHandler 구조체는 각 HTTP 경로와 요청 방식을 관리합니다.
type HttpHandler struct {
	Method   string
	Path     string
	Function func(w http.ResponseWriter, r *http.Request)
}

// 등록할 핸들러를 추가하는 함수 예시
func RegisterRoutes() {
	// 예시로 healthz 엔드포인트 등록
	HttpHandlers = append(HttpHandlers, HttpHandler{
		Method:   "GET",
		Path:     "/healthz",
		Function: HealthzHandler,
	})

	// 또 다른 예시 핸들러 등록 가능
	HttpHandlers = append(HttpHandlers, HttpHandler{
		Method:   "GET",
		Path:     "/",
		Function: RootHandler,
	})
}
