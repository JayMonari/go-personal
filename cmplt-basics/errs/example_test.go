package errs_test

import (
	"basics/errs"
	"errors"
	"fmt"
	"math"
)

func ExampleNew() {
	fmt.Printf(
		"%v\n%#v\n%#v\n%+v\n",
		errs.New("fmt"),
		errs.New("fmt"),
		errs.New("errors"),
		errs.New("🤷"),
	)
	// Output:
	// we can use fmt to have formatting verbs: "fmt"
	// &errors.errorString{s:"we can use fmt to have formatting verbs: \"fmt\""}
	// &errors.errorString{s:"an error has occurred"}
	// <nil>
}

func ExampleCustom() {
	if err := errs.Custom(); err != nil {
		fmt.Println(err)
	}
	// Output: this is a real error that can be returned if something goes wrong
}

func ExampleManyCustoms() {
	if _, err := errs.ManyCustoms(
		uint32(math.Pow(2, 32)-1), "(555)867-5309", 'A'); err != nil {
		fmt.Println(err)
	}
	if _, err := errs.ManyCustoms(0xff, "(555)67-5309", 'z'); err != nil {
		fmt.Println(err)
	}
	if _, err := errs.ManyCustoms(0b1, "(555)867-5309", '🤪'); err != nil {
		fmt.Println(err)
	}
	bearer, err := errs.ManyCustoms(0o7, "(555)867-5309", 'G')
	if err != nil {
		panic(err)
	}
	fmt.Println(bearer)
	// Output:
	// number too big: 4294967295
	// phone number must have 10 digits: (555)67-5309
	// input rune is not a valid english letter: "🤪"
	// Rise if you would. For that is our curse.
}

func ExampleExtendBasic() {
	if err := errs.ExtendBasic("555-212-4958").(errs.ConnectionError); err != nil {
		fmt.Printf("%#v\n%s\n", err, err)
	}
	if err := errs.ExtendBasic("777-390-9911").(errs.ConnectionError); err != nil {
		fmt.Printf("%#v\n%v\n", err, err)
		if err.Miss() {
			fmt.Println("Call again...")
		}
	}
	// Output:
	// errs.CallError{Number:"555-212-4958"}
	// the number you dialed could not be reached: it has been disconnected
	//
	// errs.CallError{Number:"777-390-9911"}
	// the number you dialed could not be reached: no one picked up the phone
	//
	// Call again...
}

func ExampleWrapOtherErrors() {
	if err := errs.WrapOtherErrors(); err != nil {
		fmt.Println("Wrapped error:", err)
		for err != nil {
			err = errors.Unwrap(err)
			fmt.Println("Unwrapping error:", err)
		}
	}
	// Output:
	// Wrapped error: http: Server closed: json: syntax error, unexpected ',': zip: not a valid zip file: bufio.Scanner: token too long
	// Unwrapping error: json: syntax error, unexpected ',': zip: not a valid zip file: bufio.Scanner: token too long
	// Unwrapping error: zip: not a valid zip file: bufio.Scanner: token too long
	// Unwrapping error: bufio.Scanner: token too long
	// Unwrapping error: <nil>
}

func ExampleNotNil() {
	if err := errs.NotNil(true); err != nil {
		fmt.Printf("YAH GOOFED! %#v", err)
	}

	if err := errs.NotNil(false); err != nil {
		fmt.Println("Never going to see this")
	}
	// Output:
	// YAH GOOFED! (*errs.CallError)(nil)
}
