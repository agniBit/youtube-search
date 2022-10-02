package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/agniBit/youtube-search/pkg/youtube"
	utils "github.com/agniBit/youtube-search/utl/common"
)

func main() {
	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		panic("DATABASE_URL is not set")
	}

	sqlDB, err := sql.Open("postgres", database_url)
	utils.CheckErr(err)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	utils.CheckErr(err)

	if err := db.Exec("SELECT 1").Error; err != nil {
		utils.CheckErr(err)
	}

	runMigration(db)
}

func runMigration(db *gorm.DB) {
	utils.RunSQLFromFile(db, "cmd/migration/sql/setup.sql")

	tx := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	if err := tx.Error; err != nil {
		utils.CheckErr(err)
	}

	err := db.Exec("CREATE SCHEMA IF NOT EXISTS youtube").Error
	utils.CheckErr(err)

	err = db.AutoMigrate(
		&youtube.YoutubeVideo{},
		&youtube.YoutubeVideoThumbnails{},
	)
	if err != nil {
		utils.CheckErr(err)
	}
}
