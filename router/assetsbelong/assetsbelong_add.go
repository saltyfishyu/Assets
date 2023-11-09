package assetsbelong

import (
	"Assets/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AssetsBelongAddJson struct {
	Aid    int      `json:"aid"`
	Domain []string `json:"domain"`
}

func Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		json := &AssetsBelongAddJson{}
		c.BindJSON(&json)
		var domain string
		for _, str := range json.Domain {
			domain = domain + "," + str
		}
		domain = domain[1:]

		result, msg := db.AssetsBelong_Add(json.Aid, domain)
		if result {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    msg,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    msg,
			})
		}
	}
}
