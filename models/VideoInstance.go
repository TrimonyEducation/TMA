package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type VideoInstance struct {
	gorm.Model
	VideoInstance_Thumbnail string
	VideoInstance_Link string
	VideoInstance_Description string
	VideoInstance_Duration string
	VideoInstance_TopicsTags pq.StringArray `gorm:"type:text[]"`
	VideoInstance_SubjectsTags pq.StringArray `gorm:"type:text[]"`
	UserID uint `gorm:"not null;default:null"`
	PlaylistID uint
	ReviewID uint
}