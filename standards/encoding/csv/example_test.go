package main

import (
	"bytes"
	"encoding/csv"
	"strings"
)

func Example_one() {
	srdr := csv.NewReader(strings.NewReader(`
id*firstname*lastname*email*email2*profession*birth_date
100*Marnia*Randene*Marnia.Randene@yopmail.com*Marnia.Randene@gmail.com*doctor*1948-08-04
101*Doro*Hunfredo*Doro.Hunfredo@yopmail.com*Doro.Hunfredo@gmail.com*firefighter*1963-11-14
102*Karlee*Persse*Karlee.Persse@yopmail.com*Karlee.Persse@gmail.com*worker*1924-10-16`[1:]))
	brdr := csv.NewReader(bytes.NewReader([]byte(`
id,firstname,lastname,email,email2,profession,birth_date
100,Marnia,Randene,Marnia.Randene@yopmail.com,Marnia.Randene@gmail.com,doctor,1948-08-04
101,Doro,Hunfredo,Doro.Hunfredo@yopmail.com,Doro.Hunfredo@gmail.com,firefighter,1963-11-14
102,Karlee,Persse,Karlee.Persse@yopmail.com,Karlee.Persse@gmail.com,worker,1924-10-16`[1:])))

	srdr.Comma = '*'
	srdr.TrimLeadingSpace = true

	brdr.LazyQuotes = true
	brdr.TrimLeadingSpace
}
