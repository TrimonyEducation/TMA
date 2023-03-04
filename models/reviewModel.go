package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Review struct {
	Base
	Count int  `gorm:"not null;default:0"`
	SubjectTags pq.StringArray `gorm:"type:text[]"`
	Videos []VideoInstance
	UserID uuid.UUID `gorm:"not null;default:null"`
}