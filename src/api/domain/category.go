package domain

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        uint         `json:"id" gorm:"column:id;primarykey"`
	Name      string       `json:"name" gorm:"column:name" binding:"required"`
	Photo     string       `json:"photo" gorm:"column:photo" binding:"required"`
	Products  []Product    `json:"products,omitempty"`
	CreatedAt time.Time    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at,omitempty" gorm:"column:deleted_at;index"`
}

func (c *Category) TableName() string {
	return "category"
}
