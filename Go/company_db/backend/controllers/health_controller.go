package controllers

import (
	"net/http"

	"company_db/backend/config"

	"github.com/gin-gonic/gin"
)

// HealthCheck はサーバーの状態を確認するエンドポイント
func HealthCheck(c *gin.Context) {
	// DB 接続確認
	sqlDB, err := config.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to get database instance"})
		return
	}

	// DB に Ping を送って正常か確認
	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Database connection failed"})
		return
	}

	// サーバー & DB が正常なら OK を返す
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Server is healthy"})
}

