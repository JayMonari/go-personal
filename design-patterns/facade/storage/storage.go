package storage

import "fmt"

type Storage struct{ engine string }

func New(e string) Storage { return Storage{engine: e} }

func (s *Storage) Save(comment string) {
	fmt.Printf("registered in the DB: %s\n", comment)
}
