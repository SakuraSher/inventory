package inventory

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	UnityPrice  float64   `json:"unit_price"`
	LastOrdered time.Time `json:"last_ordered"`
}

type InventoryChanged struct {
	ItemID    uint      `json:"item_id"`
	EventType string    `json:"event_type"`
	Quantity  int       `json:"quantity"`
	Timestamp time.Time `json:"time_stamp"`
}
