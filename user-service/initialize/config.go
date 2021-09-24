package initialize

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mall/user-service/global"
)

//获取环境变量值
func GetEnvTag(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func InitConfig() {
	env := GetEnvTag("MALL_ENV")
	configFilePath := "user-service/config-pro.yml"
	if env == "dev" {
		configFilePath = "user-service/config-dev.yml"
	}

	v := viper.New()

	v.SetConfigFile(configFilePath)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	//配置文件中的配置读取到struct
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}

	zap.S().Info(global.ServerConfig)

}
