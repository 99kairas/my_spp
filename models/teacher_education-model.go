package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeacherEducation struct {
	gorm.Model
	ID         uuid.UUID `json:"id" form:"id"`
	TeacherID  uuid.UUID `json:"teacher_id" form:"teacher_id"`
	StartDate  time.Time `json:"start_date" form:"start_date"`
	EndDate    time.Time `json:"end_date" form:"end_date"`
	Education  string    `json:"education" form:"education"`
	StudyPlace string    `json:"study_place" form:"study_place"`
}
