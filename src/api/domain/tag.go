package domain

import (
	"database/sql"
	"time"
)

type Tag struct {
	ID        uint         `json:"id" gorm:"column:id;primarykey"`
	Name      string       `json:"name" gorm:"column:name" binding:"required"`
	Products  []*Product   `gorm:"many2many:product_tags;"`
	CreatedAt time.Time    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at,omitempty" gorm:"column:deleted_at;index"`
}

func (T *Tag) TableName() string {
	return "tag"
}
