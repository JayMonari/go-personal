package main

import (
	"context"
	"flag"
	"grpbook/pb"
	"grpbook/sample"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddr := flag.String("addr", "0.0.0.0:9001", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddr)

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	lp := sample.NewLaptop()
	lp.Id = ""
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := pb.NewLaptopServiceClient(conn).
		CreateLaptop(ctx, &pb.CreateLaptopRequest{Laptop: lp})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Println("laptop already exists")
		} else {
			log.Fatal(err)
		}
		return
	}
	log.Printf("created laptop with id: %s", res.Id)
}
