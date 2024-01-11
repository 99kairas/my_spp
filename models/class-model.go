package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ID                uuid.UUID `json:"id" form:"id"`
	ClassName         string    `json:"class_name" form:"class_name"`
	LevelClass        string    `json:"level_class" form:"level_class"`
	StartAcademicYear time.Time `json:"start_academic_year" form:"start_academic_year"`
	EndAcademicYear   time.Time `json:"end_academic_year" form:"end_academic_year"`
	TeacherID         uuid.UUID `gorm:"index" json:"teacher_id" form:"teacher_id"`
}
