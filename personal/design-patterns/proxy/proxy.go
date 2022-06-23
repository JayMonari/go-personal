package proxy

type Proxy interface {
	GetByID(ID uint) Book
	GetAll() Books
}
