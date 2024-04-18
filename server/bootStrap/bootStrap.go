package bootStrap

import (
	"github.com/Xi-Yuer/cms/config"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/routers"
	"github.com/Xi-Yuer/cms/utils"
)

func Start() {
	if err := db.InitDB(); err != nil {
		utils.Log.Panic(err)
		return
	}

	if err := utils.File.CheckOrCreateFolder(config.Config.APP.FILEPATH); err != nil {
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
}
