
package models

import "time"

type User struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    Name         string    `gorm:"size:120" json:"name" binding:"required"`
    Email        string    `gorm:"size:120;uniqueIndex" json:"email" binding:"required,email"`
    PasswordHash string    `json:"-"`
    IsAdmin      bool      `json:"is_admin"`
    CreatedAt    time.Time `json:"-"`
    UpdatedAt    time.Time `json:"-"`
}
