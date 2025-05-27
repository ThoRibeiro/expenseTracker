package routes

import (
	"expensetracker/controllers"
	"expensetracker/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Public
	r.POST("/users/register", controllers.Register)
	r.POST("/users/login", controllers.Login)

	// Protected routes
	auth := r.Group("/users")
	auth.Use(middleware.AuthRequired())

	// User updates own profile
	auth.PUT("/me", controllers.UpdateMe)

	// Expenses endpoints under authenticated context
	auth.GET("/expenses", controllers.ListExpenses)
	auth.GET("/expenses/:id", controllers.GetExpense)
	auth.GET("/expenses/search", controllers.SearchExpenses)
	auth.POST("/expenses", controllers.CreateExpense)
	auth.PUT("/expenses/:id", controllers.UpdateExpense)
	auth.DELETE("/expenses/:id", controllers.DeleteExpense)
	auth.PUT("/expenses/bulk", controllers.BulkUpdate)
	auth.DELETE("/expenses/bulk", controllers.BulkDelete)

	// Admin-only group
	admin := auth.Group("/admin")
	admin.Use(middleware.AdminOnly())

	// Reset database
	admin.POST("/reset", controllers.ResetDB)

	// Admin can update any non-admin user
	admin.PUT("/users/:id", controllers.AdminUpdateUser)
}
