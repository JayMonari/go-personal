package singleton

import "sync"

var (
	p    *person
	once sync.Once
)

type person struct {
	mu   sync.RWMutex
	name string
	age  int
}

func GetInstance() *person {
	once.Do(func() { p = &person{} })
	return p
}

func (p *person) GetName() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.name
}
func (p *person) SetName(n string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.name = n
}

func (p *person) GetAge() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.age
}

func (p *person) IncrementAge() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.age++
}
