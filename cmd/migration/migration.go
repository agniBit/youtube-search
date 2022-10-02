package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/agniBit/youtube-search/pkg/youtube"
	"github.com/agniBit/youtube-search/utl/common"
)

func main() {
	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		panic("DATABASE_URL is not set")
	}

	sqlDB, err := sql.Open("postgres", database_url)
	checkErr(err)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		PrepareStmt: true,
	})
	checkErr(err)

	if err := db.Exec("SELECT 1").Error; err != nil {
		checkErr(err)
	}

	runMigration(db)
}

func runMigration(db *gorm.DB) {
	utils.RunSQLFromFile(db, "cmd/migration/sql/setup.sql")

	tx := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	if err := tx.Error; err != nil {
		checkErr(err)
	}

	err := db.AutoMigrate(&youtube.YoutubeVideo{})
	if err != nil {
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
