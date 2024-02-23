package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	_ "zg5/day1/rpc/userrpc/config"
	_ "zg5/day1/rpc/userrpc/model/mysql"
	"zg5/day1/rpc/userrpc/userclient"
	pb "zg5/day1/rpc/userrpc/userrpc"
)

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterStreamGreeterServer(grpcServer, new(userclient.UserRpcClient))
	listen, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Println("服务监听失败", err.Error())
	}
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println("rpc服务创建失败", err.Error())
	}
}
