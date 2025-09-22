package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // GORM will handle ID (uint), CreatedAt, UpdatedAt, and DeletedAt

	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"` // Security best practice: don't serialize the password

	// Associations: GORM will use the foreign keys on the other structs
	Products []Product
	Carts    []Cart
}
