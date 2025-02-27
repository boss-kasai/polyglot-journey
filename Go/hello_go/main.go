package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// デフォルトのミドルウェア(ロガー/リカバリ)が含まれるエンジンの生成
	router := gin.Default()

	// helloエンドポイントの定義
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// healthエンドポイントの定義
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	// サーバーの起動 (ポート8080で待ち受け)
	router.Run(":8080")
}
