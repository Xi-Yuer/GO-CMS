package db

import (
	"database/sql"
	"fmt"
	"github.com/Xi-Yuer/cms/config"
	"github.com/Xi-Yuer/cms/utils"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var DB *sql.DB

func InitDB() error {
	// 初始化数据库连接
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/cms?charset=utf8&parseTime=True", config.Config.DB.NAME, config.Config.DB.PASSWORD, config.Config.DB.HOST, config.Config.DB.PORT))
	if err != nil {
		utils.Log.Panic(err)
	}

	DB = db

	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(time.Minute * 60)
	err = DB.Ping()
	if err != nil {
		utils.Log.Panic(err)
		return err
	}
	utils.Log.Info("数据库连接成功")
	return nil
}
