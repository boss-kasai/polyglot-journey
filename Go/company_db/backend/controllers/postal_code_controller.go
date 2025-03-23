package controllers

import (
	"net/http"

	"company_db/backend/config"
	"company_db/backend/models"
	"company_db/backend/requests"
	"company_db/backend/responses"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func CreatePostalCode(c *gin.Context) {
	var req requests.CreatePostalCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.CreatePostalCodeErrorResponse{Error: err.Error()})
		return
	}

	postalCode := models.PostalCode{
		PostalCode: req.PostalCode,
		Address:    req.Address,
	}

	// ON CONFLICT DO NOTHING で実行
	result := config.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "postal_code"}},
		DoNothing: true,
	}).Create(&postalCode)

	// エラーがあれば500
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, responses.CreatePostalCodeErrorResponse{Error: result.Error.Error()})
		return
	}

	// 挿入されなかった（＝重複）
	if result.RowsAffected == 0 {
		// var existing models.PostalCode
		// config.DB.Where("postal_code = ?", req.PostalCode).First(&existing)

		c.JSON(http.StatusConflict, responses.CreatePostalCodeDuplicationResponse{
			Message: "この郵便番号はすでに登録されています",
			PostalCode: models.PostalCode{
				PostalCode: req.PostalCode, // ← リクエストから返す
			},
		})
		return
	}

	// 成功
	c.JSON(http.StatusCreated, responses.CreatePostalCodeResponse{
		Message:    "郵便番号を登録しました",
		PostalCode: postalCode,
	})
}
