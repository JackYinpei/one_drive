package main

import (
	"fmt"
	"net/http"
	"os"
	"qqq_one_drive/api"
	dao "qqq_one_drive/dao/mysql"
	"qqq_one_drive/logger"
	"qqq_one_drive/middlewares"
	"qqq_one_drive/pkg/snowflake"
	"qqq_one_drive/setting"

	"github.com/gin-gonic/gin"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: bluebell config.yaml")
		return
	}
	// 加载配置
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// snowflake Init
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// TODO 数据库链接
	dao.Databases(setting.Conf.MySQLConfig)

	gin.SetMode(setting.Conf.Mode)
	r := gin.Default()
	// r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// r.Use(logger.GinRecovery(true))
	r.LoadHTMLFiles("./pages/index.html")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)
	afterLogin := r.Group("/show")
	afterLogin.Use(middlewares.JWTAuthMiddleware())
	{
		afterLogin.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "你妈的， 登录成功")
		})
		afterLogin.POST("/note", api.PostNote)
		afterLogin.GET("/note", api.GetNote)
	}
	r.Run(":8080")
}
