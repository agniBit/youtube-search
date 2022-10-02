package postgres

import (
	"database/sql"

	"github.com/agniBit/youtube-search/type/common"
	"github.com/agniBit/youtube-search/utl/config"
	_ "github.com/lib/pq"                                 // DB adapter
	_ "github.com/newrelic/go-agent/v3/integrations/nrpq" // nrpostgres DB driver
	gpsql "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const sqlDriverName = "nrpostgres"

func NewGormDB(dbConfig *config.Database) (*gorm.DB, error) {
	sqlDB, err := sql.Open(sqlDriverName, dbConfig.URLPath)
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	if err != nil {
		return &gorm.DB{}, err
	}

	gdb, err := gorm.Open(gpsql.New(gpsql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})

	if dbConfig.LogQueries {
		gdb.Logger = logger.Default.LogMode(logger.Info)
	}

	return gdb, err
}

func Pagination(q *gorm.DB, offsetLimit *common.OffsetLimit) {
	if offsetLimit.Limit == 0 {
		offsetLimit.Limit = 10
	}

	q.Limit(offsetLimit.Limit).Offset(offsetLimit.Offset)
}
