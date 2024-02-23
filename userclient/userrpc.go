package userclient

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	_ "userrpc/config"
	_ "userrpc/model/mysql"
	"userrpc/userrpc"
)

func NewRpcServer(address int64) {
	grpcServer := grpc.NewServer()
	userrpc.RegisterStreamGreeterServer(grpcServer, new(UserRpcClient))
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", address))
	if err != nil {
		log.Println("服务监听失败", err.Error())
	}
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println("rpc服务创建失败", err.Error())
	}
}
