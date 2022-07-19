package main

import (
	"context"
	"log"
	"time"

	pb "example.com/go-usrmgmt-grpc/usrmgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	newUsers := map[string]int32{}
	newUsers["Alice"] = 43
	newUsers["Bob"] = 30
	for n, a := range newUsers {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: n, Age: a})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(r)
	}
	r, err := c.GetUsers(ctx, &pb.GetUsersParams{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r.GetUsers())
}
