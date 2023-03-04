package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Teacher struct {
	Base
	School string
	Subjects pq.StringArray `gorm:"type:text[]"`
	UserID uuid.UUID
}