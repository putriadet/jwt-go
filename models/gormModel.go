package models

import "time"

type GormModel struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAT *time.Time `json:"updated_at,omitempty"`
}
