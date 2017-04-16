package model

import (
	"time"

	"github.com/JesusIslam/lowger"
	"github.com/JesusIslam/sikritklab/database"
)

func init() {
	err := database.DB.Init(&Thread{})
	if err != nil {
		lowger.Fatal("Failed to initialize database thread:", err)
	}
}

type Thread struct {
	ID        string    `json:"id" storm:"id"`
	CreatedAt time.Time `json:"created_at" storm:"index"`
	Title     string    `json:"title"`
}
