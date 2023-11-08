package assets

import (
	"Assets/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AssetsAddJson struct {
	Content string `json:"content"`
}

func Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		json := &AssetsAddJson{}
		c.BindJSON(&json)

		result := db.Assets_Add(json.Content)
		if result {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "success",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    "false",
			})
		}
	}
}
