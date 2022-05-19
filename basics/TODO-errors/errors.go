package err

import (
	"errors"
	"fmt"
	"math"
)

// ErrorNew returns an error based on the way we want it to be made. Which can
// be done with the standard `errors` package or for more formatting options
// the `fmt` package. If the way is not recognized nil is returned.
func ErrorNew(way string) error {
	switch way {
	case "fmt":
		return fmt.Errorf("we can use fmt to have formatting verbs: %q", way)
	case "errors":
		return errors.New("an error has occured")
	default:
		return nil
	}
}

// realError is a living, breathing, 100% error. It is important to understand
// that in Go -- errors are values. If the type implements the error interface
// which has 1 method -- Error() string -- it is an error.
type realError bool

// Error is the only method in the builtin error interface. It returns a
// message of what went wrong.
func (e realError) Error() string {
	return "this is a real error that can be returned if something goes wrong"
}

// ErrorCustom shows that implementing the builtin error interface is very easy
// to do and can be used to return custom errors instead of the most common
// unexported `errorString` struct in the `errors` package.
func ErrorCustom() error {
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
	// XXX: arg e causes a recursive Error method call.
	return fmt.Sprintln("phone number must have 10 digits:", string(e))
}

// InvalidRuneError is an error that let's the caller know the input rune does
// not work with the function.
type InvalidRuneError rune

func (e InvalidRuneError) Error() string {
	// We need to explicitly convert e to a string here or else we'll get
	// XXX: arg e causes a recursive Error method call.
	return fmt.Sprintf("input rune is not a valid english letter: %q", string(e))
}

type bearer interface {
	Bearer() string
}

type UndeadWarrior struct{}

func (w UndeadWarrior) Bearer() string {
	return "Rise if you would. For that is our curse."
}

// ErrorLotsOfCustoms shows how to deal with many custom errors in a single
// function and shows that errors are just values that are returned by also
// returning a bearer which is very similar in behavior to an error.
func ErrorLotsOfCustoms(n uint32, phoneNo string, ltr rune) (bearer, error) {
	if n > uint32(math.Pow(2, 31)) {
		return nil, TooBigError(n)
	}
	nDigits := 0
	for _, r := range phoneNo {
		if r >= 'a' && r <= 'z' ||
			r >= 'A' && r <= 'Z' {
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

type PhoneError interface {
	error
	Disconnect() bool // Is the number disconnected?
	Miss() bool       // Did they miss the call?
}

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

func (e CallError) Disconnect() bool {
	if e.Number[:3] == "555" {
		return true
	}
	return false
}

func (e CallError) Miss() bool {
	if e.Number[0] == '7' {
		return true
	}
	return false
}

// ErrorExtendBasic shows how to extend the simple error interface to have more
// functionality using composition and embedding the error interface into our
// new PhoneError.
func ErrorExtendBasic(phoneNo string) PhoneError {
	return CallError{Number: phoneNo}
}

// ErrorNotNil shows that a nil error value does not equal nil. In other words
// setting an error to nil and returning that error will not give you nil. It
// shows the idomatic Go way of returning nothing if there is no error.
func ErrorNotNil(doItWrong bool) error {
	var incorrect *CallError = nil
	if doItWrong {
		return incorrect
	}
	return nil
}
