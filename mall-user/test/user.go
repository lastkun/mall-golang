package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"mall/mall-user/proto"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.Nickname, user.Password)
		checkRsp, err := userClient.CheckPassword(context.Background(), &proto.CheckPasswordRequest{
			Password:     "123456",
			EncryptedPwd: user.Password,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRsp.Success)
	}
}

func main() {
	Init()
	TestGetUserList()

	conn.Close()
}
