package main

import (
	log "github.com/GameGophers/nsq-logger"
	"google.golang.org/grpc"
	"net"
	"os"
	pb "proto"
)

const (
	_port = ":50000"
)

func main() {
	log.SetPrefix(SERVICE)
	// 监听
	lis, err := net.Listen("tcp", _port)
	if err != nil {
		log.Critical(err)
		os.Exit(-1)
	}
	log.Info("listening on ", lis.Addr())

	// 注册服务
	s := grpc.NewServer()
	ins := &server{}
	ins.init()
	pb.RegisterGeoIPServiceServer(s, ins)

	// 开始服务
	s.Serve(lis)
}
