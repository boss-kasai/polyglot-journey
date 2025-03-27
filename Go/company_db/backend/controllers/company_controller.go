package controllers

import (
	"encoding/json"
	"net/http"

	"company_db/backend/config"
	"company_db/backend/models"
	"company_db/backend/requests"
	"company_db/backend/responses"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// 企業データ登録エンドポイント
func CreateCompany(c *gin.Context) {
	var req requests.CreateCompanyRequest

	// リクエストの JSON をバインド
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// []string を JSON に変換
	urlJSON, err := json.Marshal(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert URL to JSON"})
		return
	}

	// postal_codeのIDを取得
	var postalCode models.PostalCode
	if err := config.DB.Where("postal_code = ?", req.PostalCode).First(&postalCode).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Postal code not found"})
		return
	}

	// 企業データを作成
	company := models.Company{
		Name:         req.Name,
		URL:          datatypes.JSON(urlJSON), // ✅ 修正：[]byte に変換して代入
		PhoneNumber:  req.PhoneNumber,
		PostalCodeID: postalCode.ID,
		Address:      req.Address,
	}

	err = config.DB.Transaction(func(tx *gorm.DB) error {
		// トランザクション内は常に tx を使う！
		if err := tx.Create(&company).Error; err != nil {
			return err
		}

		for _, tagName := range req.Tags {
			var tag models.Tag
			if err := tx.Where("name = ?", tagName).First(&tag).Error; err != nil {
				tag = models.Tag{Name: tagName}
				if err := tx.Create(&tag).Error; err != nil {
					return err
				}
			}
			if err := tx.Create(&models.TagCompany{
				TagID:     tag.ID,
				CompanyID: company.ID,
			}).Error; err != nil {
				return err
			}
		}

		return nil // すべて成功 → commit
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusCreated, responses.CreateCompanyResponse{Message: "Company created successfully"})
}
