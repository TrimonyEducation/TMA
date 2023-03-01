package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Video_Title  string `gorm:"unique;not null;default:null"`
    Video_Description string `gorm:"unique;not null;default:null"`
	Video_Duration string `gorm:"not null;default:null"`
	Video_FileFormat string
	Video_SizeInMB string
	Video_Resolution string
	Video_AspectRatio string
	Video_SubjectTags pq.StringArray `gorm:"type:text[];not null;default:null"`
	Video_TopicTags string
	Video_Chapters pq.StringArray `gorm:"type:text[];not null;default:null"`
	Video_Url string `gorm:"not null;default:null"`
	Video_ThumbnailUrl string `gorm:"not null;default:null"`
	Video_views int `gorm:"default:0;"`
	Video_Schoolgrade pq.StringArray `gorm:"type:text[];not null;default:null"`
	Exercises []Exercise `gorm:"many2many:video_exercise;"`
}