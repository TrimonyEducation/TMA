package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	User_Email string `gorm:"unique;not null;default:null"`
	User_Password string
	User_Name string `gorm:"unique;not null;default:null"`
	User_ProfilePicture string
	User_Role string `gorm:"not null;default:'user'"`
	User_SchoolGrade string `gorm:"not null;default:null"`
	User_SchoolLevel string `gorm:"not null;default:null"`
	User_IsPaid bool `gorm:"not null;default:false"`
	User_IsAdmin bool `gorm:"not null;default:false"`
	User_IsBanned bool `gorm:"not null;default:false"`
	User_CompletedOnboarding bool `gorm:"not null;default:false"`
	User_isTeacher bool `gorm:"not null;default:false"`
	User_EmailVerified bool `gorm:"not null;default:false"`
	Playlists []Playlist
	Review Review
	VideoInstance []VideoInstance
	Classes []Class `gorm:"many2many:user_classes;"`
	Teacher Teacher
	Take []Take
}