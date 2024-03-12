package models

import "time"

type Role struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"email" gorm:"not null;unique;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
