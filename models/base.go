package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
   }
   // BeforeCreate will set a UUID rather than numeric ID.
   func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
	 return err
	}
	base.ID = uuid
	return err
   }