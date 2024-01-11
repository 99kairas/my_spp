package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	ID                          uuid.UUID `json:"id" form:"id"`
	UserID                      uuid.UUID `gorm:"index" json:"user_id" form:"user_id"`
	TeacherIdentificationNumber string    `gorm:"type:varchar(8)" json:"teacher_identification_number" form:"teacher_identification_number"`
	Name                        string    `json:"name" form:"name"`
	Address                     string    `json:"address" form:"address"`
	Gender                      string    `gorm:"type:enum('male', 'female')" json:"gender" form:"gender"`
	TelephoneNumber             string    `json:"telephone_number" form:"telephone_number"`
	ProfileImage                string    `json:"profile_image" form:"profile_image"`
	LessonContent               string    `json:"lesson_content" form:"lesson_content"`
}
