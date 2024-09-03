package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learning-golang/db"
	"log"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

var dummyDB *gorm.DB // 仮のgorm.DBインスタンス

func init() {
	var err error

	// Load .env.test variables
	err = godotenv.Load(filepath.Join("..", ".env.test"))
	DB_USER := os.Getenv("DB_USER")
	log.Println("DB_USER:", DB_USER)

	if err != nil {
		log.Fatalf("Error loading .env.test file")
	}

	dataSourceName := db.GetDataSourceName()

	dummyDB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to initialize test database. ", err)
	}
}

func setup() {
	gin.SetMode(gin.TestMode)
}

func TestGetPosts(t *testing.T) {
	setup()
	// Ginのテスト用HTTPレスポンスライタとコンテキストを作成
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	GetPosts(c, dummyDB)

	// 応答ステータスコードと応答ボディを確認
	assert.Equal(t, 200, w.Code)
	// 応答ボディを検証（具体的な内容は実際のサービスロジックに基づく）
}

func TestGetPost(t *testing.T) {
	setup()
	// Ginのテスト用HTTPレスポンスライタとコンテキストを作成
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// パラメータとしてポストIDを追加します（存在するIDを設定）
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	GetPost(c, dummyDB)

	// 応答ステータスコードと応答ボディを確認
	assert.Equal(t, 200, w.Code)
	// 応答ボディを検証（具体的な内容は実際のサービスロジックに基づく）
}
