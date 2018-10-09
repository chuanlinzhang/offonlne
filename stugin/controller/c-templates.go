package main

import (
	"github.com/gin-gonic/gin"

)

func main() {
	router:=gin.Default()
	//加载模板
	router.LoadHTMLGlob("templates/*")
	//定义路由

	router.GET("/index", func(context *gin.Context) {
		//根据完整的文件名渲染模板，并传递参数
		//视图响应(//根据完整文件名渲染模板，并传递参数)
		//context.HTML(http.StatusOK,"index.html",gin.H{"title":"eoin"})
		//文件响应

		context.File("./t.jpg")

		//重定向
		//context.Redirect(301,"/redirect")
	})
	router.GET("/redirect", func(context *gin.Context) {
		context.JSON(200,gin.H{"name":"eoin"})
	})

	router.Run(":8080")
}
