package main

import (

	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	router:=gin.Default()
	gin.SetMode(gin.DebugMode)
	//s:=&http.Server{
	//	Addr:":8080",
	//	Handler:router,
	//	ReadTimeout:10*time.Second,
	//	WriteTimeout:10*time.Second,
	//	MaxHeaderBytes:1<<20,
	//}
	//s.ListenAndServe()
	router.GET("/welcome", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
		name:=context.DefaultQuery("name","guest")
		fmt.Println(name)
	})
	router.Run()

}
/*
表单参数通过 PostForm 方法获取

//form
router.POST("/form", func(c *gin.Context) {
	type := c.DefaultPostForm("type", "alert")//可设置默认值
	msg := c.PostForm("msg")
	title := c.PostForm("title")
	fmt.Println("type is %s, msg is %s, title is %s", type, msg, title)
})

路由群组
	someGroup := router.Group("/someGroup")
    {
        someGroup.GET("/someGet", getting)
		someGroup.POST("/somePost", posting)
	}
 */
type People interface {
	say() string
	look() string
}
