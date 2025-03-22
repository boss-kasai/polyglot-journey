package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHashString(t *testing.T) {
	input := "Hello, World!"
	// "Hello, World!" のSHA-256ハッシュ値（16進数）
	expected := "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f"

	actual := HashString(input)
	if actual != expected {
		t.Errorf("HashString(%q) = %q; want %q", input, actual, expected)
	}
}

func TestHashStringEndpoint(t *testing.T) {
	router := SetupRouter()

	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "正常入力",
			body:       `{"text":"Hello, World!"}`,
			wantStatus: http.StatusOK,
			wantBody:   `{"hashed":"dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f"}`,
		},
		{
			name:       "異常入力 - 空文字",
			body:       `{"text":""}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "Invalid request: Key: 'Text' Error:Field validation for 'Text' failed on the 'required' tag",
		},
		{
			name: "異常入力 - JSON そのものが空",
			// これだとバインド時にエラーになるので400が返るはず
			body:       ``,
			wantStatus: http.StatusBadRequest,
			// ボディ内容を固定文字列にするか、エラーメッセージの一部を含むかは実装次第
			wantBody: "Invalid request:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/hash", strings.NewReader(tt.body))
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
