package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// ...existing code...
func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{"MultipleOf3", 3, "Fizz"},
		{"MultipleOf5", 5, "Buzz"},
		{"MultipleOf15", 15, "FizzBuzz"},
		{"NoMultiple", 7, "7"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := fizzBuzz(tt.input)
			if got != tt.want {
				t.Errorf("Expected '%v', got '%v'", tt.want, got)
			}
		})
	}
}

// TestFizzBuzzEndpoint は /fizzbuzz/:num エンドポイントのテスト
func TestFizzBuzzEndpoint(t *testing.T) {
	router := SetupRouter()

	// テーブルドリブンテストで複数ケースをまとめる
	tests := []struct {
		name       string
		param      string
		wantStatus int
		wantBody   string
	}{
		{"multiple of 3", "6", http.StatusOK, "Fizz"},
		{"multiple of 5", "10", http.StatusOK, "Buzz"},
		{"multiple of 15", "15", http.StatusOK, "FizzBuzz"},
		{"not multiple of 3 or 5", "7", http.StatusOK, "7"},
		{"invalid param", "abc", http.StatusBadRequest, "Invalid number: abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 擬似的なレスポンスとリクエストを用意
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/fizzbuzz/"+tt.param, nil)

			// Ginルーターにリクエストを流す
			router.ServeHTTP(w, req)

			// ステータスコードの検証
			if w.Code != tt.wantStatus {
				t.Errorf("Status code = %d; want %d", w.Code, tt.wantStatus)
			}

			// レスポンスボディの検証
			// (今回は簡単にstring比較していますが、実装や要件によっては部分一致などに変更)
			gotBody := w.Body.String()
			if gotBody != tt.wantBody {
				t.Errorf("Response body = %q; want %q", gotBody, tt.wantBody)
			}
		})
	}
}

// ...existing code...
