package main

import (
	"fmt"
	"log"
	"net"

	api "github.com/weworksandbox/sg2019/api/payments"
	"google.golang.org/grpc"
)

// main start a gRPC server and waits for connection
func main() {
	// create a listener on TCP port 50053
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50053))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := api.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	api.RegisterPaymentsServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
