package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learning-golang/constants"
	"os"
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
