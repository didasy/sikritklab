package model

import (
	"time"
)

type Thread struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Title     string     `json:"title"`
	Tags      string     `json:"tags"`
}
