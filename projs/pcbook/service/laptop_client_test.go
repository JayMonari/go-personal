package service_test

import (
	"context"
	"grpbook/pb"
	"grpbook/sample"
	"grpbook/serializer"
	"grpbook/service"
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	lpServer, addr := startTestLaptopServer(t, &service.InMemoryLaptopStore{})

	laptop := sample.NewLaptop()
	wantID := laptop.Id
	res, err := newTestLaptopClient(t, addr).
		CreateLaptop(context.Background(),
			&pb.CreateLaptopRequest{Laptop: laptop})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.Id, wantID)

	other, err := lpServer.Store.Find(res.Id)
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

	_, addr := startTestLaptopServer(t, store)
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

func startTestLaptopServer(t *testing.T, store service.LaptopStore) (*service.LaptopServer, string) {
	svr := service.LaptopServer{Store: store}
	grpcSvr := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcSvr, svr)
	lis, err := net.Listen("tcp", ":0") // random available port
	require.NoError(t, err)

	go grpcSvr.Serve(lis)
	return &svr, lis.Addr().String()
}

func newTestLaptopClient(t *testing.T, address string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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
