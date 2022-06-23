package main

import (
	"fmt"
	"proto"
)

func main() {
	clr := "orange"
	pp := map[string]string{"home": "(123)859-0390", "work": "(342)238-8943"}
	p1 := proto.Prototype{Name: "Jay", Age: 929, Friends: []string{"K", "I"}, Color: &clr, Phones: pp}
	pcopy := p1
	pcopy.Age = 743
	pcopy.Name = "Joy"
	// XXX: Does not work because we have a slice of **TYPE** and that's just a
	// pointer under the hood, so when we change a value in that slice it changes
	// for all values associated to that pointer.
	// pcopy.Friends[0] = "Y"
	// pcopy.Friends[1] = "X"
	// pcopy.Color = "Green"
	// pcopy.Phones["home"] = "(899)283-9999"

  pclone := p1.Clone()
	pcopy.Age = 4444
	pcopy.Name = "Cloy"
	pclone.Friends[0] = "Y"
	pclone.Friends[1] = "X"
  c := "green"
	pclone.Color = &c
	pclone.Phones["home"] = "(899)283-9999"

  // We can still affect p1 pointers as well.
  p1.Friends = append(p1.Friends, "H")
  clr = "Gray"
  pp["work"] = "(777)889-9000"

	fmt.Println(p1.String())
	fmt.Println(pclone.String())
}
