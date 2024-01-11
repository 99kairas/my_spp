package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	ID      uuid.UUID `json:"id" form:"id"`
	UserID  uuid.UUID `gorm:"index" json:"user_id" form:"user_id"`
	Title   string    `json:"title" form:"title"`
	Content string    `gorm:"type:text" json:"content" form:"content"`
	Image   string    `json:"image" form:"image"`
	View    int       `gorm:"default:0" json:"view" form:"view"`
	Date    time.Time `json:"date" form:"date"`
	Status  string    `gorm:"type:enum('publish', 'draft')" json:"status" form:"status"`
}
