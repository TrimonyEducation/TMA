package models

import "github.com/google/uuid"

type Answers struct {
	Base
	Answer string `gorm:"not null; default:null; validate:string(255)"` 
	IsCorrect bool `gorm:"default:false"`
	ProblemID uuid.UUID `gorm:"not null; default:null"`
	TakeID uuid.UUID `gorm:"not null; default:null"`
}