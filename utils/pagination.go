
package utils

import (
    "strconv"

    "github.com/gin-gonic/gin"
)

func GetPagination(c *gin.Context) (page, size int) {
    page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
    size, _ = strconv.Atoi(c.DefaultQuery("size", "20"))
    if page < 1 {
        page = 1
    }
    if size < 1 {
        size = 20
    }
    return
}
