package database

import (
	"os"
	"strconv"

	"github.com/JesusIslam/lowger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB                 *gorm.DB
	PostLimitPerThread int
	NotFound           = gorm.ErrRecordNotFound
	ThreadDisplayOrder = "created_at DESC"
)

func init() {
	var err error

	PostLimitPerThread, err = strconv.Atoi(os.Getenv("POST_LIMIT_PER_THREAD"))
	if err != nil {
		lowger.Fatal("POST_LIMIT_PER_THREAD environment variable is not a valid integer:", err)
	}

	sqlConnString := os.Getenv("SQL_CONNECTION_STRING")
	dialect := os.Getenv("SQL_DIALECT")

	DB, err = gorm.Open(dialect, sqlConnString)
	if err != nil {
		lowger.Fatal("Failed to establish database connection:", err)
	}
}

func New() *gorm.DB {
	return DB.New()
}
