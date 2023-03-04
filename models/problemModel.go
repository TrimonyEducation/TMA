package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Problem struct {
	Base 
	QST string `gorm:"not null;default:null"`
	Body string `gorm:"not null;default:null"`
	Type string `gorm:"not null;default:null"`
	Options pq.GenericArray `gorm:"type:json[]"`
	Difficulty string `gorm:"not null;default:null"`
	Answer string `gorm:"not null;default:null"`
	AssetLinks pq.GenericArray `gorm:"type:json[]"`
	SolutionText string 
	SolutionVideo string
	IsPublished bool `gorm:"not null;default:false"`
	IsPublic bool `gorm:"not null;default:false"`
	ExerciseID uuid.UUID `gorm:"not null;default:null"`
}