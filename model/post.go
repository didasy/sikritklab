package model

import "time"

type Post struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	ThreadID  int64      `json:"thread_id"`
	Title     string     `json:"title,omitempty"`
	Content   string     `json:"content"`
}
