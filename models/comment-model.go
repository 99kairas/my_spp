package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID              uuid.UUID `json:"id" form:"id"`
	NewsID          uuid.UUID `gorm:"column:news_id;index" json:"news_id" form:"news_id"`
	UserID          uuid.UUID `gorm:"column:user_id;index" json:"user_id" form:"user_id"`
	ParentCommentID uuid.UUID `gorm:"index;null" json:"parent_comment_id" form:"parent_comment_id"`
	Comment         string    `gorm:"type:text;" json:"comment" form:"comment"`
	Date            time.Time `json:"date" form:"date"`
}
