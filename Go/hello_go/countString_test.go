package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCountStringEndpoint(t *testing.T) {
	router := SetupRouter()

	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "正常入力",
			body:       `{"text":"Hello"}`,
			wantStatus: http.StatusOK,
			wantBody:   `{"count":5}`,
		},
		{
			name:       "異常入力 - 空文字",
			body:       `{"text":""}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid request",
		},
		{
			name: "異常入力 - JSON そのものが空",
			// これだとバインド時にエラーになるので400が返るはず
			body:       ``,
			wantStatus: http.StatusBadRequest,
			// ボディ内容を固定文字列にするか、エラーメッセージの一部を含むかは実装次第
			wantBody: "Invalid request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/countstring", strings.NewReader(tt.body))
			// JSONを送信するので、Content-Typeをapplication/jsonに設定
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("status code = %d; want %d", w.Code, tt.wantStatus)
			}

			// レスポンスボディがJSONの場合、tt.wantJSONもJSONとみなして比較する方が厳密
			// ここでは文字列として比較しているが、必要に応じてjson.Unmarshalして
			// "hashed" フィールドだけを取り出して確認するなども可能
			gotBody := w.Body.String()
			if !strings.Contains(gotBody, tt.wantBody) {
				t.Errorf("body = %q; want substring %q", gotBody, tt.wantBody)
			}
		})
	}
}
