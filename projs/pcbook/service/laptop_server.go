package service

import (
	"context"
	"errors"
	"grpbook/pb"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer is an implementation of the GRPC laptop_service
type LaptopServer struct {
	Store LaptopStore
	pb.UnimplementedLaptopServiceServer
}

func (s LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := req.Laptop
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		if _, err := uuid.Parse(laptop.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument,
				"laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal,
				"cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}
	// time.Sleep(time.Second) // NOTE(jay): For testing deadlines
	switch ctx.Err() {
	case context.DeadlineExceeded:
		log.Println("deadline has exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline has exceeded")
	case context.Canceled:
		log.Println("request was canceled")
		return nil, status.Error(codes.DeadlineExceeded, "request was canceled")
	}
	if err := s.Store.Save(laptop); err != nil {
		c := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			c = codes.AlreadyExists
		}
		return nil, status.Errorf(c, "cannot save laptop to the store: %v", err)
	}
	log.Printf("saved laptop with id: %s", laptop.Id)
	return &pb.CreateLaptopResponse{Id: laptop.Id}, nil
}

func (s LaptopServer) SearchLaptop(
	req *pb.SearchLaptopRequest,
	stream pb.LaptopService_SearchLaptopServer,
) error {
	fil := req.GetFilter()
	log.Printf("receive a search-laptop request with filter: %v", fil)
	err := s.Store.Search(
		stream.Context(),
		fil,
		func(lp *pb.Laptop) error {
			res := &pb.SearchLaptopResponse{Laptop: lp}
			if err := stream.Send(res); err != nil {
				return err
			}
			log.Printf("sent laptop with ID: %s", lp.GetId())
			return nil
		})
	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}
	return nil
}
