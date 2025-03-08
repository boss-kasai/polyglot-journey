package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPrime(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  []int
	}{
		{"正常型 : 2", 2, []int{2, 3}},
		{"正常型 : 10", 10, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := getPrimes(tt.input)
			if !equalSlices(got, tt.want) {
				t.Errorf("Expected '%v', got '%v'", tt.want, got)
			}
		})
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestPrimeEndpoint(t *testing.T) {
	router := SetupRouter()

	// height, weight をクエリパラメータで受け取る例
	tests := []struct {
		name       string
		param      string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "正常入力 : 2",
			param:      "2",
			wantStatus: http.StatusOK,
			wantBody:   `[2,3]`,
		},
		{
			name:       "正常入力 : 10",
			param:      "10",
			wantStatus: http.StatusOK,
			wantBody:   "[2,3,5,7,11,13,17,19,23,29]",
		},
		{
			name:       "異常入力 : ゼロ",
			param:      "0",
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid input",
		},
		{
			name:       "異常入力 : マイナス値",
			param:      "-1",
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid input",
		},
		{
			name:       "異常入力 : 数値以外",
			param:      "abc",
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid number: abc",
		},
		{
			name:       "異常入力 : 空文字",
			param:      "",
			wantStatus: http.StatusNotFound,
			wantBody:   "404 page not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/prime/"+tt.param, nil)

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
