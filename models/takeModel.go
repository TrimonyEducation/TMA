package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Take struct {
	Base
	Progress int `gorm:"not null;default:0"`
	Answers pq.GenericArray `gorm:"type:json[]"`
	UserID uuid.UUID `gorm:"not null;default:null"`
	ExerciseID uuid.UUID `gorm:"not null;default:null"`
		
}