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
	var postalCodeID = postalCode.ID

	// Tagsを登録(既存のデータがある場合は登録しないでIDを取得)

	// 企業データを作成
	company := models.Company{
		Name:        req.Name,
		URL:         datatypes.JSON(urlJSON), // ✅ 修正：[]byte に変換して代入
		PhoneNumber: req.PhoneNumber,
		PostalCode:  postalCodeID,
		Address:     req.Address,
	}

	// データベースに保存
	if err := config.DB.Create(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusCreated, responses.CreateCompanyResponse{Message: "Company created successfully"})
}
