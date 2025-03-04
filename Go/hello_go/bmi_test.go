package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBmiEndpoint(t *testing.T) {
	router := SetupRouter()

	// height, weight をクエリパラメータで受け取る例
	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "正常入力",
			body:       `{"height": 169.5, "weight": 60}`,
			wantStatus: http.StatusOK,
			wantBody:   `{"bmi":20.88,"category":"Normal weight"}`,
		},
		{
			name:       "必須項目不足",
			body:       `{"weight": 60}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   `Invalid request: Key: 'Height' Error:Field validation for 'Height' failed on the 'required' tag`,
		},
		{
			name:       "値が0以下になっている",
			body:       `{"height": 169.5, "weight": 0}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   `Invalid request: Key: 'Weight' Error:Field validation for 'Weight' failed on the 'required' tag`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/bmi", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Status code = %d; want %d", w.Code, tt.wantStatus)
			}
			gotBody := w.Body.String()
			if gotBody != tt.wantBody {
				t.Errorf("Response body = %q; want %q", gotBody, tt.wantBody)
			}
		})
	}
}
