package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Problem_QST string
	Problem_Body string
	Problem_Type string
	Problem_Options pq.GenericArray `gorm:"type:json[]"`
	Problem_Difficulty string
	Problem_Answer string
	Problem_AssetLinks pq.GenericArray `gorm:"type:json[]"`
	Problem_SolutionText string
	Problem_SolutionVideo string
	ExerciseID uint `gorm:"not null;default:null"`
}