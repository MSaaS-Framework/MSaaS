package app

import (
	"fmt"
	"net/http"
)

// HealthzHandler는 healthz 엔드포인트를 처리합니다.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

// RootHandler는 루트 경로를 처리합니다.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Wizcraft API 서버에 오신 것을 환영합니다.")
}
