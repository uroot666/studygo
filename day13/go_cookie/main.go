package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie 并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusOK, gin.H{"error": "err"})
		c.Abort()
	}
}

func main() {

	// 默认使用了两个中间件 Logger(), Recovery()
	r := gin.Default()
	// 服务端要给客户端cookie
	r.GET("cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 设置cookie
			// maxAge int, 单位秒
			// path.cookie 所在目录
			// domain 域名
			// secure 是否只能通过https访问
			// httpOnly 是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
			fmt.Println(c.Cookie("key_cookie"))
		}
		fmt.Printf("cookie 的值是: %s\n", cookie)

	})

	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("abc", "123", 60, "/",
			"localhost", false, true)
		c.String(200, "Login success!")
	})

	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})

	r.Run(":80")

}
