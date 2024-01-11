package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID            uuid.UUID `json:"id" form:"id"`
	UserID        uuid.UUID `json:"user_id" form:"user_id"`
	SppID         uuid.UUID `json:"spp_id" form:"spp_id"`
	Date          time.Time `json:"date" form:"date"`
	PaymentMethod string    `gorm:"type:enum('transfer', 'manual');" json:"payment_method" form:"payment_method"`
	TotalAmount   float64   `json:"total_amount" form:"total_amount"`
	Status        string    `gorm:"type:enum('pending', 'partial', 'success', 'failed');default:pending" json:"status" form:"status"`
}
