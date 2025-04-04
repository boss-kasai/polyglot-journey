package company_test

import (
	"bytes"
	"company_db/backend/config"
	"company_db/backend/models"
	"company_db/backend/requests"
	"company_db/backend/responses"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompaniesTableExists(t *testing.T) {
	exists := config.DB.Migrator().HasTable("companies")
	assert.True(t, exists, "companies テーブルが存在しません")
}

func TestCreateCompany_Success(t *testing.T) {
	body := requests.CreateCompanyRequest{
		Name:        "株式会社テスト",
		URL:         []string{"https://example.com", "https://example.org"},
		PhoneNumber: "09012345678",
		PostalCode:  "0123456",
		Address:     "東京都港区",
		Tags:        []string{"IT", "Web"},
	}
	jsonBody, _ := json.Marshal(body) // JSONに変換

	// 事前に郵便番号を登録
	postalCodeBody := requests.CreatePostalCodeRequest{
		PostalCode: "0123456",
		Address:    "東京都港区",
	}
	postalCodeJsonBody, _ := json.Marshal(postalCodeBody)
	postalCodeReq, _ := http.NewRequest("POST", "/postal_codes", bytes.NewBuffer(postalCodeJsonBody))
	postalCodeReq.Header.Set("Content-Type", "application/json")
	postalCodeW := httptest.NewRecorder()
	TestRouter.ServeHTTP(postalCodeW, postalCodeReq)
	assert.Equal(t, http.StatusCreated, postalCodeW.Code)

	// 会社情報を登録
	req, _ := http.NewRequest("POST", "/companies", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	TestRouter.ServeHTTP(w, req)
	var resp responses.CreateCompanyResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.NoError(t, err, "レスポンスのJSONパースに失敗しました")
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "Company created successfully", resp.Message)

	// DBが正しく挿入されたかを確認
	var result models.Company
	db_err := config.DB.
		Preload("PostalCode"). // 外部キーの郵便番号
		Preload("Tags").       // 多対多のタグ
		Where("name = ?", "株式会社テスト").
		First(&result).Error

	assert.NoError(t, db_err)
	assert.Equal(t, "株式会社テスト", result.Name)
	assert.Equal(t, "09012345678", result.PhoneNumber)
	assert.Equal(t, "0123456", result.PostalCode.PostalCode)
	assert.Equal(t, "東京都港区", result.Address)
	assert.Equal(t, 2, len(result.Tags))

	tagNames := make([]string, len(result.Tags))
	for i, tag := range result.Tags {
		tagNames[i] = tag.Name
	}
	assert.ElementsMatch(t, []string{"IT", "Web"}, tagNames) // 順番は関係なしで要素が一致しているか

}

// Name の値が空の場合のテスト
func TestCreateCompany_Fail_NoName(t *testing.T) {
	// Name が空の場合
	body := requests.CreateCompanyRequest{
		Name:        "",
		URL:         []string{"https://example.com", "https://example.org"},
		PhoneNumber: "09012345678",
		PostalCode:  "0123456",
		Address:     "東京都港区",
		Tags:        []string{"IT", "Web"},
	}
	jsonBody, _ := json.Marshal(body) // JSONに変換

	// 会社情報を登録
	req, _ := http.NewRequest("POST", "/companies", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	TestRouter.ServeHTTP(w, req)
	var resp responses.CreateCompanyErrorResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.NoError(t, err, "レスポンスのJSONパースに失敗しました")
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Key: 'CreateCompanyRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag", resp.Error)
}

// 郵便番号が事前登録されたものでない
func TestCreateCompany_Fail_NoPostalCode(t *testing.T) {
	body := requests.CreateCompanyRequest{
		Name:        "株式会社テスト",
		URL:         []string{"https://example.com", "https://example.org"},
		PhoneNumber: "09012345678",
		PostalCode:  "0023456",
		Address:     "東京都港区",
		Tags:        []string{"IT", "Web"},
	}
	jsonBody, _ := json.Marshal(body) // JSONに変換

	// 会社情報を登録
	req, _ := http.NewRequest("POST", "/companies", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	TestRouter.ServeHTTP(w, req)
	var resp responses.CreateCompanyErrorResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.NoError(t, err, "レスポンスのJSONパースに失敗しました")
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Postal code not found", resp.Error)
}
