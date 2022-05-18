package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
	"net/http"
)

func main()  {
	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化路由绑定
	 bootstrap.SetupRoute(router)


	// 注册一个路由
	router.GET("/", func(c *gin.Context) {
		// 以 json 格式响应
		c.JSON(http.StatusOK,gin.H{
			"Hello":"World!",
		})
	})
	// 运行服务
	err := router.Run(":8000")
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
