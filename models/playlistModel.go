package models

import (
	"gorm.io/gorm"
)

type Playlist struct {
	gorm.Model
	Playlist_Count int `gorm:"not null;default:0"`
	Playlist_Name string `gorm:"not null;default:null"`
	VideoInstance []VideoInstance
	UserID uint `gorm:"not null;default:null"`
}