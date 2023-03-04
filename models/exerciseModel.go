package models

import (
	"github.com/lib/pq"
)

type Exercise struct {
	Base
	Title string `gorm:"not null;default:null"`
	Description string `gorm:"not null;default:null"`
	NumOfProblems int 
	DifficultyLevel string `gorm:"not null;default:null"`
	SubjectTags pq.StringArray `gorm:"type:text[];not null;default:null"`
	Chapters pq.StringArray `gorm:"type:text[];not null;default:null"`
	TopicsTags pq.StringArray `gorm:"type:text[]"`
	Takes []Take
	Problems []Problem
	Videos []Video `gorm:"many2many:video_exercise;"`
}