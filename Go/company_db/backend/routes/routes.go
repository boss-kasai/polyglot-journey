package routes

import (
	"company_db/backend/controllers"

	"github.com/gin-gonic/gin"
)

// ルーティング設定
func SetupRoutes(r *gin.Engine) {
	// 企業登録エンドポイント
	r.POST("/companies", controllers.CreateCompany)
	// 郵便番号登録エンドポイント
	r.POST("/postal_codes", controllers.CreatePostalCode)
	// ヘルスチェック
	r.GET("/health", controllers.HealthCheck)
}
