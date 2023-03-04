package models

import (
	"github.com/google/uuid"
)

type Take struct {
	Base
	Progress int `gorm:"not null;default:0"`
	Answers []Answers
	IsFinished bool `gorm:"not null;default:false"`
	UserID uuid.UUID `gorm:"not null;default:null"`
	ExerciseID uuid.UUID `gorm:"not null;default:null"`	
}