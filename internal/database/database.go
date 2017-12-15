package database

import (
	"crypto/rand"
	"os"
	"time"

	"github.com/JesusIslam/lowger"
	"github.com/JesusIslam/sikritklab/internal/constant"
	"github.com/asdine/storm"
	bolt "github.com/coreos/bbolt"
	"github.com/oklog/ulid"
)

var (
	DB  *storm.DB
	err error
)

func init() {
	dbPath := os.Getenv(constant.EnvDatabasePath)
	if dbPath == "" {
		dbPath = constant.DefaultDBPath
	}
	DB, err = storm.Open(dbPath, storm.BoltOptions(0644, &bolt.Options{
		Timeout: time.Second,
	}))
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
