
package models

import "time"

type Expense struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `json:"user_id"`
    Label     string    `gorm:"size:120" json:"label" binding:"required"`
    Amount    float64   `json:"amount" binding:"required"`
    Category  string    `gorm:"size:60" json:"category" binding:"required"`
    Date      time.Time `json:"date" binding:"required"`
    Notes     string    `gorm:"size:500" json:"notes"`
    CreatedAt time.Time `json:"-"`
    UpdatedAt time.Time `json:"-"`
}
