package db

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
	"learning-golang/constants"
)

func RunMigrations(db *gorm.DB) error {
	// gorm.DBからdatabase/sql.DBへの変換
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("%s: %v", constants.ErrGetDBConnection, err)
	}

	// MySQLドライバーのインスタンスを作成
	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("%s: %v", constants.ErrCreateMySQLDriver, err)
	}

	// マイグレーションのインスタンスを作成
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		return fmt.Errorf("%s: %v", constants.ErrConnectMySQL, err)
	}

	// マイグレーションを実行
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("%s: %w", constants.ErrMigrationFailed, err)
		}
	}

	return nil
}
