package models

import "github.com/lib/pq"

type Chapter struct {
	Base
	Title string 
	Videos []Video
	Subject string `gorm:"not null;default:null"`
	TopicTags pq.StringArray `gorm:"type:text[];not null;default:null"`
}