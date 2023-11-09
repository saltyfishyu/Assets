package assets

import (
	"Assets/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func All() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, msg := db.Assets_All()
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    msg,
			"data":   data,
		})
	}
}
