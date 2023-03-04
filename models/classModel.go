package models

import (
	"github.com/lib/pq"
)

type Class struct {
	Base
	Title string `gorm:"not null;default:null"`
	Description string
	JoinCode string `gorm:"not null;default:null;unique"`
	SchoolGrade pq.StringArray `gorm:"type:text[]"`
	SchoolLevel pq.StringArray `gorm:"type:text[]"`
	NumOfStudents string
	Subjects pq.StringArray `gorm:"type:text[]"`
	Students []User `gorm:"many2many:user_classes;"`
	Teacher []Teacher `gorm:"many2many:class_teacher;"`
}