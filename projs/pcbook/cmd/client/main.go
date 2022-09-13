package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"grpbook/pb"
	"grpbook/sample"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	testRateLaptop(client)
}

func uploadImage(client pb.LaptopServiceClient, laptopID, imgPath string) {
	file, err := os.Open(imgPath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	if err = stream.Send(&pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptopId:  laptopID,
				ImageType: filepath.Ext(imgPath),
			},
		},
	}); err != nil {
		log.Fatal("cannot send image info: ", err, stream.RecvMsg(nil))
	}

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024)
	for n, err := reader.Read(buf); err != io.EOF; n, err = reader.Read(buf) {
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		if err := stream.Send(&pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buf[:n],
			},
		}); err != nil {
			log.Fatal("cannot send chunk to server: ", err, stream.RecvMsg(nil))
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}
	log.Printf("image uploaded with ID: %s, size %d", res.Id, res.Size)
}

func testUploadImage(client pb.LaptopServiceClient) {
	lp := sample.NewLaptop()
	createLaptop(client, lp)
	uploadImage(client, lp.Id, "/tmp/laptop.jpg")
}

func testCreateLaptop(client pb.LaptopServiceClient) {
	createLaptop(client, sample.NewLaptop())
}

func testSearchLaptop(client pb.LaptopServiceClient) {
	for i := 0; i < 10; i++ {
		createLaptop(client, sample.NewLaptop())
	}
	searchLaptop(client, &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	})
}

func testRateLaptop(c pb.LaptopServiceClient) {
	n := 3
	lpIDs := make([]string, n)
	for i := 0; i < n; i++ {
		lp := sample.NewLaptop()
		lpIDs[i] = lp.Id
		createLaptop(c, lp)
	}

	scores := make([]float64, n)
	for {
		fmt.Println("rate laptop (y/n)?")
		var ans string
		fmt.Scan(&ans)
		if strings.ToLower(ans) != "y" {
			break
		}

		for i := 0; i < n; i++ {
			scores[i] = sample.RandomLaptopScore()
		}

		if err := rateLaptop(c, lpIDs, scores); err != nil {
			log.Fatal(err)
		}
	}
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

func rateLaptop(c pb.LaptopServiceClient, lpIDs []string, scores []float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := c.RateLaptop(ctx)
	if err != nil {
		return fmt.Errorf("cannot rate laptop: %w", err)
	}

	waitResp := make(chan error)
	go func() {
		for {
			res, err := stream.Recv()
			switch {
			case err == io.EOF:
				log.Print("no more responses")
				waitResp <- nil
				return
			case err != nil:
				waitResp <- fmt.Errorf("cannot receive stream response: %v", err)
				return
			}
			log.Println("recieved reponse: ", res)
		}
	}()

	for i, lpID := range lpIDs {
		if err := stream.Send(&pb.RateLaptopRequest{
			LaptopId: lpID,
			Score:    scores[i],
		}); err != nil {
			return fmt.Errorf("cannot send stream request: %v - %v", err, stream.RecvMsg(nil))
		}
	}
	if err = stream.CloseSend(); err != nil {
		return fmt.Errorf("cannot close send: %w", err)
	}
	return <-waitResp
}

func createLaptop(client pb.LaptopServiceClient, lp *pb.Laptop) {
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
