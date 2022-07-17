package main

import (
	"fmt"
	"net/http"
	"os"
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
	gin.SetMode(setting.Conf.Mode)
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.LoadHTMLFiles("./pages/index.html")
	r.Static("/static", "./static")
	// r := gin.Default()
	// r.LoadHTMLFiles("./pages/index.html")
	// r.GET("/hello", controller.GetIndex)
	// r.Run()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
