package service_test

import (
	"context"
	"grpbook/pb"
	"grpbook/sample"
	"grpbook/serializer"
	"grpbook/service"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	lpServer, addr := startTestLaptopServer(t)

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

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	svr := service.LaptopServer{Store: &service.InMemoryLaptopStore{}}
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
