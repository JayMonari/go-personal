package main

import (
	"context"
	"flag"
	"grpbook/pb"
	"grpbook/sample"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddr := flag.String("addr", "0.0.0.0:9001", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddr)

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewLaptopServiceClient(conn)
	for i := 0; i < 10; i++ {
		createLaptop(client)
	}

	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}
	searchLaptop(client, filter)
}

func searchLaptop(client pb.LaptopServiceClient, filter *pb.Filter) {
	log.Print("search filter: ", filter)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.SearchLaptop(ctx, &pb.SearchLaptopRequest{Filter: filter})
	if err != nil {
		log.Fatal("cannot search laptop: ", err)
	}
	for res, err := stream.Recv(); err != io.EOF; res, err = stream.Recv() {
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}
		{
			lp := res.Laptop
			log.Printf(`
- found: %v
	+ brand: %v
	+ name: %v
	+ cpu cores: %v
	+ cpu min ghz: %v
	+ ram: %v
	+ price: %v`[1:],
				lp.Id, lp.Brand, lp.Name, lp.Cpu.NumberCores, lp.Cpu.MinGhz, lp.Ram,
				lp.PriceUsd)
		}
	}
}

func createLaptop(client pb.LaptopServiceClient) {
	lp := sample.NewLaptop()
	lp.Id = ""
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := client.CreateLaptop(ctx, &pb.CreateLaptopRequest{Laptop: lp})
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
