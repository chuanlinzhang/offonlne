package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

//如果一个字段被 binding:"required" 修饰而值却是空的，请求会失败并返回错误。
type Login struct {
	User       string         `form:"user" json:"user " binding:"required"`
	Password   string         `form:"password" json:"password" binding:"required"`
	Uploadfile multipart.File `form:"uploadfile" json:"uploadfile" `
}

type po struct {
	name string
}

func (*po) say() string {
	panic("implement me")
}

func (*po) look() string {
	panic("implement me")
}

func main() {
	router := gin.Default()

	router.POST("/loginJSON", func(context *gin.Context) {
		var json Login
		//使用自动推断，这里是表单提交
		if context.Bind(&json) == nil {
			if json.User == "manu" && json.Password == "123" {
				context.JSON(http.StatusOK, gin.H{"status": "login in"})
			} else {
				context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})
	router.POST("/upload", func(context *gin.Context) {
		file, header, err := context.Request.FormFile("upload")
		filename := header.Filename
		out, err := os.Create("./" + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
	})
	router.POST("/form", func(context *gin.Context) {
		var up Login
		fmt.Println(context.Bind(&up))
		if context.Bind(&up) == nil {
			file, header, err := context.Request.FormFile("uploadfile")
			out, err := os.Create("./" + header.Filename)
			if err != nil {
				log.Fatal(err)
			}
			io.Copy(out, file)
			up.Uploadfile = file
			fmt.Println(up)
		}
		//获取路径参数
		//id := context.Param("id")
		//表单提交上来的，可以用个获取部分信息
		//context.PostForm()
		//context.String(http.StatusOK,"upload succsess")
		context.JSON(http.StatusOK, gin.H{"user": "Lena", "Message": "hey", "Number": 123})
		//context.XML(http.StatusOK, gin.H{"user": "Lena", "Message": "hey", "Number": 123})

	})
	router.Run(":8080")
}

/*
上传多个文件
单个文件上传很简单，别以为多个文件就会很麻烦。依葫芦画瓢，所谓多个文件，无非就是多一次遍历文件，然后一次copy数据存储即可。下面只写handler，省略main函数的初始化路由和开启服务器监听了：

router.POST("/multi/upload", func(c *gin.Context) {
        err := c.Request.ParseMultipartForm(200000)
        if err != nil {
            log.Fatal(err)
        }

        formdata := c.Request.MultipartForm

        files := formdata.File["upload"]
        for i, _ := range files { /
            file, err := files[i].Open()
            defer file.Close()
            if err != nil {
                log.Fatal(err)
            }

            out, err := os.Create(files[i].Filename)

            defer out.Close()

            if err != nil {
                log.Fatal(err)
            }

            _, err = io.Copy(out, file)

            if err != nil {
                log.Fatal(err)
            }

            c.String(http.StatusCreated, "upload successful")

        }

    })
与单个文件上传类似，只不过使用了c.Request.MultipartForm得到文件句柄，再获取文件数据，然后遍历读写。
*/
