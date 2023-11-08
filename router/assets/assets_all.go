package assets

import (
	"Assets/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func All() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := db.Assets_All()
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"data":   data,
		})
	}
}
