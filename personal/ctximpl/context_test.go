package ctximpl_test

import (
	"ctximpl"
	"fmt"
	"math"
	"testing"
	"time"
)

func TestBackgroundNotTODO(t *testing.T) {
	todo := fmt.Sprint(ctximpl.TODO())
	bg := fmt.Sprint(ctximpl.Background())

	if todo == bg {
		t.Errorf("TODO and Background are equal: %q vs %q", todo, bg)
	}
}

func TestWithCancel(t *testing.T) {
	ctx, cancel := ctximpl.WithCancel(ctximpl.Background())
	if err := ctx.Err(); err != nil {
		t.Errorf("error should be nil first, got %v", err)
	}
	cancel()
	<-ctx.Done()
}

func TestWithCancelConcurrent(t *testing.T) {
	ctx, cancel := ctximpl.WithCancel(ctximpl.Background())
	time.AfterFunc(1*time.Second, cancel)
	if err := ctx.Err(); err != nil {
		t.Errorf("error should be nil first, got %v", err)
	}
	<-ctx.Done()
	if err := ctx.Err(); err != ctximpl.Canceled {
		t.Errorf("error should be canceled now, got %v", err)
	}
}

func TestWithCancelPropagation(t *testing.T) {
	ctxA, cancelA := ctximpl.WithCancel(ctximpl.Background())
	ctxB, cancelB := ctximpl.WithCancel(ctxA)
	defer cancelB()
	cancelA()

	select {
	case <-ctxB.Done():
	case <-time.After(1 * time.Second):
		t.Errorf("time out")
	}

	if err := ctxB.Err(); err != ctximpl.Canceled {
		t.Errorf("error should be canceled now, got %v", err)
	}
}

func TestWithDeadline(t *testing.T) {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := ctximpl.WithDeadline(ctximpl.Background(), deadline)

	if d, ok := ctx.Deadline(); !ok || d != deadline {
		t.Errorf("deadline: wanted %v, got %v", deadline, d)
	}

	then := time.Now()
	<-ctx.Done()
	if d := time.Since(then); math.Abs(d.Seconds()-2.0) > 0.1 {
		t.Errorf("should have been done after 2.0 seconds, took %v", d)
	}
	if err := ctx.Err(); err != ctximpl.DeadlineExceeded {
		t.Errorf("error should be DeadlineExceeded, got %v", err)
	}

	cancel()
	if err := ctx.Err(); err != ctximpl.DeadlineExceeded {
		t.Errorf("error should still be DeadlineExceeded, got %v", err)
	}
}

func TestWithTimeout(t *testing.T) {
	timeout := 2 * time.Second
	deadline := time.Now().Add(timeout)
	ctx, cancel := ctximpl.WithTimeout(ctximpl.Background(), timeout)

	if d, ok := ctx.Deadline(); !ok || d.Sub(deadline) > time.Millisecond {
		t.Errorf("deadline: wanted %v, got %v", deadline, d)
	}

	then := time.Now()
	<-ctx.Done()
	if d := time.Since(then); math.Abs(d.Seconds()-2.0) > 0.1 {
		t.Errorf("should have been done after 2.0 seconds, took %v", d)
	}
	if err := ctx.Err(); err != ctximpl.DeadlineExceeded {
		t.Errorf("error should be DeadlineExceeded, got %v", err)
	}
	cancel()
	if err := ctx.Err(); err != ctximpl.DeadlineExceeded {
		t.Errorf("error should still be DeadlineExceeded, got %v", err)
	}
}

func TestWithValue(t *testing.T) {
	tc := []struct {
		key, val, keyRet, valRet interface{}
		shouldPanic              bool
	}{
		{"a", "b", "a", "b", false},
		{"a", "b", "c", nil, false},
		{42, true, 42, true, false},
		{42, true, int64(42), nil, false},
		{nil, true, nil, nil, true},
		{[]int{1, 2, 3}, true, []int{1, 2, 3}, nil, true},
	}

	for _, c := range tc {
		fmt.Println("TODO", c)
	}
}
