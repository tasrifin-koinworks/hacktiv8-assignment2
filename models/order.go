package models

import (
	"time"
)

type Order struct {
	ID           int        `gorm:"primaryKey" json:"order_id"`
	CustomerName string     `gorm:"not null;type:varchar(100)" json:"customer_name"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"-"`
	Items        []Item     `json:"items"`
}
