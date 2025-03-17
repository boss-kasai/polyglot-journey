package controllers

import (
	"net/http"
	"strings"

	"company_db/backend/config"
	"company_db/backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// リクエスト用構造体
type CreatePostalCodeRequest struct {
	PostalCode string `json:"postal_code" binding:"required"`
	Address    string `json:"address" binding:"required"`
}

func CreatePostalCode(c *gin.Context) {
	var req CreatePostalCodeRequest

	// リクエストの JSON をバインド
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 郵便番号データを作成
	postalCode := models.PostalCode{
		PostalCode: req.PostalCode,
		Address:    req.Address,
	}

	// データベースに保存
	if err := config.DB.Create(&postalCode).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Postal code already exists"})
			return
		}
		// `ON CONFLICT DO NOTHING` を適用
		result := config.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "postal_code"}},
			DoNothing: true,
		}).Create(&postalCode)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		if result.RowsAffected == 0 {
			// 既存のデータがあった場合
			c.JSON(http.StatusConflict, gin.H{
				"message":     "この郵便番号はすでに登録されています",
				"postal_code": postalCode.PostalCode,
			})
			return
		}

		// 新規登録成功
		c.JSON(http.StatusCreated, gin.H{
			"message":     "郵便番号を登録しました",
			"postal_code": postalCode.PostalCode,
		})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusCreated, gin.H{"message": "Postal code created successfully", "postal_code": postalCode})
}
