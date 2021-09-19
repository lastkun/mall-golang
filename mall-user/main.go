package main

import (
	"google.golang.org/grpc"
	"net"

	"mall/mall-user/handler"
	"mall/mall-user/proto"
)

func main() {
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	//服务端使用0.0.0.0 表示本机中所有的IPV4地址 监听0.0.0.0的端口，就是监听本机中所有IP的端口
	listener, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		panic("failed to listen " + err.Error())
	}

	err = server.Serve(listener)
	if err != nil {
		panic("failed to start server" + err.Error())
	}

}
