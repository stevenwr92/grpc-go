package models

import "time"

type RolesRight struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	RoleId    uint `gorm`
	Route     string
	Section   string
	Path      string
	RCreate   uint
	RRead     uint
	RUpdate   uint
	RDelete   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
