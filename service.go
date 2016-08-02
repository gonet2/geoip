package main

import (
	"errors"
	pb "geoip/proto"
	"net"
	"os"
	"strings"

	"github.com/oschwald/maxminddb-golang"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	SERVICE = "[GEOIP]"
)

var (
	ERROR_CANNOT_QUERY_IP = errors.New("cannot query ip")
)

// read the following fields only
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

type server struct {
	mmdb *maxminddb.Reader
}

func (s *server) init() {
	// 载入IP表
	log.Debug("Loading GEOIP City...")
	reader, err := maxminddb.Open(s.data_path())
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}

	s.mmdb = reader
	log.Debug("GEOIP City Load Complete.")
}

// get correct data path from GOPATH
func (s *server) data_path() (path string) {
	paths := strings.Split(os.Getenv("GOPATH"), ":")
	for k := range paths {
		path = paths[k] + "/src/geoip/GeoIP2-City.mmdb"
		_, err := os.Lstat(path)
		if err == nil {
			return path
		}
	}
	return
}

func (s *server) query(ip net.IP) *City {
	city := &City{}
	err := s.mmdb.Lookup(ip, city)
	if err != nil {
		log.Error(err)
		return nil
	}

	return city
}

// 查询IP所属国家
func (s *server) QueryCountry(ctx context.Context, in *pb.GeoIP_IP) (*pb.GeoIP_Name, error) {
	ip := net.ParseIP(in.Ip)
	if city := s.query(ip); city != nil {
		return &pb.GeoIP_Name{Name: city.Country.IsoCode}, nil
	}
	return nil, ERROR_CANNOT_QUERY_IP
}

// 查询IP所属城市
func (s *server) QueryCity(ctx context.Context, in *pb.GeoIP_IP) (*pb.GeoIP_Name, error) {
	ip := net.ParseIP(in.Ip)
	if city := s.query(ip); city != nil {
		return &pb.GeoIP_Name{Name: city.City.Names["en"]}, nil
	}
	return nil, ERROR_CANNOT_QUERY_IP
}

// 查询IP所属地区(省)
func (s *server) QuerySubdivision(ctx context.Context, in *pb.GeoIP_IP) (*pb.GeoIP_Name, error) {
	ip := net.ParseIP(in.Ip)
	if city := s.query(ip); city != nil {
		if len(city.Subdivisions) > 0 {
			return &pb.GeoIP_Name{Name: city.Subdivisions[0].Names["en"]}, nil
		}
		return nil, ERROR_CANNOT_QUERY_IP
	}
	return nil, ERROR_CANNOT_QUERY_IP
}
