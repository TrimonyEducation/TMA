package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Review_Count int  `gorm:"not null;default:0"`
	Review_SubjectTags pq.StringArray `gorm:"type:text[];not null;default:null"`
	VideoInstance []VideoInstance
	UserID uint `gorm:"not null;default:null"`
}