package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("go in test")
	r := gin.Default()
	r.GET("/contact", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Blog":  "xiarong715.github.io",
			"email": "xiarong2010@163.com",
		})
	})

	// http://localhost:8080/users/lsy
	r.GET("/users/:id", func(c *gin.Context) {
		c.String(200, c.Param("id"))
	})
	// http://localhost:8080/?wechat=flysnow_org
	// http://locahost:8080/?media=wechat&media=weibo
	r.GET("/", func(c *gin.Context) {
		// c.String(200, c.DefaultQuery("wechat", "sos"))
		c.JSON(200, c.QueryArray("media"))
	})
	r.GET("/map", func(c *gin.Context) {
		c.JSON(200, c.QueryMap("ids"))
	})
	r.Run(":8080")
}
