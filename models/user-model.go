package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"id" form:"id"`
	Email    string    `json:"email" form:"email"`
	Password string    `json:"password" form:"password"`
	OTP      string    `json:"otp" form:"otp"`
	Status   string    `gorm:"type:enum('verified', 'unverified');default:unverified" json:"status" form:"status"`
}
