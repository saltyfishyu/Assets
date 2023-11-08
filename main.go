package main

import (
	"Assets/cors"
	"Assets/db"
	"Assets/router/assets"
	"Assets/router/assetsbelong"
	"Assets/router/index"
	"github.com/gin-gonic/gin"
)

func main() {
	db.New()

	// 切换release环境 不加是debug环境
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	//静态资源
	//r.LoadHTMLGlob("templates/*")
	//r.StaticFile("/favicon.ico", "./templates/favicon.ico")
	//r.StaticFS("/static", http.Dir("./static"))

	r.Use(cors.Cors_backend())

	//入口路由
	indexrouter := r.Group("/")
	{
		indexrouter.GET("/", index.Index())
		indexrouter.POST("/", index.Index())
	}

	assetsrouter := r.Group("/assets")
	{
		assetsrouter.GET("/all", assets.All())
		assetsrouter.POST("/add", assets.Add())
	}
	assetsbelongrouter := r.Group("/assetsbelong")
	{
		assetsbelongrouter.POST("/add", assetsbelong.Add())
	}

	r.Run("127.0.0.1:7070")
}
