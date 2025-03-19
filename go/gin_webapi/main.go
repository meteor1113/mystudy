package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong（最简单的回应）",
		})
	})

	r.GET("/message", func(c *gin.Context) {
		name := c.Query("name")
		data := map[string]interface{}{
			"name": name,
		}
		c.JSON(200, gin.H{
			"status":  1000,
			"message": "响应鹦鹉学舌",
			"data":    data,
		})
	})

	type Info struct {
		Name  string `json:"name"`
		Score int    `json:"score"`
	}

	r.POST("/movie", func(c *gin.Context) {

		// 以Info为模板初始化data
		var data Info

		// 将请求参数绑定到data
		c.BindJSON(&data)

		fmt.Println("=data=>>", data)

		c.JSON(200, gin.H{
			"status":  1000,
			"message": "返回电影名称和评分",
			"data":    data,
		})

	})

	r.Run()
}
