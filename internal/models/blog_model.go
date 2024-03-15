package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Blog struct {
	Id         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"user_id" validate:"required"`
	Title      string    `json:"title" validate:"required"`
	Content    string    `jsong:"content" validate:"required"`
	Category   string    `json:"category" validate:"required"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type GetBlogResBody struct {
	Id         uuid.UUID      `json:"id"`
	Title      string         `json:"title"`
	Content    string         `jsong:"content"`
	Category   pq.StringArray `json:"category"`
	Created_At time.Time      `json:"created_at"`
	Updated_At time.Time      `json:"updated_at"`
}

type PostBlogReqModel struct {
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Category string `json:"category" validate:"required"`
}

type UpdateBlogReqModel struct {
	BlogId   uuid.UUID      `json:"blog_id" validate:"required"`
	Title    string         `json:"title" validate:"required"`
	Content  string         `json:"content" validate:"required"`
	Category pq.StringArray `json:"category" validate:"required"`
}

type DeleteBlogReqModel struct {
	BlogId uuid.UUID `json:"blog_id" validate:"required"`
}
