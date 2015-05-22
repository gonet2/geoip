package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	pb "proto"
	"testing"
)

const (
	address    = "localhost:50000"
	default_ip = "106.187.101.141"
)

func TestGeoIP(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGeoIPServiceClient(conn)

	// Contact the server and print out its response.
	r, err := c.QueryCountry(context.Background(), &pb.GeoIP_IP{Ip: default_ip})
	if err != nil {
		log.Fatalf("could not query: %v", err)
	}
	log.Printf("Country: %s", r.Name)
}
