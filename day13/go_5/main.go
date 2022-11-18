package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 创建路由
	r := gin.Default()
	// api参数
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+"is"+action)
	})

	r.GET("/welcom", func(c *gin.Context) {
		// DefaultQuery 第二个参数是默认值
		name := c.DefaultQuery("name", "Jack")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))

	})

	r.POST("/form", func(c *gin.Context) {
		// 表单参数设置默认值
		type1 := c.DefaultPostForm("type", "alert")
		// 接收其它的
		username := c.PostForm("username")
		password := c.PostForm("password")
		hobbys := c.PostFormArray("hobby")
		c.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s, hobbys is %v",
			type1, username, password, hobbys))
	})

	r.POST("/upload", func(c *gin.Context) {
		// 表单取文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		// 传到项目根目录，名字就用本身的
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("%s upload", file.Filename))
	})
	r.Run(":80")

}

func getting(c *gin.Context) {

}
