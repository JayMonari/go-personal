package circular

import "errors"

type Buffer struct{ buf chan byte }

func NewBuffer(size int) *Buffer { return &Buffer{buf: make(chan byte, size)} }

func (b *Buffer) ReadByte() (byte, error) {
	if len(b.buf) == 0 {
		return 0, errors.New("buffer is empty")
	}
	return <-b.buf, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if len(b.buf) == cap(b.buf) {
		return errors.New("buffer has reached capacity")
	}
	b.buf <- c
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if err := b.WriteByte(c); err != nil {
		<-b.buf
		b.buf <- c
	}
}

func (b *Buffer) Reset() { b.buf = make(chan byte, cap(b.buf)) }
