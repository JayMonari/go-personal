package main

import (
	"context"
	"io"
	"os"

	"prim/primitive"
)

func main() {
	f, err := os.Open("./input.jpg")
	if err != nil {
		panic(err)
	}
	out, err := primitive.Transform(context.TODO(), f, 10)
	if err != nil {
		panic(err)
	}
	outF, err := os.Create("output.jpg")
	if _, err = io.Copy(outF, out); err != nil {
		panic(err)
	}
}
