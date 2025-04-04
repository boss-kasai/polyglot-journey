package company_test

import (
	"company_db/backend/config"
	"company_db/backend/models"
	"company_db/backend/routes"
	"company_db/backend/tests/testutils"
	"context"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/testcontainers/testcontainers-go"
)

var TestRouter *gin.Engine
var TestContainer testcontainers.Container

func TestMain(m *testing.M) {
	dsn, container, err := testutils.SetupPostgresContainer()
	if err != nil {
		panic(err)
	}
	TestContainer = container

	_, err = config.ConnectDatabase(dsn)
	if err != nil {
		panic("DB接続に失敗: " + err.Error())
	}

	config.DB.AutoMigrate(&models.PostalCode{})
	config.DB.AutoMigrate(&models.Company{})
	config.DB.AutoMigrate(&models.Tag{})
	config.DB.AutoMigrate(&models.TagCompany{})

	gin.SetMode(gin.TestMode)
	TestRouter = gin.Default()
	routes.SetupRoutes(TestRouter)

	// テスト実行
	code := m.Run()

	// defer の代わりにここで terminate
	TestContainer.Terminate(context.Background())

	os.Exit(code)
}
