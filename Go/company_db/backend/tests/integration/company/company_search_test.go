package company_test

import (
	"company_db/backend/config"
	"company_db/backend/models"
	"company_db/backend/requests"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func SeedCompanies(t *testing.T) {
	// --- 郵便番号を事前に登録 ---
	var postalCodes []models.PostalCode

	for i := 0; i < 5; i++ {
		pcs := models.PostalCode{
			PostalCode: fmt.Sprintf("%07d", i), // ← 7桁ゼロパディング
			Address:    fmt.Sprintf("東京都港区%d丁目", i+1),
		}
		postalCodes = append(postalCodes, pcs)
	}

	err := config.DB.Create(&postalCodes).Error
	assert.NoError(t, err)

	// --- 郵便番号 -> IDマップを作る
	pcMap := make(map[string]int)
	for _, pc := range postalCodes {
		pcMap[pc.PostalCode] = pc.ID
	}

	// --- タグを登録（IT, Web） ---
	tags := []models.Tag{
		{Name: "IT"},
		{Name: "Web"},
		{Name: "デザイン"},
		{Name: "コンサルティング"},
		{Name: "情報通信"},
		{Name: "製造業"},
		{Name: "小売業"},
	}
	err = config.DB.Create(&tags).Error
	assert.NoError(t, err)

	// --- タグ名 → IDマップを作る
	tagMap := make(map[string]string)
	for _, tag := range tags {
		tagMap[tag.Name] = tag.ID
	}

	// --- 企業を5件作成 ---
	prefixes := []string{"東京都新宿区", "埼玉県さいたま市", "岐阜県郡上市", "福岡県博多市", "北海道札幌市"}
	var companies []models.Company
	for i := 1; i <= 5; i++ {
		urls := fmt.Sprintf(`["https://example.com/%d", "https://example.org/%d"]`, i, i)
		addr := prefixes[i%len(prefixes)]

		company := models.Company{
			Name:         fmt.Sprintf("株式会社テスト_No_%d", i),
			URL:          datatypes.JSON([]byte(urls)),
			PhoneNumber:  fmt.Sprintf("090123456%02d", i),
			PostalCodeID: pcMap[fmt.Sprintf("%07d", i)],
			Address:      addr,
		}
		companies = append(companies, company)
	}

	err = config.DB.Create(&companies).Error
	assert.NoError(t, err)

	// --- 中間テーブルにタグを紐付け（例: 企業1→1,2 企業2→2,3 ...）
	var tagCompanies []models.TagCompany
	for i, c := range companies {
		// 2件ずつ割り当て（タグの数より少なくてもOK）
		tagIdx1 := i % len(tags)
		tagIdx2 := (i + 1) % len(tags)

		tagCompanies = append(tagCompanies,
			models.TagCompany{
				CompanyID: c.ID,
				TagID:     tags[tagIdx1].ID,
			},
			models.TagCompany{
				CompanyID: c.ID,
				TagID:     tags[tagIdx2].ID,
			},
		)
	}

	err = config.DB.Create(&tagCompanies).Error
	assert.NoError(t, err)
}

func TestSearchCompany_Success(t *testing.T) {
	SeedCompanies(t) // テストデータを登録

	t.Run("検索: 会社名", func(t *testing.T) {
		body := requests.SearchCompanyRequest{
			Name: "株式会社テスト_No_1",
		}
	})

	t.Run("検索: 郵便番号", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/companies?postal_code=0000001", nil)
		w := httptest.NewRecorder()
		TestRouter.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		results := decodeCompanyList(t, w.Body)
		assert.GreaterOrEqual(t, len(results), 1)
		assert.Equal(t, "株式会社テスト_No_1", results[0]["name"])
	})
}
