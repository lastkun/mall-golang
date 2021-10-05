package initialize

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"mall/goods-service/global"
)

//初始化配置信息
func InitConfig() {
	//读取本地配置 本地只配了nacos信息
	viper.AutomaticEnv()
	env := viper.GetString("MALL_ENV")

	configFilePath := "goods-service/config-pro.yml"
	if env == "dev" {
		configFilePath = "goods-service/config-dev.yml"
	}

	v := viper.New()

	v.SetConfigFile(configFilePath)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	//配置文件中的配置读取到struct
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}

	//从nacos拉取服务配置
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		zap.S().Fatal("nacos连接信息有误", err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
	})

	if err != nil {
		zap.S().Panic("获取配置失败 ", err)
	}
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		zap.S().Fatalf("读取配置失败 %s", err.Error())
	}
	zap.S().Info(&global.ServerConfig)
}
