package models

import (
	"github.com/lib/pq"
)

type Video struct {
	Base
	Title  string `gorm:"unique;not null;default:null"`
    Description string `gorm:"unique;not null;default:null"`
	Duration string `gorm:"not null;default:null"`
	FileFormat string
	SizeInMB string
	Resolution string
	AspectRatio string
	SubjectTags pq.StringArray `gorm:"type:text[];not null;default:null"`
	TopicTags string
	Chapters pq.StringArray `gorm:"type:text[];not null;default:null"`
	Url string `gorm:"not null;default:null"`
	ThumbnailUrl string `gorm:"not null;default:null"`
	Views int `gorm:"default:0;"`
	Schoolgrade pq.StringArray `gorm:"type:text[];not null;default:null"`
	Exercises []Exercise `gorm:"many2many:video_exercise;"`
}