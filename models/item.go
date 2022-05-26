package models

import (
	"time"
)

type Item struct {
	ID          uint       `gorm:"primaryKey" json:"item_id"`
	ItemCode    string     `json:"item_code"`
	Description string     `json:"description"`
	OrderId     uint       `json:"order_id"`
	Quantity    int        `json:"quantity"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"-"`
}
