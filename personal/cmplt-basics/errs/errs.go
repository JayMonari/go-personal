package errs

import (
	"errors"
	"fmt"
	"math"
)

// New returns an error based on the way we want it to be made. Which can
// be done with the standard `errors` package or for more formatting options
// the `fmt` package. If the way is not recognized nil is returned.
func New(way string) error {
	switch way {
	case "fmt":
		return fmt.Errorf("we can use fmt to have formatting verbs: %q", way)
	case "errors":
		return errors.New("an error has occurred")
	default:
		return nil
	}
}

// realError is a living, breathing, 100% real error. It is important to
// understand that in Go -- errors are values. If the type implements the error
// interface which has 1 method -- Error() string -- it is an error.
type realError bool

// Error is the only method in the builtin error interface. It returns a
// message of what went wrong.
func (e realError) Error() string {
	return "this is a real error that can be returned if something goes wrong"
}

// Custom shows that implementing the builtin error interface is very easy
// to do and can be used to return custom errors instead of the most common
// unexported `errorString` struct in the `errors` package.
func Custom() error {
	return realError(true)
}

// TooBigError is an exported error that will tell the caller if the number
// input is too big.
type TooBigError int64

func (e TooBigError) Error() string {
	return fmt.Sprintf("number too big: %d", e)
}

// phoneNumberError is an unexported error that informs the caller when a
// bad phone number was passed in.
type phoneNumberError string

func (e phoneNumberError) Error() string {
	// We need to explicitly convert e to a string here or else we'll get
	// NOTE(jay): arg e causes a recursive Error method call.
	return fmt.Sprint("phone number must have 10 digits: ", string(e))
}

// InvalidRuneError is an error that let's the caller know the input rune does
// not work with the function.
type InvalidRuneError rune

func (e InvalidRuneError) Error() string {
	// We need to explicitly convert e to a string here or else we'll get
	// NOTE(jay): arg e causes a recursive Error method call.
	return fmt.Sprintf("input rune is not a valid english letter: %q", string(e))
}

// bearer is a simple interface much like error. It is important to recognize
// errors are values, meaning bearer and error are no different from one
// another. Anything you'd expect a normal value to do, error can to.
type bearer interface {
	Bearer() string
}

// UndeadWarrior is a bearer of a great curse and must travel to distant lands
// in hopes of finally removing it from themself.
type UndeadWarrior struct{}

// Bearer is a method like Error that takes no arguments and returns a string.
func (w UndeadWarrior) Bearer() string {
	return "Rise if you would. For that is our curse."
}

func (w UndeadWarrior) String() string { return w.Bearer() }

// ManyCustoms shows how to deal with many custom errors in a single
// function and shows that errors are just values that are returned by also
// returning a bearer which is very similar in behavior to an error.
func ManyCustoms(n uint32, phoneNo string, ltr rune) (bearer, error) {
	if n > uint32(math.Pow(2, 31)) {
		return nil, TooBigError(n)
	}
	nDigits := 0
	for _, r := range phoneNo {
		if r >= '0' && r <= '9' {
			nDigits++
		}
	}
	if nDigits != 10 {
		return nil, phoneNumberError(phoneNo)
	}
	if !(ltr >= 'A' && ltr <= 'Z' || ltr >= 'a' && ltr <= 'z') {
		return nil, InvalidRuneError(ltr)
	}
	return UndeadWarrior{}, nil
}

// ConnectionError extends the behavior of a basic error with more methods that
// could be useful for someone making a call. It can be used to check if
// someone missed the call and try again or if the number had been
// disconnected.
type ConnectionError interface {
	error
	Disconnect() bool // Is the person disconnected?
	Miss() bool       // Did they miss the contact?
}

// CallError implements a ConnectionError. We can imagine other Errors that
// implement ConnectionError like: TransreceiverError, MorseError,
// NetworkError, etc ...
type CallError struct{ Number string }

func (e CallError) Error() string {
	var reason string
	switch {
	case e.Disconnect():
		reason = "it has been disconnected"
	case e.Miss():
		reason = "no one picked up the phone"
	default:
		reason = "something went wrong, please try again"
	}
	return fmt.Sprintln("the number you dialed could not be reached:", reason)
}

// Disconnect satisfies part of ConnectionError.
func (e CallError) Disconnect() bool {
	if e.Number[:3] == "555" {
		return true
	}
	return false
}

// Miss satisfies part of ConnectionError.
func (e CallError) Miss() bool {
	if e.Number[0] == '7' {
		return true
	}
	return false
}

// ExtendBasic shows how to extend the simple error interface to have more
// functionality using composition and embedding the error interface into our
// new ConnectionError.
func ExtendBasic(phoneNo string) ConnectionError {
	return CallError{Number: phoneNo}
}

// WrapOtherErrors shows how to put an error inside of another error. This
// is very helpful when you have many moving parts in your application. We want
// to know **where** the error originated and what places it went along the
// way.
func WrapOtherErrors() error {
	if err := pkgBufioCall(); err != nil {
		return pkgHTTPCall(pkgJSONCall(pkgZipCall(err)))
	}
	return nil
}

func pkgHTTPCall(e error) error {
	return fmt.Errorf("http: Server closed: %w", e)
}

func pkgJSONCall(e error) error {
	return fmt.Errorf("json: syntax error, unexpected ',': %w", e)
}

func pkgZipCall(e error) error {
	return fmt.Errorf("zip: not a valid zip file: %w", e)
}

func pkgBufioCall() error {
	return errors.New("bufio.Scanner: token too long")
}

// NotNil shows that a nil error value does not equal nil. In other words
// setting an error to nil and returning that error will not give you nil. It
// shows the idiomatic Go way of returning nothing if there is no error.
func NotNil(doItWrong bool) error {
	// var incorrect *CallError = nil 👈 Same as below, but this is wrong because
	// nil is the zero value, but just to show you can do the above as well.
	var incorrect *CallError
	if doItWrong {
		return incorrect
	}
	return nil
}
