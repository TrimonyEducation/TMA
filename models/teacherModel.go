package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Teacher_School string
	Teacher_Subjects pq.StringArray `gorm:"type:text[]"`
	UserID int
}