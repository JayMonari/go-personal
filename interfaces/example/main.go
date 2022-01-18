package main

import "fmt"

type Walker interface {
	Walk() int
}

type DefaultWalker struct {
	energy int
}

func (dw *DefaultWalker) Walk() int {
	if dw.energy == 0 {
		dw.energy = 100
	}
	dw.energy -= 3
	return dw.energy
}

type Man interface {
	Walker
	Talk() string
}

type Manny struct {
	*DefaultWalker
}

func (m Manny) Talk() string {
	return "Hey, how's it hangin' pal?"
}

type Malint struct {
	*DefaultWalker
}

func (m Malint) Talk() string {
	return "Follow me, old chum."
}

func (m Malint) Grunt() string {
	return "RAGHRLURRRR!!!!"
}

type Bear interface {
	Walker
	Grunt() string
}

type PandaBear struct{}

func (p PandaBear) Name() string {
	return "Xiao Ming"
}

func (p PandaBear) Walk() (noop int) {
	return
}

func (p PandaBear) Grunt() string {
	return "*sniff* *sniff* urgh"
}

func (p PandaBear) SitAround() {
	fmt.Println("... I ain't doin' 💩")
}

type WereBear interface {
	Walker
	Man
	Bear
}

type Pig interface {
	Talk() string
}

type Police struct{}

func (p Police) Talk() string {
	return "STOP RESISTING!"
}

type ManBearPig interface {
	Man
	Bear
	Pig
}

type Human struct{}

type manBearPig struct {
	Malint
	PandaBear
	Police
}

func (mbp manBearPig) Talk() string {
	return "NOBODY EXPECTS THE SPANISH INQUISITION!"
}

func main() {
	m1 := Manny{&DefaultWalker{}}
	m2 := Malint{&DefaultWalker{}}
	// b := PandaBear{}
	converse(m1, m2)
	// stroll(m1, m2, b)
	// attack(m1, m2)
	// m := &manBearPig{}
	// m.Talk()
}

func attack(m1 Man, m2 WereBear) {
	var e1, e2 int
	for e1 >= 0 && e2 >= 0 {
		m1.Walk()
		m1.Walk()
		e1 = m1.Walk()
		e2 = m2.Walk()
		switch {
		case e1 <= 0:
			fmt.Println("Time to Feast!", m2.Grunt())
		case e2 <= 0:
			fmt.Println("The Prey lived to see another day", m1.Talk())
		}
	}
}

func stroll(m1, m2 Man, b Bear) {
	var e1, e2 int
	for i := 0; i < 10; i++ {
		e1 = m1.Walk()
		e2 = m2.Walk()
	}
	fmt.Println("Energy Level of Manny:", e1, "Energy Level of Malint2:", e2)
	fmt.Println(b.Grunt(), "Oh hey look, What a fat stupid bear.")
}

func converse(m1, m2 Man) {
	fmt.Println(m1.Talk())
	fmt.Println(m2.Talk())
}
