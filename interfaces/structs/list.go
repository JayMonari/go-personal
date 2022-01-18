package structs

import "fmt"

type printer interface {
	Print()
}

type Stock []printer

func (s Stock) Print() {
	if len(s) == 0 {
		fmt.Println("Sorry. We're waiting for delivery.")
	}
	for _, v := range s {
		v.Print()
	}
}
