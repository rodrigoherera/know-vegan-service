package domain

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          uint         `json:"id" gorm:"column:id;primarykey"`
	Name        string       `json:"name" gorm:"column:name" binding:"required"`
	Description string       `json:"description" gorm:"column:description"`
	Ingredients string       `json:"ingredients" gorm:"column:ingredients"`
	Photo       string       `json:"photo" gorm:"column:photo"`
	CategoryID  uint         `json:"category_id" gorm:"column:category_id" binding:"required"`
	Tags        []*Tag       `gorm:"many2many:product_tags;"`
	CreatedAt   time.Time    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at,omitempty" gorm:"column:deleted_at;index"`
}

func (p *Product) TableName() string {
	return "product"
}
