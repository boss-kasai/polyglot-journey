package integration

import (
	"bytes"
	"company_db/backend/config"
	"company_db/backend/requests"
	"company_db/backend/responses"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// var postalCodeContainer testcontainers.Container
// var postalCodeRouter *gin.Engine

func TestPostalCodesTableExists(t *testing.T) {
	exists := config.DB.Migrator().HasTable("postal_codes")
	assert.True(t, exists, "postal_codes テーブルが存在しません")
}

// 成功パターンのテスト
func TestCreatePostalCode_Success(t *testing.T) {
	body := requests.CreatePostalCodeRequest{
		PostalCode: "1234567",
		Address:    "東京都港区",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/postal_codes", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	TestRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp responses.CreatePostalCodeResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.NoError(t, err, "レスポンスのJSONパースに失敗しました")

	assert.Equal(t, "1234567", resp.PostalCode.PostalCode)
	assert.Equal(t, "東京都港区", resp.PostalCode.Address)
	assert.Equal(t, "郵便番号を登録しました", resp.Message)
}

// 失敗パターンのテスト
func TestCreatePostalCode_MissingField(t *testing.T) {
	body := requests.CreatePostalCodeRequest{
		Address: "東京都港区", // postal_code がない
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/postal_codes", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	TestRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error"`)
}

func TestCreatePostalCode_Duplicate(t *testing.T) {
	body := requests.CreatePostalCodeRequest{
		PostalCode: "9998888",
		Address:    "大阪府大阪市",
	}
	jsonBody, _ := json.Marshal(body)

	// 1回目：成功
	req1, _ := http.NewRequest("POST", "/postal_codes", bytes.NewBuffer(jsonBody))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	TestRouter.ServeHTTP(w1, req1)
	assert.Equal(t, http.StatusCreated, w1.Code)

	// 2回目：重複
	req2, _ := http.NewRequest("POST", "/postal_codes", bytes.NewBuffer(jsonBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	TestRouter.ServeHTTP(w2, req2)

	// 重複時のレスポンスを検証
	var resp responses.CreatePostalCodeDuplicationResponse
	err := json.NewDecoder(w2.Body).Decode(&resp)
	assert.NoError(t, err, "レスポンスのJSONパースに失敗しました")

	assert.Equal(t, http.StatusConflict, w2.Code)
	assert.Contains(t, "この郵便番号はすでに登録されています", resp.Message)
	assert.Equal(t, "9998888", resp.PostalCode.PostalCode)

}
