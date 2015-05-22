package main

import (
	log "github.com/GameGophers/nsq-logger"
	"github.com/oschwald/maxminddb-golang"
	"google.golang.org/grpc"
	"net"
	"os"
	pb "proto"
)

const (
	_port = ":50000"
)

var (
	_mmdb *maxminddb.Reader
	_path = os.Getenv("GOPATH") + "/src/geoip/GeoIP2-City.mmdb"
)

func main() {
	// 载入IP表
	log.Trace(SERVICE, "Loading GEOIP City...")
	reader, err := maxminddb.Open(_path)
	if err != nil {
		log.Critical(SERVICE, err)
		os.Exit(-1)
	}

	_mmdb = reader
	log.Trace(SERVICE, "GEOIP City Load Complete.")

	// 监听
	lis, err := net.Listen("tcp", _port)
	if err != nil {
		log.Critical(SERVICE, err)
	}
	log.Info(SERVICE, "listening on ", lis.Addr())

	// 注册服务
	s := grpc.NewServer()
	pb.RegisterGeoIPServiceServer(s, &server{})

	// 开始服务
	s.Serve(lis)
}
