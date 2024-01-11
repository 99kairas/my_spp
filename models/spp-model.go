package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Spp struct {
	gorm.Model
	ID                uuid.UUID `json:"id" form:"id"`
	ClassID           uuid.UUID `gorm:"index" json:"class_id" form:"class_id"`
	StartAcademicYear time.Time `json:"start_academic_year" form:"start_academic_year"`
	EndAcademicYear   time.Time `json:"end_academic_year" form:"end_academic_year"`
	Amount            float64   `json:"amount" form:"amount"`
	Description       string    `gorm:"type:text;" json:"description" form:"description"`
}
