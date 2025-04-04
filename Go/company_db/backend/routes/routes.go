package routes

import (
	"company_db/backend/controllers"

	"github.com/gin-gonic/gin"
)

// ルーティング設定
func SetupRoutes(r *gin.Engine) {
	// 企業登録エンドポイント
	r.POST("/companies", controllers.CreateCompany)
	// 企業検索エンドポイント
	r.POST("/companies/search", controllers.SearchCompany)
	// 郵便番号登録エンドポイント
	r.POST("/postal_codes", controllers.CreatePostalCode)
	// 郵便番号検索エンドポイント
	r.GET("/postal_codes", controllers.SearchPostalCode)
	// ヘルスチェック
	r.GET("/health", controllers.HealthCheck)
}
