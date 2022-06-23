package defers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// RunAtEnd shows that using the `defer` keyword will make that function
// run at the very end of our function.
func RunAtEnd() {
	defer fmt.Println("defer: this is the first line, but the last output.")
	fmt.Println("This is the second line, but first to be printed.")
	fmt.Println("A couple of more lines for good measure.")
	fmt.Println("Never hurt anyone.")
}

// LIFO shows that if using multiple defers they work in a Last In First
// Out (LIFO) fashion.
func LIFO() {
	defer fmt.Println("First In. Third Out")
	fmt.Println("Random line 1")
	fmt.Println("Random line 2")
	defer fmt.Println("Second In. Second Out")
	fmt.Println("Random line 3")
	defer fmt.Println("Third In. First Out")
}

// ArgumentsEvaluated shows the arguments passed to a `defer`'s function
// are immediately evaluated and not evaluated at the end of the function's
// lifetime
func ArgumentsEvaluated() {
	whatWillIBe := 42
	defer fmt.Println("Number in defer:", whatWillIBe)
	whatWillIBe = 9001
	fmt.Println("Number at end of function:", whatWillIBe)
}

// NamedReturn shows that before we return a named value, if we change it
// in a deferred function, it will change it before returning the value.
// Compare with Return.
func NamedReturn() (gi string) {
	gi = "Golden Idol"
	defer indianaJones(&gi)
	return gi
}

// Return shows that depending on the scope, in the function or out of the
// function, the return value will change. Compare with NamedReturn.
func Return() string {
	gi := "Golden Idol"
	defer indianaJones(&gi)
	return gi
}

func indianaJones(s *string) { *s = "Bag of sand" }

// RecoverPanic shows how to recover from a panic by using a `defer`
// statement and the `recover` keyword that will be covered in the next lesson.
func RecoverPanic(calmdown bool) {
	if calmdown {
		defer recuperate()
	}
	panics()
	fmt.Println("Looks like we recovered in time.")
}

func panics() {
	panic("WE'RE ALL GOING DOWN, THIS IS THE END!!!")
}

func recuperate() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from:", r)
	}
}

// FileClose shows that we **always close** file descriptors after we open
// them.
func FileClose() {
	f, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE, 0744)
	if err != nil {
		fmt.Println(err)
	}
	defer closeFile(f)

	_, err = f.WriteString("👋 🌏")
	if err != nil {
		fmt.Println(err)
	}
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	if err := f.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// TempFileRemoveClose shows that with a temp file (and temp directories)
// we should clean up after ourselves by removing the file and closing the file
// descriptor.... **Always remove and close** temp file descriptors.
func TempFileRemoveClose() {
	f, err := os.CreateTemp(".", "tempfile")
	if err != nil {
		fmt.Println(err)
	}
	defer removeFile(f)
	defer closeFile(f)
}

func removeFile(f *os.File) {
	fmt.Println("removing")
	if err := os.Remove(f.Name()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// HTTPBodyClose shows that we **always close** our `*http.Response` body
// after it successfully completes.
func HTTPBodyClose() {
	res, err := http.Get("https://gophergo.dev")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	buf := make([]byte, 4096)
	io.ReadFull(res.Body, buf)
	fmt.Printf("Does the first 4KiB of data have Gopher Go in it? %t\n",
		bytes.Contains(buf, []byte("Gopher Go")))
}

// CancelContext shows a contrived way of canceling a context. Though it
// is a toy example, it is still very important that you **always cancel** your
// context after you're done using it.
func CancelContext(d time.Duration) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()
	t := time.NewTimer(d)
	select {
	case <-ctx.Done():
		fmt.Println("You're too slow!!!")
	case <-t.C:
		fmt.Println("Don't forget to cancel!")
	}
}

// account acts as a bank account with a balance
type account struct {
	balance int64
	sync.Mutex
}

// NewAccount ...
func NewAccount(balance int64) *account { return &account{balance: balance} }

// Balance gets the amount of money stored in cents. It locks the account from
// other transactions, such as a deposit, so the statement is accurate.
func (a *account) Balance() int64 {
	a.Lock()
	defer a.Unlock()

	return a.balance
}

// Deposit places money into the account. It locks the account so that other
// transactions, such as Withdrawal, cannot effect the balance while it's being
// updated.
func (a *account) Deposit(amount int64) (int64, error) {
	if amount < 0 {
		return -1, errors.New("Cannot deposit a negative amount.")
	}
	a.Lock()
	defer a.Unlock()

	a.balance += amount
	return a.balance, nil
}
