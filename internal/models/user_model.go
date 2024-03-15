package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	Id           uuid.UUID      `json:"id"`
	Name         string         `json:"name"`
	Username     string         `json:"username"`
	Password     string         `json:"password"`
	RefreshToken string         `json:"refresh_token"`
	Category     pq.StringArray `json:"category"`
}
