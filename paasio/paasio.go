package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	r io.Reader
	c counter
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{
		r: r,
		c: counter{},
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	m, err := rc.r.Read(p)
	rc.c.addBytes(m)
	return m, err
}

// ReadCount returns the total number of bytes successfully read along with the
// total number of calls to Read()
func (rc *readCounter) ReadCount() (int64, int) {
	return rc.c.count()
}

type writeCounter struct {
	w io.Writer
	c counter
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{
		w: w,
		c: counter{},
	}
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.w.Write(p)
	wc.c.addBytes(n)
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	return wc.c.count()
}

type counter struct {
	bytes int64
	ops   int
	mu    sync.Mutex
}

func (c *counter) addBytes(n int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.bytes += int64(n)
	c.ops++
}

func (c *counter) count() (int64, int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.bytes, c.ops
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return readWriteCounter{
		NewReadCounter(rw),
		NewWriteCounter(rw),
	}
}
