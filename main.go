package main

import (
	"fmt"
	"net/http"
	"os"
	dao "qqq_one_drive/dao/mysql"
	"qqq_one_drive/logger"
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

	// TODO 数据库链接
	dao.Databases(setting.Conf.MySQLConfig)

	gin.SetMode(setting.Conf.Mode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.LoadHTMLFiles("./pages/index.html")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.Run(":8080")
}
