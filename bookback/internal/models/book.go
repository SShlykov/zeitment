package models

import (
	"time"
)

type Book struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Owner       int       `json:"owner"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"is_public"`
	Publication NullTime  `json:"publication"`
	DeletedAt   NullTime  `json:"deleted_at"`
	//Chapters    []Chapter
}
