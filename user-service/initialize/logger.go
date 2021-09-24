package initialize

import "go.uber.org/zap"

func InitGlobalLogger() {
	//使用zap.S()之前需要生成全局logger
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
