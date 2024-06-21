package gormDB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	userModel "micro/model/user"
	"os"
	"time"
)

var DB *gorm.DB

func NewGorm() *gorm.DB {
	if DB != nil {
		return DB
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:2214380963Wx!!@tcp(127.0.0.1:3306)/micro_cms?charset=utf8&parseTime=True&loc=Local", // 数据源名称
		DefaultStringSize:         256,                                                                                       // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                      // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                      // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                      // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                     // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err := db.AutoMigrate(&userModel.User{}); err != nil {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return db
}
