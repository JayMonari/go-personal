package main

import (
	"flag"
	"fmt"
	"grpbook/pb"
	"grpbook/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 9001, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer,
		service.LaptopServer{Store: &service.InMemoryLaptopStore{}})
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
