package proxy

type Local struct {
	Remote *Data
	cache  Books
}

func NewLocal() Proxy {
	return &Local{
		Remote: New("postgresql://served.com", 9999, "usr", "pass"),
		cache:  make(Books, 0, 5),
	}
}

func (l *Local) GetByID(ID uint) Book {
	for _, b := range l.cache {
		if b.ID == ID {
			return b
		}
	}
  b := l.Remote.ByID(ID)
  l.cache = append(l.cache, b)
	return b
}

func (l *Local) GetAll() Books {
	return l.Remote.All()
}
