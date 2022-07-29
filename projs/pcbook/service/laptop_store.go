package service

import (
	"errors"
	"fmt"
	"grpbook/pb"
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
	other := &pb.Laptop{}
	if err := copier.Copy(other, lp); err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
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
	other := &pb.Laptop{}
	if err := copier.Copy(other, lp); err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}
	return other, nil
}
