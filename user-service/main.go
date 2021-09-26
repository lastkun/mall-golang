package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"syscall"

	"mall/user-service/global"
	"mall/user-service/handler"
	"mall/user-service/initialize"
	"mall/user-service/proto"
	"mall/user-service/utils"
)

func main() {
	//初始化logger
	initialize.InitGlobalLogger()
	//初始化配置
	initialize.InitConfig()
	//初始化DB
	initialize.InitDB()

	//动态获取可用端口号
	Port, _ := utils.GetFreePort()

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	//服务端使用0.0.0.0 表示本机中所有的IPV4地址 监听0.0.0.0的端口，就是监听本机中所有IP的端口
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", Port))
	if err != nil {
		panic("failed to listen " + err.Error())
	}

	zap.S().Infof("start [user-service] on port : %d", Port)

	//健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//服务注册
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host,
		global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("192.168.1.2:%d", Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serviceID := fmt.Sprintf("%s", uuid.NewV4()) //使用uuid作为服务的id
	registration.ID = serviceID
	registration.Port = Port
	registration.Tags = []string{"user", "service"}
	registration.Address = "192.168.1.2"
	registration.Check = check
	//1. 如何启动两个服务
	//2. 即使我能够通过终端启动两个服务，但是注册到consul中的时候也会被覆盖
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	go func() {
		//serve方法是阻塞的
		err = server.Serve(listener)
		if err != nil {
			panic("failed to start server" + err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = client.Agent().ServiceDeregister(serviceID)
	if err != nil {
		zap.S().Errorf("[user-service]注销失败 : %s", err.Error())
	}
	zap.S().Info("[user-service]注销成功")
}
