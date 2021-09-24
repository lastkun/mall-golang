package global

import (
	"gorm.io/gorm"
	"mall/user-service/config"
)

var (
	DB           *gorm.DB
	ServerConfig = &config.ServerConfig{}
)

//在某package定义了init方法 当import的时候这个方法自动执行
//func init() {
//	dsn := "root:root@tcp(192.168.1.6:3306)/mall_user?charset=utf8mb4&parseTime=True&loc=Local"
//	//日志配置
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags),
//		logger.Config{
//			SlowThreshold:             time.Second, // 慢 SQL 阈值
//			LogLevel:                  logger.Info, // 日志级别
//			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
//			Colorful:                  true,        // 彩色打印
//		},
//	)
//
//	// 全局模式
//	var err error
//	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: newLogger,
//		//自定义建表命名规则
//		NamingStrategy: schema.NamingStrategy{
//			//TablePrefix: "",
//			SingularTable: true,
//		},
//	})
//	if err != nil {
//		panic(err)
//	}
//
//}
