package models

import (
	"time"
)

type User struct {
	ID                  string `gorm:"primary_key;"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time `sql:"index"`
	Email               string     `gorm:"unique;not null;default:null; validate:email"`
	Name                string     `gorm:"unique;not null;default:null"`
	ProfilePicture      string
	Role                string `gorm:"not null;default:'user'"`
	SchoolGrade         string 
	SchoolLevel         string 
	IsPaid              bool   `gorm:"not null;default:false"`
	IsAdmin             bool   `gorm:"not null;default:false"`
	IsBanned            bool   `gorm:"not null;default:false"`
	CompletedOnboarding bool   `gorm:"not null;default:false"`
	IsTeacher           bool   `gorm:"not null;default:false"`
	EmailVerified       bool   `gorm:"not null;default:false"`
	Playlists           []Playlist
	Review              Review
	VideoInstance       []VideoInstance
	Classes             []Class `gorm:"many2many:user_classes;"`
	Teacher             Teacher
	Take                []Take
}
