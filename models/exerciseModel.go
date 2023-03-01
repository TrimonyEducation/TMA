package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	Exercise_Title string `gorm:"not null;default:null"`
	Exercise_Description string `gorm:"not null;default:null"`
	Exercise_NumOfProblems int 
	Exercise_DifficultyLevel string `gorm:"not null;default:null"`
	Exercise_SubjectTags pq.StringArray `gorm:"type:text[]"`
	Exercise_Chapters pq.StringArray `gorm:"type:text[]"`
	Exercise_TopicsTags pq.StringArray `gorm:"type:text[]"`
	Takes []Take
	Problems []Problem
	Videos []Video `gorm:"many2many:video_exercise;"`
}