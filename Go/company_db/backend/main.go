package main

import (
	"log"

	"company_db/backend/config"
	"company_db/backend/models"
	"company_db/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// データベース接続
	config.ConnectDB()

	// マイグレーション実行
	err := config.DB.AutoMigrate(&models.Company{}, &models.PostalCode{}, &models.Tag{}, &models.TagCompany{})
	if err != nil {
		log.Fatal("マイグレーションに失敗しました:", err)
	}

	log.Println("マイグレーションが完了しました！")

	// Gin ルーターを設定
	r := gin.Default()

	// ルーティングを設定
	routes.SetupRoutes(r)

	// サーバー起動
	log.Println("Server is running on port 8080...")
	r.Run(":8080") // ポート 8080 で起動
}

