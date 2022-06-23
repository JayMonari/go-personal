package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	sweaters := Inventory{Material: "wool", Count: 17}
	tmpl, err := template.New("example").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		log.Fatal(err)
	}
}

type Inventory struct {
	Material string
	Count    uint
}
