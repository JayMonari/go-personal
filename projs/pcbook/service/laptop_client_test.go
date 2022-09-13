package service_test

import (
	"bufio"
	"context"
	"fmt"
	"grpbook/pb"
	"grpbook/sample"
	"grpbook/serializer"
	"grpbook/service"
	"io"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestClientRateLaptop(t *testing.T) {
	t.Parallel()

	lpStore := service.InMemoryLaptopStore{}
	laptop := sample.NewLaptop()
	require.NoError(t, lpStore.Save(laptop))

	stream, err := newTestLaptopClient(
		t,
		startTestLaptopServer(t, &lpStore, nil, service.NewInMemoryRatingStore()),
	).RateLaptop(context.Background())
	require.NoError(t, err)

	scores := []float64{8, 7.5, 10}
	avgs := []float64{8, 7.75, 8.5}
	n := len(scores)
	for i := 0; i < n; i++ {
		require.NoError(t, stream.Send(&pb.RateLaptopRequest{
			LaptopId: laptop.Id,
			Score:    scores[i],
		}))
	}

	require.NoError(t, stream.CloseSend())

	for i := 0; ; i++ {
		res, err := stream.Recv()
		if err == io.EOF {
			require.Equal(t, n, i)
			return
		}

		require.NoError(t, err)
		require.Equal(t, laptop.Id, res.LaptopId)
		require.Equal(t, uint32(i+1), res.RatedCount)
		require.Equal(t, avgs[i], res.AverageScore)
	}
}

func TestClientUploadImage(t *testing.T) {
	t.Parallel()

	lpStore := service.InMemoryLaptopStore{}
	laptop := sample.NewLaptop()
	require.NoError(t, lpStore.Save(laptop))

	imgDir, imgType := "../img", ".jpg"
	file, err := os.Open("/tmp/laptop.jpg")
	require.NoError(t, err)
	defer file.Close()

	stream, err := newTestLaptopClient(
		t,
		startTestLaptopServer(t, &lpStore, service.NewDiskImageStore(imgDir), nil),
	).UploadImage(context.Background())
	require.NoError(t, err)

	if err = stream.Send(&pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptopId:  laptop.Id,
				ImageType: imgType,
			},
		},
	}); err != nil {
		t.Fatal("cannot send image info: ", err, stream.RecvMsg(nil))
	}

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024)
	size := 0
	for n, err := reader.Read(buf); err != io.EOF; n, err = reader.Read(buf) {
		require.NoError(t, err)
		size += n
		if err := stream.Send(&pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buf[:n],
			},
		}); err != nil {
			t.Fatal("cannot send chunk to server: ", err, stream.RecvMsg(nil))
		}
	}

	res, err := stream.CloseAndRecv()
	require.NoError(t, err)
	require.NotZero(t, res.Id)
	require.EqualValues(t, size, res.Size)

	require.FileExists(t, fmt.Sprintf("%s/%s%s", imgDir, res.Id, imgType))
	require.NoError(t, os.Remove(fmt.Sprintf("%s/%s%s", imgDir, res.Id, imgType)))
}

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	lpStore := service.InMemoryLaptopStore{}
	addr := startTestLaptopServer(t, &lpStore, nil, nil)

	laptop := sample.NewLaptop()
	wantID := laptop.Id
	res, err := newTestLaptopClient(t, addr).
		CreateLaptop(context.Background(),
			&pb.CreateLaptopRequest{Laptop: laptop})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.Id, wantID)

	other, err := lpStore.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	requireSameLaptop(t, laptop, other)
}

func TestClient_SearchLaptop(t *testing.T) {
	t.Parallel()

	store := &service.InMemoryLaptopStore{}
	expectedIDs := make(map[string]bool)
	for i := 0; i < 6; i++ {
		laptop := sample.NewLaptop()
		switch i {
		case 0:
			laptop.PriceUsd = 2500
		case 1:
			laptop.Cpu.NumberCores = 2
		case 2:
			laptop.Cpu.MinGhz = 2.0
		case 3:
			laptop.Ram = &pb.Memory{Value: 4096, Unit: pb.Memory_GIGABYTE}
		case 4:
			laptop.PriceUsd = 1999
			laptop.Cpu.NumberCores = 4
			laptop.Cpu.MinGhz = 2.5
			laptop.Cpu.MaxGhz = 4.5
			laptop.Ram = &pb.Memory{Value: 16, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptop.Id] = true
		case 5:
			laptop.PriceUsd = 2000
			laptop.Cpu.NumberCores = 6
			laptop.Cpu.MinGhz = 2.8
			laptop.Cpu.MaxGhz = 5.0
			laptop.Ram = &pb.Memory{Value: 64, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptop.Id] = true
		}
		err := store.Save(laptop)
		require.NoError(t, err)
	}

	addr := startTestLaptopServer(t, store, nil, nil)
	client := newTestLaptopClient(t, addr)
	stream, err := client.SearchLaptop(
		context.Background(),
		&pb.SearchLaptopRequest{
			Filter: &pb.Filter{
				MaxPriceUsd: 2000,
				MinCpuCores: 4,
				MinCpuGhz:   2.2,
				MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
			},
		},
	)
	require.NoError(t, err)
	found := 0
	for res, err := stream.Recv(); err != io.EOF; res, err = stream.Recv() {
		require.NoError(t, err)
		require.Contains(t, expectedIDs, res.GetLaptop().GetId())
		found++
	}
	require.Equal(t, len(expectedIDs), found)
}

func startTestLaptopServer(
	t *testing.T,
	lpStore service.LaptopStore,
	iStore service.ImageStore,
	rStore service.RatingStore,
) string {
	svr := service.NewLaptopServer(lpStore, iStore, rStore)
	grpcSvr := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcSvr, svr)
	lis, err := net.Listen("tcp", ":0") // random available port
	require.NoError(t, err)

	go grpcSvr.Serve(lis)
	return lis.Addr().String()
}

func newTestLaptopClient(t *testing.T, address string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, lp1, lp2 *pb.Laptop) {
	json1, err := serializer.ProtobufToJSON(lp1)
	require.NoError(t, err)
	json2, err := serializer.ProtobufToJSON(lp1)
	require.NoError(t, err)
	require.Equal(t, json1, json2)
}
