package main

import (
	"context"
	"github.com/Xi-Yuer/cms/config"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/routers"
	"github.com/Xi-Yuer/cms/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := db.InitDB()
	if err != nil {
		utils.Log.Panic(err)
		return
	}

	r := routers.SetUpRouters()
	go func() {
		err := r.Run(config.Config.APP.PORT)
		if err != nil {
			utils.Log.Panic(err)
		}
	}()

	utils.Log.Info("服务器启动成功，运行端口", config.Config.APP.PORT)
	utils.Log.Info("接口文档地址", config.Config.APP.SWAGPATH)
	// 优雅退出程序
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}
