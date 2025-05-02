package controllers

import (
	"expensetracker/database"
	"expensetracker/models"
	"expensetracker/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User credentials"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /users [post]
func Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "invalid_input", err.Error())
		return
	}

	hash, _ := utils.HashPassword(input.Password)
	user := models.User{Name: input.Name, Email: input.Email, PasswordHash: hash}
	if err := database.DB.Create(&user).Error; err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "email_exists", err.Error())
		return
	}
	token, _ := utils.GenerateToken(user.ID, false)
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

// @Summary Login
// @Tags auth
// @Accept json
// @Produce json
// @Param creds body map[string]string true "Credentials"
// @Success 200 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "invalid_input", err.Error())
		return
	}
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		utils.ErrorJSON(c, http.StatusUnauthorized, "invalid_credentials", "wrong email or password")
		return
	}
	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		utils.ErrorJSON(c, http.StatusUnauthorized, "invalid_credentials", "wrong email or password")
		return
	}
	token, _ := utils.GenerateToken(user.ID, user.IsAdmin)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
