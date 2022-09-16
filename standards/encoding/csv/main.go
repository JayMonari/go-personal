package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("extendsclass.csv")
	if err != nil {
		log.Fatal(err)
	}
	r2, err := os.OpenFile("extendsclass.csv", os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	r.Comment = '#'
	// r.ReuseRecord = true
	r.TrimLeadingSpace = true

	offset := int64(0)
	var records [][]string
	var badRecs [][]string
	var errs []error
	for rec, err := r.Read(); !errors.Is(err, io.EOF); rec, err = r.Read() {
		if err != nil {
			switch {
			case errors.Is(err, csv.ErrBareQuote):
				b := make([]byte, r.InputOffset()-offset)
				fmt.Printf("r.InputOffset(): %v\n", r.InputOffset())
				fmt.Printf("offset: %v\n", offset)
				fmt.Printf("len(b): %v\n", len(b))
				r2.ReadAt(b, offset)

				_, col, _ := strings.Cut(err.Error(), " column ")
				col, _, _ = strings.Cut(col, ":")
				colNo, _ := strconv.Atoi(col)
				records = append(records, strings.Fields(string(append(b[:colNo-1], b[colNo:len(b)-1]...))))
			case errors.Is(err, csv.ErrQuote):
				b := make([]byte, r.InputOffset()-offset)
				fmt.Printf("r.InputOffset(): %v\n", r.InputOffset())
				fmt.Printf("offset: %v\n", offset)
				fmt.Printf("len(b): %v\n", len(b))
				r2.ReadAt(b, offset)
				_, col, _ := strings.Cut(err.Error(), " column ")
				col, _, _ = strings.Cut(col, ":")
				colNo, _ := strconv.Atoi(col)
				records = append(records,
					strings.Fields(string(append(b[:colNo-1], b[colNo:len(b)-1]...))))
				fmt.Printf("b: %s\n", b)
			case errors.Is(err, csv.ErrFieldCount):
				// FIXME(sd): Send to errs
				errs = append(errs, err)
			}
			badRecs = append(badRecs, rec)
			offset = r.InputOffset()
			continue
		}
		records = append(records, rec)
		offset = r.InputOffset()
	}
	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records)
	// fmt.Printf("records: %v\n", records[:10])
	fmt.Printf("errs: %v\n", errs)
	fmt.Printf("badRecs: %v\n", badRecs)
}

// https://github.com/jszwec/csvutil/blob/2b7b86bab664a79d0e2e5c3d285c195df32aadc9/decoder.go#L439
