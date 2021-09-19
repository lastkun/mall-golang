package main

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"time"

	"mall/mall-user/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func createMd5(code string) string {
	hash := md5.New()
	io.WriteString(hash, code)
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	dsn := "root:root@tcp(192.168.1.6:3306)/mall_user?charset=utf8mb4&parseTime=True&loc=Local"
	//日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		//自定义建表命名规则
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "",
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&model.User{}) //此处应该有sql语句
}
