package service

import (
	"context"
	"errors"
	"fmt"
	"grpbook/pb"
	"log"
	"sync"

	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("record already exists")

// LaptopStore is an interface to store laptops.
type LaptopStore interface {
	// Save persists the laptop into the store.
	Save(lp *pb.Laptop) error
	// Find returns a laptop by it's ID or else throws an error.
	Find(id string) (*pb.Laptop, error)
	// Search searches for laptops with filter, returns one by one via found.
	Search(ctx context.Context, fil *pb.Filter, found func(lp *pb.Laptop) error) error
}

// InMemoryLaptopStore stores laptop in memory.
type InMemoryLaptopStore struct {
	mu   sync.RWMutex
	data map[string]*pb.Laptop
}

// Save persists the laptop to the Store.
func (s *InMemoryLaptopStore) Save(lp *pb.Laptop) error {
	if s.data == nil {
		s.data = make(map[string]*pb.Laptop)
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[lp.Id]; exists {
		return ErrAlreadyExists
	}
	other, err := deepCopy(lp)
	if err != nil {
		return err
	}
	s.data[other.Id] = other
	return nil
}

// Find find a laptop by ID or nil. It only errors if the laptop couldn't be
// copied to the result.
func (s *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	lp := s.data[id]
	if lp == nil {
		return nil, nil
	}
	return deepCopy(lp)
}

// Search searches for laptops with filter, returns one by one via found.
func (s *InMemoryLaptopStore) Search(
	ctx context.Context,
	fil *pb.Filter,
	found func(lp *pb.Laptop) error,
) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, lp := range s.data {
		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			log.Print("context is cancelled")
			return errors.New("context is cancelled")
		}
		// time.Sleep(time.Second) NOTE(jay): For testing context deadline
		if isQualified(fil, lp) {
			other, err := deepCopy(lp)
			if err != nil {
				return err
			}
			if err = found(other); err != nil {
				return err
			}
		}
	}
	return nil
}

func isQualified(f *pb.Filter, lp *pb.Laptop) bool {
	switch {
	case lp.GetPriceUsd() > f.GetMaxPriceUsd():
		return false
	case lp.GetCpu().GetNumberCores() < f.GetMinCpuCores():
		return false
	case lp.GetCpu().GetMinGhz() < f.GetMinCpuGhz():
		return false
	case toBit(lp.GetRam()) < toBit(f.GetMinRam()):
		return false
	default:
		return true
	}
}

func toBit(mem *pb.Memory) uint64 {
	val := mem.GetValue()
	switch mem.GetUnit() {
	case pb.Memory_BIT:
		return val
	case pb.Memory_BYTE:
		return val << 3 // 8 = 2^3
	case pb.Memory_KILOBYTE:
		return val << 13 // 1024 * 8 = 2^10 * 2^3 = 2^13
	case pb.Memory_MEGABYTE:
		return val << 23
	case pb.Memory_GIGABYTE:
		return val << 33
	case pb.Memory_TERABYTE:
		return val << 43
	default:
		return 0
	}
}

func deepCopy(lp *pb.Laptop) (*pb.Laptop, error) {
	other := &pb.Laptop{}
	if err := copier.Copy(other, lp); err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}
	return other, nil
}
