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

	// データベースに保存
	if err := config.DB.Create(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
		return
	}

	// タグを登録
	for _, tagName := range req.Tags {
		var tag models.Tag
		if err := config.DB.Where("name = ?", tagName).First(&tag).Error; err != nil {
			// タグが存在しない場合は新規作成
			tag = models.Tag{Name: tagName}
			if err := config.DB.Create(&tag).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag"})
				return
			}
		}
		// 中間テーブルに保存
		err := config.DB.Create(&models.TagCompany{TagID: tag.ID, CompanyID: company.ID}).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag_company"})
			return
		}
	}

	// 成功レスポンス
	c.JSON(http.StatusCreated, responses.CreateCompanyResponse{Message: "Company created successfully"})
}
