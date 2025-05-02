
package database

import (
    "expensetracker/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func ConnectDatabase() {
    db, err := gorm.Open(sqlite.Open("expenses.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    db.AutoMigrate(&models.User{}, &models.Expense{})
    DB = db
}
