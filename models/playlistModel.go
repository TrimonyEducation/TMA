package models

import "github.com/google/uuid"

type Playlist struct {
	Base
	Count int `gorm:"not null;default:0"`
	Title string `gorm:"not null;default:null"`
	Videos []VideoInstance
	UserID uuid.UUID `gorm:"not null;default:null"`
}