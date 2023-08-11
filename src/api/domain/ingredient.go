package domain

import (
	"database/sql"
	"time"
)

type Ingredient struct {
	ID          uint         `json:"id" gorm:"column:id;primarykey"`
	Name        string       `json:"name" gorm:"column:name" binding:"required"`
	Description string       `json:"description" gorm:"column:description"`
	Tags        []*Tag       `gorm:"many2many:ingredient_tags;"`
	CreatedAt   time.Time    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at,omitempty" gorm:"column:deleted_at;index"`
}

func (i *Ingredient) TableName() string {
	return "ingredient"
}
