package controllers

import (
	"expensetracker/database"
	"expensetracker/models"
	"expensetracker/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Update own profile
// @Tags users
// @Security BearerAuth
// @Accept json
// @Param profile body object true "Profile fields" example({"name":"Alice","email":"alice@ex.com","password":"newpass"})
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /users/me [put]
func UpdateMe(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input struct {
		Name     string `json:"name" binding:"omitempty"`
		Email    string `json:"email" binding:"omitempty,email"`
		Password string `json:"password" binding:"omitempty,min=6"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "invalid_input", err.Error())
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		utils.ErrorJSON(c, http.StatusNotFound, "not_found", "user not found")
		return
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Password != "" {
		hash, _ := utils.HashPassword(input.Password)
		user.PasswordHash = hash
	}

	if err := database.DB.Save(&user).Error; err != nil {
		utils.ErrorJSON(c, http.StatusInternalServerError, "db_error", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"name":     user.Name,
		"email":    user.Email,
		"is_admin": user.IsAdmin,
	})
}

// @Summary Admin update user (non-admin)
// @Tags admin
// @Security BearerAuth
// @Accept json
// @Param id path int true "User ID"
// @Param data body object true "Fields to update" example({"name":"Bob","email":"bob@ex.com","is_admin":false})
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /admin/users/{id} [put]
func AdminUpdateUser(c *gin.Context) {
	targetID := c.Param("id")

	var input struct {
		Name    string `json:"name" binding:"omitempty"`
		Email   string `json:"email" binding:"omitempty,email"`
		IsAdmin *bool  `json:"is_admin" binding:"omitempty"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorJSON(c, http.StatusBadRequest, "invalid_input", err.Error())
		return
	}

	var user models.User
	if err := database.DB.First(&user, targetID).Error; err != nil {
		utils.ErrorJSON(c, http.StatusNotFound, "not_found", "user not found")
		return
	}
	if user.IsAdmin {
		utils.ErrorJSON(c, http.StatusForbidden, "forbidden", "cannot modify another admin")
		return
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.IsAdmin != nil {
		user.IsAdmin = *input.IsAdmin
	}

	if err := database.DB.Save(&user).Error; err != nil {
		utils.ErrorJSON(c, http.StatusInternalServerError, "db_error", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"name":     user.Name,
		"email":    user.Email,
		"is_admin": user.IsAdmin,
	})
}
