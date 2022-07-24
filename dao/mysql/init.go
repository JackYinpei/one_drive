package dao

import (
	"fmt"
	"log"
	"os"
	"qqq_one_drive/setting"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Databases(conn *setting.MySQLConfig) {
	connString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conn.User, conn.Password, conn.Port, conn.DB)
	// 初始化GORM日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})
	// Error
	if connString == "" || err != nil {
		zap.L().Panic("mysql lost")
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Panic("mysql lost")
		panic(err)
	}

	sqlDB.SetMaxOpenConns(setting.Conf.MySQLConfig.MaxOpenConns)
	sqlDB.SetMaxOpenConns(setting.Conf.MySQLConfig.MaxIdleConns)
	DB = db

	migration()
}

func migration() {
	_ = DB.AutoMigrate(&User{}, &Note{})
}
