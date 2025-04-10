package main

import (
	"math"
	"net/http"
	"strconv"

	"unicode/utf8"

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

	// bmiエンドポイントの定義
	router.POST("/bmi", func(c *gin.Context) {
		var req struct {
			Height float64 `json:"height" binding:"required"`
			Weight float64 `json:"weight" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "Invalid request: %v", err)
			return
		}
		bmi_row := req.Weight / math.Pow(req.Height/100, 2)
		// 小数点第3位を四捨五入
		bmi := float64(int(bmi_row*100+0.5)) / 100
		// BMIの基準値で判定
		var result string
		if bmi < 18.5 {
			result = "Underweight"
		} else if bmi < 24.9 {
			result = "Normal weight"
		} else if bmi < 30 {
			result = "Overweight"
		} else {
			result = "Obesity"
		}

		c.JSON(http.StatusOK, gin.H{
			"bmi":      bmi,
			"category": result,
		})

	})

	// fiboエンドポイントの定義
	router.GET("/fibo/:num", func(c *gin.Context) {
		param := c.Param("num")
		n, err := strconv.Atoi(param)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid number: %s", param)
			return
		} else if n <= 0 {
			c.String(http.StatusBadRequest, "Invalid input")
			return
		} else if n > 100 {
			c.String(http.StatusBadRequest, "Too large number")
			return
		}
		result := fibo(n)
		c.String(http.StatusOK, result)
	})

	// primeエンドポイントの定義
	router.GET("/prime/:num", func(c *gin.Context) {
		param := c.Param("num")
		n, err := strconv.Atoi(param)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid number: %s", param)
			return
		} else if n <= 0 {
			c.String(http.StatusBadRequest, "Invalid input")
			return
		}
		result := getPrimes(n)
		c.JSON(http.StatusOK, result)
	})

	// loanエンドポイントの定義
	router.POST("/loan", func(c *gin.Context) {
		var req struct {
			Amount     float64
			years      int
			annualRate float64
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "Invalid request: %v", err)
			return
		}
		result := calculateLoanPrincipal(req.Amount, req.years, req.annualRate)
		c.JSON(http.StatusOK, gin.H{
			"principal": result,
		})
	})

	// hashエンドポイントの定義
	router.POST("/hash", func(c *gin.Context) {
		var req struct {
			Text string `json:"text" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "Invalid request: %v", err)
			return
		}
		result := HashString(req.Text)
		c.JSON(http.StatusOK, gin.H{
			"hashed": result,
		})
	})

	// 文字数を数えるエンドポイントの定義
	router.POST("/countstring", func(c *gin.Context) {
		var req struct {
			Text string `json:"text" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "Invalid request")
			return
		}
		result := utf8.RuneCountInString(req.Text)
		c.JSON(http.StatusOK, gin.H{
			"count": result,
		})
	})

	return router
}
