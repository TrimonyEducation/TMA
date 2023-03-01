package models

import (
	"gorm.io/gorm"
	"github.com/lib/pq"
)

type Class struct {
	gorm.Model
	Class_Title string
	Class_Description string
	Class_SchoolGrade pq.StringArray `gorm:"type:text[]"`
	Class_SchoolLevel string //havo, vmbo, havo-vwo vmbo-havo vwo
	Class_NumOfStudents string
	Class_Subjects pq.StringArray `gorm:"type:text[]"`
	Teacher []Teacher `gorm:"many2many:class_teacher;"`
}