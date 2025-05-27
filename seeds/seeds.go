package seed

import (
	"expensetracker/config"
	"expensetracker/database"
	"expensetracker/models"
	"expensetracker/utils"
)

// SeedAdmin crée un utilisateur admin si aucun n’existe
func SeedAdmin() {
	email := config.GetEnv("ADMIN_EMAIL", "admin@example.com")
	pass := config.GetEnv("ADMIN_PASS", "admin")

	var count int64
	database.DB.Model(&models.User{}).
		Where("email = ?", email).
		Count(&count)

	if count == 0 {
		hash, _ := utils.HashPassword(pass)
		admin := models.User{
			Name:         "Admin",
			Email:        email,
			PasswordHash: hash,
			IsAdmin:      true,
		}
		database.DB.Create(&admin)
	}
}
