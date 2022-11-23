package model

import (
	"time"
)

type Item struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Status    int       `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
