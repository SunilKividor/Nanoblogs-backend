package models

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	Id         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"user_id"`
	Title      string    `json:"title"`
	Content    string    `jsong:"content"`
	Category   string    `json:"category"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type GetBlogResBody struct {
	Id         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Content    string    `jsong:"content"`
	Category   string    `json:"category"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type PostBlogReqModel struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

type UpdateBlogReqModel struct {
	BlogId   uuid.UUID `json:"blog_id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Category string    `json:"category"`
}

type DeleteBlogReqModel struct {
	BlogId uuid.UUID `json:"blog_id"`
}
