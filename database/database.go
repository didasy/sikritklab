package database

import (
	"crypto/rand"
	"os"
	"time"

	"github.com/JesusIslam/lowger"
	"github.com/asdine/storm"
	"github.com/oklog/ulid"
)

const (
	DBPathEnv = "SIKRITKLAB_DATABASE_PATH"
)

var (
	DB  *storm.DB
	err error
)

func init() {
	dbPath := os.Getenv(DBPathEnv)
	if dbPath == "" {
		dbPath = "./sikritklab.db"
	}
	DB, err = storm.Open(dbPath)
	if err != nil {
		lowger.Fatal("Failed to open database file:", err)
	}
}

func NewThreadID() (string, error) {
	t := ulid.Timestamp(time.Now())
	id, err := ulid.New(t, rand.Reader)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
