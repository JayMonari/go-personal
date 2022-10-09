package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing/quick"

	"golang.org/x/exp/constraints"
)

var rnd = rand.New(rand.NewSource(0))

func ExampleCheck() {
	// Output:
	//
}

func ExampleCheckEqual() {
	// Output:
	//
}

func ExampleValue() {
	// NOTE(jay): Placed in anonymous function so that function will continue
	// to print next example.
	func() {
		type causePanic struct {
			private bool
		}
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("a panic happened:", r)
			}
		}()
		quick.Value(reflect.TypeOf(causePanic{}), rnd)
	}()

	type kitchenSink struct {
		S    string
		Ui   uint
		Ui8  uint8
		Ui16 uint16
		Ui32 uint32
		Ui64 uint64
		I    int
		I8   int8
		I16  int16
		I32  int32
		I64  int64
		F32  float32
		F64  float64
		C64  complex64
		C128 complex128
		P    uintptr
	}
	v, _ := quick.Value(reflect.TypeOf(kitchenSink{}), rnd)
	fmt.Printf("\n%+v\n", v)
	// Output:
	// a panic happened: reflect: reflect.Value.Set using value obtained using unexported field
	//
	// {S:񦓩󋶒򑮿􋐍󽖁𯅴񗯉򞀒񆈔󙔆􊀐񓒪宎󼺶 Ui:16368296284793757383 Ui8:95 Ui16:21699 Ui32:2929938124 Ui64:3287288577352441706 I:-5292444157415598862 I8:109 I16:-22950 I32:-1400774486 I64:-2830929173373128928 F32:-1.3675876e+38 F64:-1.2285373489200797e+308 C64:(3.5394412e+37-5.148254e+37i) C128:(5.647763477622366e+307-1.1644320106541681e+308i) P:14225330248022666787}
}

func ExampleCheckEqualError() {
	// Output:
	//
}

func ExampleCheckError() {
	// Output:
	//
}

func ExampleConfig() {
	// Output:
	//
}

type CPU struct {
	Brand         string
	Name          string
	NumberCores   uint32
	NumberThreads uint32
	MinGhz        float64
	MaxGhz        float64
}

func rNum[N constraints.Integer](r *rand.Rand, min, max N) N {
	var randVal N
	switch any(min).(type) {
	case int64:
		randVal = N(rand.Int63())
	case int, int32:
		randVal = N(rand.Int31())
	case int16:
		randVal = N(rand.Intn(math.MaxInt16))
	case int8:
		randVal = N(rand.Intn(math.MaxInt8))
	case uint64:
		randVal = N(rand.Uint64())
	case uint, uint32:
		randVal = N(rand.Uint32())
	case uint16:
		randVal = N(rand.Intn(math.MaxUint16))
	case uint8:
		randVal = N(rand.Intn(math.MaxUint8))
	}
	return min + randVal%(max-min-1)
}

func rFloat[N constraints.Float](r *rand.Rand, min, max N) N {
	return min + N(rand.Float64())*(max-min)
}

var (
	brands   = [...]string{"Intel", "AMD"}
	amdNames = [...]string{
		"Ryzen 3 PRO 5350G",
		"Ryzen 7 PRO 5875U",
		"Ryzen 9 5900HS",
		"Ryzen Threadripper Pro 3945WX",
		"EPYC 7003",
	}
	intelNames = [...]string{
		"Pentium G7400",
		"Celeron G6900",
		"Core i9-9980HK",
		"Core i7-9750H",
		"Core i5-9400F",
		"Core i3-1005G1",
	}
)

func rCPUName(r *rand.Rand, brand string) string {
	if brand == brands[0] {
		return intelNames[r.Intn(len(intelNames))]
	}
	return amdNames[r.Intn(len(amdNames))]
}

func (c CPU) String() string {
	var sb strings.Builder
	sb.WriteString("Brand:\t" + c.Brand + "\n")
	sb.WriteString("Name:\t" + c.Name + "\n")
	sb.WriteString("Number of Cores:\t" + strconv.Itoa(int(c.NumberCores)) + "\n")
	sb.WriteString("Number of Threads:\t" + strconv.Itoa(int(c.NumberThreads)) + "\n")
	sb.WriteString("Minimum GHz:\t" + strconv.FormatFloat(c.MinGhz, 'f', 4, 64) + "\n")
	sb.WriteString("Max GHz:\t" + strconv.FormatFloat(c.MaxGhz, 'f', 4, 64) + "\n")
	return sb.String()
}

func (c CPU) Generate(r *rand.Rand, size int) reflect.Value {
	b := brands[r.Intn(len(brands))]
	cores := rNum[uint32](r, 2, 8)
	minGhz := rFloat(r, 2.0, 3.5)
	return reflect.ValueOf(CPU{
		Brand:         b,
		Name:          rCPUName(r, b),
		NumberCores:   cores,
		NumberThreads: rNum(r, cores, 12),
		MinGhz:        minGhz,
		MaxGhz:        rFloat(r, minGhz, 5.0),
	})
}

func ExampleGenerator() {
	var _ quick.Generator = CPU{}
	for i := 0; i < 5; i++ {
		v, _ := quick.Value(reflect.TypeOf(CPU{}), rnd)
		fmt.Println(v)
	}
	// Output:
	//
	// Brand:	Intel
	// Name:	Celeron G6900
	// Number of Cores:	4
	// Number of Threads:	6
	// Minimum GHz:	3.4108
	// Max GHz:	4.1064
	//
	// Brand:	Intel
	// Name:	Pentium G7400
	// Number of Cores:	4
	// Number of Threads:	10
	// Minimum GHz:	3.0302
	// Max GHz:	3.3385
	//
	// Brand:	Intel
	// Name:	Celeron G6900
	// Number of Cores:	4
	// Number of Threads:	6
	// Minimum GHz:	2.4514
	// Max GHz:	4.5250
	//
	// Brand:	AMD
	// Name:	Ryzen Threadripper Pro 3945WX
	// Number of Cores:	2
	// Number of Threads:	10
	// Minimum GHz:	2.5710
	// Max GHz:	3.7099
	//
	// Brand:	Intel
	// Name:	Pentium G7400
	// Number of Cores:	4
	// Number of Threads:	4
	// Minimum GHz:	2.4397
	// Max GHz:	2.9992
}

func ExampleSetupError() {
	// Output:
	//
}
