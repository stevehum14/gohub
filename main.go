package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
	"net/http"
)

func init()  {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main()  {
	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env,"env","","加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Logger
	bootstrap.SetupLogger()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化 db
	bootstrap.SetupDB()

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
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
