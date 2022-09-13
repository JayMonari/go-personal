package service

import (
	"bytes"
	"context"
	"errors"
	"grpbook/pb"
	"io"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// maximum of 1 mebibyte
const maxImageSize = 1 << 20

// LaptopServer is an implementation of the GRPC laptop_service
type LaptopServer struct {
	lpStore     LaptopStore
	imgStore    ImageStore
	ratingStore RatingStore

	pb.UnimplementedLaptopServiceServer
}

func NewLaptopServer(lpStore LaptopStore, iStore ImageStore, rStore RatingStore,
) *LaptopServer {
	return &LaptopServer{
		lpStore:     lpStore,
		imgStore:    iStore,
		ratingStore: rStore,
	}
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
	if err := contextError(ctx); err != nil {
		return nil, err
	}
	if err := s.lpStore.Save(laptop); err != nil {
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
	err := s.lpStore.Search(
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

// UploadImage is a client-streaming RPC to pload a laptop image
func (s *LaptopServer) UploadImage(stream pb.LaptopService_UploadImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive image info"))
	}

	laptopID := req.GetInfo().LaptopId
	imgType := req.GetInfo().ImageType
	log.Printf("recieve an upload-image request for laptop %s with image type %s",
		laptopID, imgType)

	laptop, err := s.lpStore.Find(laptopID)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
	}
	if laptop == nil {
		return logError(status.Errorf(codes.InvalidArgument, "laptop %s doesn't exist", laptopID))
	}

	var imgData bytes.Buffer
	imgSize := 0
	for req, err := stream.Recv(); err != io.EOF; req, err = stream.Recv() {
		if err := contextError(stream.Context()); err != nil {
			return err
		}
		log.Print("waiting to receive more data")
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		chunk := req.GetChunkData()
		log.Printf("received a chunk with size: %d", len(chunk))
		if imgSize += len(chunk); imgSize > maxImageSize {
			return logError(status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imgSize, maxImageSize))
		}
		if _, err := imgData.Write(chunk); err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	imgID, err := s.imgStore.Save(laptop.Id, imgType, imgData)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save image to the store: %v", err))
	}

	res := &pb.UploadImageResponse{
		Id:   imgID,
		Size: uint32(imgSize),
	}
	if err = stream.SendAndClose(res); err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}
	log.Printf("saved image with ID: %s, size: %d", imgID, imgSize)
	return nil
}

// RateLaptop is a bidirectional-streaming RPC that allows client to rate a
// stream of laptops with a score and returns a stream of average score for
// each rated laptop.
func (s *LaptopServer) RateLaptop(stream pb.LaptopService_RateLaptopServer) error {
	for req, err := stream.Recv(); err != io.EOF; req, err = stream.Recv() {
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}
		if err := contextError(stream.Context()); err != nil {
			return err
		}

		lpID := req.GetLaptopId()
		score := req.GetScore()
		log.Printf("received a rate-laptop request: id = %q, score = %.2f", lpID, score)

		found, err := s.lpStore.Find(lpID)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
		}
		if found == nil {
			return logError(status.Errorf(codes.NotFound, "laptopID %q is not found", lpID))
		}

		rating, err := s.ratingStore.Add(lpID, score)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot add rating to the store: %v", err))
		}

		if err = stream.Send(&pb.RateLaptopResponse{
			LaptopId:     lpID,
			RatedCount:   rating.Count,
			AverageScore: rating.Sum / float64(rating.Count),
		}); err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot send stream response: %v", err))
		}
	}
	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.DeadlineExceeded:
		log.Println("deadline has exceeded")
		return status.Error(codes.DeadlineExceeded, "deadline has exceeded")
	case context.Canceled:
		log.Println("request was canceled")
		return status.Error(codes.DeadlineExceeded, "request was canceled")
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Println(err)
	}
	return err
}
