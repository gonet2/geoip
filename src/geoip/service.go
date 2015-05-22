package main

import (
	"errors"
	log "github.com/GameGophers/nsq-logger"
	"golang.org/x/net/context"
	"net"
	pb "proto"
)

const (
	SERVICE = "[GEOIP]"
)

//---------------------------------------------------------- read the following fields only
type City struct {
	City struct {
		Names map[string]string `maxminddb:"names"`
	} `maxminddb:"city"`

	Country struct {
		GeoNameID uint   `maxminddb:"geoname_id"`
		IsoCode   string `maxminddb:"iso_code"`
	} `maxminddb:"country"`

	Subdivisions []struct {
		Names map[string]string `maxminddb:"names"`
	} `maxminddb:"subdivisions"`
}

func _query(ip net.IP) *City {
	city := &City{}
	err := _mmdb.Lookup(ip, city)
	if err != nil {
		log.Error(SERVICE, err)
		return nil
	}

	return city
}

type server struct{}

// 查询IP所属国家
func (s *server) QueryCountry(ctx context.Context, in *pb.GeoIP_IP) (*pb.GeoIP_Name, error) {
	ip := net.ParseIP(in.Ip)
	if city := _query(ip); city != nil {
		return &pb.GeoIP_Name{Name: city.Country.IsoCode}, nil
	}
	return nil, errors.New("cannot query ip")
}

// 查询IP所属城市
func (s *server) QueryCity(ctx context.Context, in *pb.GeoIP_IP) (*pb.GeoIP_Name, error) {
	ip := net.ParseIP(in.Ip)
	if city := _query(ip); city != nil {
		return &pb.GeoIP_Name{Name: city.City.Names["en"]}, nil
	}
	return nil, errors.New("cannot query ip")
}

// 查询IP所属地区(省)
func (s *server) QuerySubdivision(ctx context.Context, in *pb.GeoIP_IP) (*pb.GeoIP_Name, error) {
	ip := net.ParseIP(in.Ip)
	if city := _query(ip); city != nil {
		if len(city.Subdivisions) > 0 {
			return &pb.GeoIP_Name{Name: city.Subdivisions[0].Names["en"]}, nil
		}
		return nil, errors.New("cannot query ip")
	}
	return nil, errors.New("cannot query ip")
}
