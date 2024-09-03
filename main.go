package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"learning-golang/api"
	"learning-golang/constants"
	"learning-golang/db"
	"learning-golang/docs"
	"learning-golang/router"
	"os"
)

// @title			Swagger Example API
// @version		1.0
// @description	This is a sample server.
// @host			localhost:8080
// @BasePath		/api/v1
func main() {
	// .env ファイルを読み込む
	err := loadEnv()
	if err != nil {
		// エラーメッセージをクライアントに返す
		handleError(err)
		return
	}

	// データベースの初期化
	database, err := db.InitDB()
	if err != nil {
		// エラーメッセージをクライアントに返す
		handleError(err)
		return
	}

	// Ginのルーターのセットアップ
	r := router.SetupRouter(database)

	// APIグループの設定
	setupAPIRoutes(r, database)

	// Swaggerの設定
	setupSwagger(r)

	// サーバーの起動
	appPort := os.Getenv("APP_PORT")
	r.Run(":" + appPort)
}

// .env ファイルを読み込む
func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("%s: %v", constants.ErrLoadEnvFile, err)
	}
	return nil
}

func handleError(err error) {
	gin.DefaultErrorWriter.Write([]byte(err.Error()))
}

func setupAPIRoutes(r *gin.Engine, db *gorm.DB) {
	v1 := r.Group("/api/v1")
	api.SetupV1Routes(v1, db)
}

func setupSwagger(r *gin.Engine) {
	docs.SetupSwagger(r)
}
