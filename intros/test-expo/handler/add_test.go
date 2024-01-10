package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"

	"testexpo/handler"
)

func ExampleAdd() {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`numberOne=4&numberTwo=10`))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	handler.Add(rr, r)
	fmt.Print(rr.Body.String())
	// Output: 14
}

func TestAdd(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("numberOne=10&numberTwo=20"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler.Add(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if expected := "30"; rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// This shows how go has built in protection for data races that may occur in your program. This
// includes things like accessing something that's being modified and lock contention.
// Run:
//
//	go test -race ./...
func TestAddRaceCondition(t *testing.T) {
	t.Parallel()
	for name, tc := range map[string]struct {
		num1 string
		num2 string
		want string
	}{
		"1 + 2 = 3": {
			num1: "1",
			num2: "2",
			want: "3",
		},
		"2 + 1 = 3": {
			num1: "2",
			num2: "1",
			want: "3",
		},
		"10 + 20 = 30": {
			num1: "10",
			num2: "20",
			want: "30",
		},
		"0 + 0 = 0": {
			num1: "0",
			num2: "0",
			want: "0",
		},
	} {
		tc := tc // NOTE(jay): Will no longer need after go 1.22
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodPost, "/",
				strings.NewReader("numberOne="+tc.num1+"&numberTwo="+tc.num2),
			)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			rr := httptest.NewRecorder()
			handler.AddRaceCondition(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			if want, got := tc.want, rr.Body.String(); want != got {
				t.Errorf("handler returned unexpected body: want %v got %v", want, got)
			}
		})
	}
}

// https://eli.thegreenplace.net/2023/common-pitfalls-in-go-benchmarking/
// Run:
//
//	go test -bench=. -benchmem ./...
func BenchmarkAdd(b *testing.B) {
	// NOTE(jay): Not needed here, but doesn't hurt and required for simpler functions
	var rr *httptest.ResponseRecorder
	var r *http.Request

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rr = httptest.NewRecorder()
		r = httptest.NewRequest(http.MethodPost, "/",
			strings.NewReader(`numberOne=18446744073709551615&numberTwo=18446744073709551615`),
		)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		b.StartTimer()

		handler.Add(rr, r)
	}

	// NOTE(jay): Not needed here, but doesn't hurt and required for simpler functions
	runtime.KeepAlive(rr)
	runtime.KeepAlive(r)
}

// https://go.dev/doc/security/fuzz/
// Run:
//
//	go test -fuzz=FuzzAdd ./handler/
func FuzzAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, one, two string) {
		rr := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/",
			strings.NewReader(fmt.Sprintf(`numberOne=%s&numberTwo=%s`, one, two)),
		)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		handler.Add(rr, r)
		if sc := rr.Result().StatusCode; sc != http.StatusOK {
			t.Fatalf("want: %d, got: %d\n", http.StatusOK, sc)
		}
	})
}
