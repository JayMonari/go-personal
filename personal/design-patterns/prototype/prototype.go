package proto

import "fmt"

type Prototype struct {
	Name    string
	Age     int
	Friends []string
	Color   *string
	Phones  map[string]string
}

func (p *Prototype) String() string {
	return fmt.Sprintf(
		"Name: %s, Age: %d, Friends: %s Color: %s Phones: %s\n",
		p.Name, p.Age, p.Friends, *p.Color, p.Phones)
}

func (p *Prototype) Clone() Prototype {
	f := make([]string, len(p.Friends))
	copy(f, p.Friends)

	c := *p.Color

	pp := make(map[string]string, len(p.Phones))
	for k, v := range p.Phones {
		pp[k] = v
	}

	return Prototype{
		Name:    p.Name,
		Age:     p.Age,
		Friends: f,
		Color:   &c,
		Phones:  pp,
	}
}
