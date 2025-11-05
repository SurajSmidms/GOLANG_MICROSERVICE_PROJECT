package models

import "time"

type Payment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderID   uint      `json:"order_id"`
	UserID    uint      `json:"user_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"` // e.g., "SUCCESS", "FAILED", "PENDING"
	CreatedAt time.Time `json:"created_at"`
}
