package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ManualPayment struct {
	gorm.Model
	ID            uuid.UUID `json:"id" form:"id"`
	TransactionID uuid.UUID `gorm:"index" json:"transaction_id" form:"transaction_id"`
	AmountPaid    float64   `json:"amount_paid" form:"amount_paid"`
	Date          time.Time `json:"date" form:"date"`
	PaymentImage  string    `json:"payment_image" form:"payment_image"`
	Status        string    `gorm:"type:enum('incomplete', 'pay', 'complete');default:incomplete" json:"status" form:"status"`
}
