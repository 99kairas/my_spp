package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID                          uuid.UUID `json:"id" form:"id"`
	UserID                      uuid.UUID `gorm:"index" json:"user_id" form:"user_id"`
	StudentIdentificationNumber string    `gorm:"type:varchar(8)" json:"student_identification_number" form:"student_identification_number"`
	Name                        string    `json:"name" form:"name"`
	Address                     string    `json:"address" form:"address"`
	Gender                      string    `gorm:"type:enum('male', 'female')" json:"gender" form:"gender"`
	DateOfBirth                 time.Time `json:"date_of_birth" form:"date_of_birth"`
	ProfileImage                string    `json:"profile_image" form:"profile_image"`
	ClaimCode                   string    `json:"claim_code" form:"claim_code"`
	ClassID                     uuid.UUID `gorm:"index" json:"class_id" form:"class_id"`
}
