package postal_code_test

import (
	"company_db/backend/config"
	"company_db/backend/models"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func decodePostalCodeList(t *testing.T, body io.Reader) []map[string]interface{} {
	var resp map[string]interface{}
	err := json.NewDecoder(body).Decode(&resp)
	assert.NoError(t, err)

	rawList := resp["postal_codes"].([]interface{})
	list := make([]map[string]interface{}, len(rawList))
	for i, item := range rawList {
		list[i] = item.(map[string]interface{}) // 型変換をまとめる
	}
	return list
}

func TestSearchPostalCode(t *testing.T) {
	// --- データ挿入 ---
	pc := models.PostalCode{
		PostalCode: "1112222",
		Address:    "北海道札幌市",
	}
	err := config.DB.Create(&pc).Error
	assert.NoError(t, err)

	t.Run("検索: 郵便番号だけ", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/postal_codes?postal_code=111", nil)
		w := httptest.NewRecorder()
		TestRouter.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		results := decodePostalCodeList(t, w.Body)
		assert.GreaterOrEqual(t, len(results), 1)
		assert.Equal(t, "1112222", results[0]["postal_code"])
		assert.Equal(t, "北海道札幌市", results[0]["address"])
	})

	t.Run("検索: 住所だけ", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/postal_codes?address=札幌", nil)
		w := httptest.NewRecorder()
		TestRouter.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		results := decodePostalCodeList(t, w.Body)
		assert.GreaterOrEqual(t, len(results), 1)
		assert.Equal(t, "1112222", results[0]["postal_code"])
		assert.Equal(t, "北海道札幌市", results[0]["address"])
	})

	t.Run("検索: 郵便番号 + 住所", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/postal_codes?postal_code=222&address=札幌", nil)
		w := httptest.NewRecorder()
		TestRouter.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		results := decodePostalCodeList(t, w.Body)
		assert.GreaterOrEqual(t, len(results), 1)
		assert.Equal(t, "1112222", results[0]["postal_code"])
		assert.Equal(t, "北海道札幌市", results[0]["address"])
	})

	t.Run("検索: 該当なし", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/postal_codes?postal_code=333", nil)
		w := httptest.NewRecorder()
		TestRouter.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		results := decodePostalCodeList(t, w.Body)
		assert.Equal(t, 0, len(results))
	})
}

func TestSearchPostalCode_Error(t *testing.T) {
	t.Run("検索: 条件なし", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/postal_codes", nil)
		w := httptest.NewRecorder()
		TestRouter.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
