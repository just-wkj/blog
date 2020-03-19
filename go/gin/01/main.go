package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code": 1,
			"msg":  "ok",
		})
	})

	r.GET("/index", func(context *gin.Context) {
		context.HTML(200, "index.tmpl", gin.H{
			"title": "titlexx",
		})
	})
	r.Run(":9988")
}
