package model

import (
	"time"

	"github.com/JesusIslam/lowger"
	"github.com/JesusIslam/sikritklab/internal/constant"
	"github.com/JesusIslam/sikritklab/internal/database"
)

func init() {
	err := database.DB.Init(&Thread{})
	if err != nil {
		lowger.Fatal(constant.ErrorFailedToInitializeDatabaseThread, err)
	}
}

type Thread struct {
	ID        string    `json:"id" storm:"id"`
	CreatedAt time.Time `json:"created_at" storm:"index"`
	Title     string    `json:"title"`
}
