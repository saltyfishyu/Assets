package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "you see me",
		})
	}
}
