package models

import (
	"github.com/lib/pq"
)

type Class struct {
	Base
	Title string `gorm:"not null;default:null"`
	Description string
	SchoolGrade pq.StringArray `gorm:"type:text[]"`
	SchoolLevel pq.StringArray `gorm:"type:text[]"`
	NumOfStudents string
	Subjects pq.StringArray `gorm:"type:text[]"`
	Teacher []Teacher `gorm:"many2many:class_teacher;"`
}