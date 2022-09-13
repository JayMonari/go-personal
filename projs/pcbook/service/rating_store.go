package service

import "sync"

// RatingStore is an interface to store laptop ratings.
type RatingStore interface {
	// Add adds a new laptop score to the store and returns its rating.
	Add(laptopID string, score float64) (*Rating, error)
}

// Rating contains the rating information of a laptop.
type Rating struct {
	Count uint32
	Sum   float64
}

// InMemoryRatingStore stores laptop ratings in memory.
type InMemoryRatingStore struct {
	mu      sync.RWMutex
	ratings map[string]*Rating
}

func NewInMemoryRatingStore() *InMemoryRatingStore {
	return &InMemoryRatingStore{
		ratings: map[string]*Rating{},
	}
}

func (s *InMemoryRatingStore) Add(laptopID string, score float64) (*Rating, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, ok := s.ratings[laptopID]
	if ok {
		r.Count++
		r.Sum += score
	} else {
		r = &Rating{
			Count: 1,
			Sum:   score,
		}
	}
	s.ratings[laptopID] = r
	return r, nil
}
