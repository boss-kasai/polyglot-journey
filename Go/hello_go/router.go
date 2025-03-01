package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SetupRouter はエンドポイントを設定した *gin.Engine を生成して返す
func SetupRouter() *gin.Engine {
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

	// fizzbuzzエンドポイントの定義
	router.GET("/fizzbuzz/:num", func(c *gin.Context) {
		param := c.Param("num")
		n, err := strconv.Atoi(param)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid number: %s", param)
			return
		}
		result := fizzBuzz(n)
		c.String(http.StatusOK, result)
	})

	return router
}
