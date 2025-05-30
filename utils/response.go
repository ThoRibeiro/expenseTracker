
package utils

import "github.com/gin-gonic/gin"

func ErrorJSON(c *gin.Context, code int, errKey, msg string) {
    c.AbortWithStatusJSON(code, gin.H{"error": errKey, "message": msg})
}
