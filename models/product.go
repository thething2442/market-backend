package models
import (
	"time"
    "gorm.io/gorm"
	
)
type Product struct {
    ID        string `gorm:"primaryKey"`
    Email     string `gorm:"unique" json:"email"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
    UserID uint
	User   User
}