package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learning-golang/constants"
	"os"
	"time"
)

func InitDB() (*gorm.DB, error) {
	// .env ファイルをロード
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("%s: %v", constants.ErrLoadEnvFile, err)
	}

	// 接続文字列を組み立てる
	dataSourceName := GetDataSourceName()

	// MySQLへの接続
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %v", constants.ErrConnectMySQL, err)
	}

	// データベース接続のPingを行う
	if err := db.Exec("SELECT 1").Error; err != nil {
		return nil, fmt.Errorf("%s: %v", constants.ErrPingDatabase, err)
	}

	// モデルの一括マイグレーションを実行
	if err := RunMigrations(db); err != nil {
		return nil, fmt.Errorf("%s: %v", constants.ErrMigration, err)
	}

	return db, nil
}

func GetDataSourceName() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	return dataSourceName
}

func GetDataSourceName2() string {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		// エラーハンドリング
	}
	c := gorm.Config{
		DBName:    "db",
		User:      "user",
		Passwd:    "password",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	return c
}
