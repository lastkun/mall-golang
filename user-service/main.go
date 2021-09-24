package main

import (
	"google.golang.org/grpc"
	"mall/user-service/initialize"
	"net"

	"mall/user-service/handler"
	"mall/user-service/proto"
)

func main() {
	//初始化logger
	initialize.InitGlobalLogger()
	//初始化配置
	initialize.InitConfig()
	//初始化DB
	initialize.InitDB()

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	//服务端使用0.0.0.0 表示本机中所有的IPV4地址 监听0.0.0.0的端口，就是监听本机中所有IP的端口
	listener, err := net.Listen("tcp", "0.0.0.0:9001")
	if err != nil {
		panic("failed to listen " + err.Error())
	}

	err = server.Serve(listener)
	if err != nil {
		panic("failed to start server" + err.Error())
	}

}
