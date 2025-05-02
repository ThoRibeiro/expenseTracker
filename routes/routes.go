
package routes

import (
    "expensetracker/controllers"
    "expensetracker/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/users", controllers.Register)
    r.POST("/login", controllers.Login)

    auth := r.Group("/")
    auth.Use(middleware.AuthRequired())

    // user self update
    auth.PUT("/me", controllers.Register) // simple reuse, could be separate

    // expenses
    auth.GET("/expenses", controllers.ListExpenses)
    auth.GET("/expenses/:id", controllers.GetExpense)
    auth.GET("/expenses/search", controllers.SearchExpenses)
    auth.POST("/expenses", controllers.CreateExpense)
    auth.PUT("/expenses/:id", controllers.UpdateExpense)
    auth.DELETE("/expenses/:id", controllers.DeleteExpense)
    auth.PUT("/expenses/bulk", controllers.BulkUpdate)
    auth.DELETE("/expenses/bulk", controllers.BulkDelete)

    admin := auth.Group("/admin")
    admin.Use(middleware.AdminOnly())
    admin.POST("/reset", controllers.ResetDB)
}
