package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Take struct {
	gorm.Model
	Take_Progress int
	Take_Answers pq.GenericArray `gorm:"type:json[]"`
	UserID uint `gorm:"not null;default:null"`
	ExerciseID uint `gorm:"not null;default:null"`
		
}