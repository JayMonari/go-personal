package proxy

import "time"

type Data struct {
	books    Books
	server   string
	port     uint16
	user     string
	password string
}

func New(server string, port uint16, user, password string) *Data {
	d := &Data{
		server:   server,
		port:     port,
		user:     user,
		password: password,
	}
	d.load()
	return d
}

func (d *Data) ByID(ID uint) Book {
	time.Sleep(2 * time.Second)
	for _, v := range d.books {
		if v.ID == ID {
			return v
		}
	}

	return Book{}
}

func (d *Data) All() Books {
	time.Sleep(4 * time.Second)
	return d.books
}

func (d *Data) load() {
	d.books = make(Books, 0, 5)
	d.books = append(
		d.books,
		Book{ID: 1, Name: "Title 1", Author: "Author 1"},
		Book{ID: 2, Name: "Title 2", Author: "Author 2"},
		Book{ID: 3, Name: "Title 3", Author: "Author 3"},
		Book{ID: 4, Name: "Title 4", Author: "Author 4"},
		Book{ID: 5, Name: "Title 5", Author: "Author 5"},
	)
}
