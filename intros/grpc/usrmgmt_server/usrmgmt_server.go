package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "example.com/go-usrmgmt-grpc/usrmgmt"
	"google.golang.org/grpc"
)

const (
	port = ":9001"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	userList *pb.UserList
}

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{userList: &pb.UserList{}}
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
	u := &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Id:   int32(rand.Intn(10_000)),
	}
	s.userList.Users = append(s.userList.Users, u)
	return u, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	return s.userList, nil
}

func main() {
	if err := NewUserManagementServer().Run(); err != nil {
		log.Fatal(err)
	}
}
