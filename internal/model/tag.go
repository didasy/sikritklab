package model

import (
	"time"

	"github.com/JesusIslam/lowger"
	"github.com/JesusIslam/sikritklab/internal/constant"
	"github.com/JesusIslam/sikritklab/internal/database"
)

func init() {
	err := database.DB.Init(&Tag{})
	if err != nil {
		lowger.Fatal(constant.ErrorFailedToInitializeDatabaseTag, err)
	}
}

type Tag struct {
	ID        int       `json:"id" storm:"id,increment"`
	CreatedAt time.Time `json:"created_at" storm:"index"`
	ThreadID  string    `json:"thread_id" storm:"index"`
	Tag       string    `json:"tag" storm:"index"`
}
