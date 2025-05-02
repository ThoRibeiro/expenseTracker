
package controllers

import (
    "expensetracker/database"
    "expensetracker/models"
    "expensetracker/utils"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "time"
)

// @Summary List expenses
// @Tags expenses
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param size query int false "Page size"
// @Param category query string false "Category"
// @Param min query number false "Min amount"
// @Param max query number false "Max amount"
// @Param from query string false "From date RFC3339"
// @Param to query string false "To date RFC3339"
// @Success 200 {array} models.Expense
// @Router /expenses [get]
func ListExpenses(c *gin.Context) {
    userID := c.GetUint("user_id")
    page, size := utils.GetPagination(c)
    var expenses []models.Expense
    q := database.DB.Where("user_id = ?", userID)
    if cat := c.Query("category"); cat != "" {
        q = q.Where("category = ?", cat)
    }
    if min := c.Query("min"); min != "" {
        q = q.Where("amount >= ?", min)
    }
    if max := c.Query("max"); max != "" {
        q = q.Where("amount <= ?", max)
    }
    if from := c.Query("from"); from != "" {
        if t, e := time.Parse(time.RFC3339, from); e == nil {
            q = q.Where("date >= ?", t)
        }
    }
    if to := c.Query("to"); to != "" {
        if t, e := time.Parse(time.RFC3339, to); e == nil {
            q = q.Where("date <= ?", t)
        }
    }
    q.Offset((page - 1) * size).Limit(size).Order("date desc").Find(&expenses)
    c.JSON(http.StatusOK, expenses)
}

// @Summary Get expense by id
// @Tags expenses
// @Security BearerAuth
// @Param id path int true "Expense ID"
// @Success 200 {object} models.Expense
// @Router /expenses/{id} [get]
func GetExpense(c *gin.Context) {
    userID := c.GetUint("user_id")
    id := c.Param("id")
    var expense models.Expense
    if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&expense).Error; err != nil {
        utils.ErrorJSON(c, http.StatusNotFound, "not_found", "expense not found")
        return
    }
    c.JSON(http.StatusOK, expense)
}

// @Summary Search expenses
// @Tags expenses
// @Security BearerAuth
// @Param q query string true "Query"
// @Success 200 {array} models.Expense
// @Router /expenses/search [get]
func SearchExpenses(c *gin.Context) {
    userID := c.GetUint("user_id")
    qStr := c.Query("q")
    var expenses []models.Expense
    database.DB.Where("user_id = ? AND (label LIKE ? OR notes LIKE ? OR category LIKE ?)",
        userID, "%"+qStr+"%", "%"+qStr+"%", "%"+qStr+"%").Find(&expenses)
    c.JSON(http.StatusOK, expenses)
}

// @Summary Create expense
// @Tags expenses
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param expense body models.Expense true "Expense"
// @Success 201 {object} models.Expense
// @Router /expenses [post]
func CreateExpense(c *gin.Context) {
    userID := c.GetUint("user_id")
    var input models.Expense
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "invalid_input", err.Error())
        return
    }
    input.UserID = userID
    if err := database.DB.Create(&input).Error; err != nil {
        utils.ErrorJSON(c, http.StatusInternalServerError, "db_error", err.Error())
        return
    }
    c.JSON(http.StatusCreated, input)
}

// @Summary Update expense
// @Tags expenses
// @Security BearerAuth
// @Accept json
// @Param id path int true "Expense ID"
// @Param expense body models.Expense true "Expense"
// @Success 200 {object} models.Expense
// @Router /expenses/{id} [put]
func UpdateExpense(c *gin.Context) {
    userID := c.GetUint("user_id")
    id := c.Param("id")
    var expense models.Expense
    if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&expense).Error; err != nil {
        utils.ErrorJSON(c, http.StatusNotFound, "not_found", "expense not found")
        return
    }
    var input models.Expense
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "invalid_input", err.Error())
        return
    }
    input.ID = expense.ID
    input.UserID = expense.UserID
    if err := database.DB.Model(&expense).Updates(input).Error; err != nil {
        utils.ErrorJSON(c, http.StatusInternalServerError, "db_error", err.Error())
        return
    }
    c.JSON(http.StatusOK, expense)
}

// @Summary Delete expense
// @Tags expenses
// @Security BearerAuth
// @Param id path int true "Expense ID"
// @Success 204
// @Router /expenses/{id} [delete]
func DeleteExpense(c *gin.Context) {
    userID := c.GetUint("user_id")
    id := c.Param("id")
    if err := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Expense{}).Error; err != nil {
        utils.ErrorJSON(c, http.StatusInternalServerError, "db_error", err.Error())
        return
    }
    c.Status(http.StatusNoContent)
}

// @Summary Bulk update expenses
// @Tags expenses
// @Security BearerAuth
// @Accept json
// @Param payload body map[string]interface{} true "ids and fields"
// @Success 200 {string} string "updated"
// @Router /expenses/bulk [put]
func BulkUpdate(c *gin.Context) {
    userID := c.GetUint("user_id")
    var payload struct {
        IDs    []uint                `json:"ids" binding:"required"`
        Fields map[string]interface{} `json:"fields" binding:"required"`
    }
    if err := c.ShouldBindJSON(&payload); err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "invalid_input", err.Error())
        return
    }
    res := database.DB.Model(&models.Expense{}).Where("user_id = ? AND id IN ?", userID, payload.IDs).Updates(payload.Fields)
    if res.Error != nil {
        utils.ErrorJSON(c, http.StatusInternalServerError, "db_error", res.Error.Error())
        return
    }
    c.JSON(http.StatusOK, gin.H{"updated": res.RowsAffected})
}

// @Summary Bulk delete expenses
// @Tags expenses
// @Security BearerAuth
// @Accept json
// @Param payload body map[string][]uint true "ids"
// @Success 200 {string} string "deleted"
// @Router /expenses/bulk [delete]
func BulkDelete(c *gin.Context) {
    userID := c.GetUint("user_id")
    var payload struct {
        IDs []uint `json:"ids" binding:"required"`
    }
    if err := c.ShouldBindJSON(&payload); err != nil {
        utils.ErrorJSON(c, http.StatusBadRequest, "invalid_input", err.Error())
        return
    }
    res := database.DB.Where("user_id = ? AND id IN ?", userID, payload.IDs).Delete(&models.Expense{})
    if res.Error != nil {
        utils.ErrorJSON(c, http.StatusInternalServerError, "db_error", res.Error.Error())
        return
    }
    c.JSON(http.StatusOK, gin.H{"deleted": res.RowsAffected})
}

// @Summary Reset database
// @Tags admin
// @Security BearerAuth
// @Success 200
// @Router /admin/reset [post]
func ResetDB(c *gin.Context) {
    database.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Expense{})
    database.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Where("is_admin = false").Delete(&models.User{})
    c.JSON(http.StatusOK, gin.H{"status": "database reset"})
}
