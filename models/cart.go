package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID          string `gorm:"primaryKey" json:"id"`
	ItemName    string `json:"itemName"`
	Price uint `json:"price"`
	Description string `json:"description"`
	ItemLocalID string `json:"itemLocalID"`
	Quantity    uint   `json:"quantity"` // Quantity is typically a number, not a string
	UniqueID 	string `json:"uniqueid"`
	UserID uint
	User   User
}