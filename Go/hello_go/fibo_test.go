package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFiboEndpoint(t *testing.T) {
	router := SetupRouter()

	// height, weight をクエリパラメータで受け取る例
	tests := []struct {
		name       string
		param      string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "正常入力 - 2",
			param:      "2",
			wantStatus: http.StatusOK,
			wantBody:   "1",
		},
		{
			name:       "正常入力 - 10",
			param:      "10",
			wantStatus: http.StatusOK,
			wantBody:   "55",
		},
		{
			name:       "異常入力 - ゼロ",
			param:      "0",
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid input",
		},
		{
			name:       "異常入力 - マイナス値",
			param:      "-1",
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid input",
		},
		{
			name:       "異常入力 - 数値以外",
			param:      "abc",
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid number: abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/fibo/"+tt.param, nil)

			// Ginルーターにリクエストを流す
			router.ServeHTTP(w, req)

			// ステータスコードの検証
			if w.Code != tt.wantStatus {
				t.Errorf("Status code = %d; want %d", w.Code, tt.wantStatus)
			}

			// レスポンスボディの検証
			gotBody := w.Body.String()
			if gotBody != tt.wantBody {
				t.Errorf("Response body = %q; want %q", gotBody, tt.wantBody)
			}
		})
	}
}
