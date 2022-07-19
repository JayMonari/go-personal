package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"

	pb "example.com/go-usrmgmt-grpc/usrmgmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	port = ":9001"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (svc *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, svc)
	log.Println("server listening on", lis.Addr())
	return s.Serve(lis)
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Recieved: %q\n", in.GetName())
	b, err := os.ReadFile("users.json")
	ul := pb.UserList{}
	u := &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Id:   int32(rand.Intn(10_000)),
	}
	switch {
	case err == nil:
		if err = protojson.Unmarshal(b, &ul); err != nil {
			log.Fatal(err)
		}
		ul.Users = append(ul.Users, u)
		marshal(&ul)
	case os.IsNotExist(err):
		log.Println("File not found. Creating a new file")
		ul.Users = append(ul.Users, u)
		marshal(&ul)
	default:
		log.Fatal(err)
	}
	return u, nil
}

func marshal(ul *pb.UserList) {
	jsonb, err := protojson.Marshal(ul)
	if err != nil {
		log.Fatal(err)
	}
	if err = os.WriteFile("users.json", jsonb, 0664); err != nil {
		log.Fatal(err)
	}
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	jsonb, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}
	ul := pb.UserList{}
	if err = protojson.Unmarshal(jsonb, &ul); err != nil {
		log.Fatal(err)
	}
	return &ul, nil
}

func main() {
	if err := (&UserManagementServer{}).Run(); err != nil {
		log.Fatal(err)
	}
}
