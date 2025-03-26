package models

import "time"

// Inventory represents a storage location or collection
type Inventory struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"` // e.g., "Main Warehouse"
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Item represents a product type stored in an inventory
type Item struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	InventoryID   int       `json:"inventory_id" gorm:"not null"` // Foreign key to Inventory
	Name          string    `json:"name" gorm:"type:varchar(100);not null"`
	Description   string    `json:"description" gorm:"type:text"`
	Quantity      int       `json:"quantity" gorm:"not null"`
	LastRestockAt time.Time `json:"last_restock_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// RestockHistory tracks restocking events for an item
type RestockHistory struct {
	ID               int       `json:"id" gorm:"primaryKey"`
	ItemID           int       `json:"item_id" gorm:"not null"` // Foreign key to Item
	RestockAmount    int       `json:"restock_amount" gorm:"not null"`
	RestockTimestamp time.Time `json:"restock_timestamp"`
}
