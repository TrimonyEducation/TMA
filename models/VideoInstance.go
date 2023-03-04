package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type VideoInstance struct {
	Base
	Thumbnail string `gorm:"not null;default:null"`
	Link string  `gorm:"not null;default:null"`
	Description string `gorm:"not null;default:null"`
	Duration string  `gorm:"not null;default:null"`
	TopicsTags pq.StringArray `gorm:"type:text[]"`
	SubjectsTags pq.StringArray `gorm:"type:text[]"`
	UserID uuid.UUID `gorm:"not null;default:null"`
	PlaylistID uuid.UUID
	ReviewID uuid.UUID
}