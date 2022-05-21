package err_test

import (
	"basics/err"
	"errors"
	"fmt"
	"math"
)

func ExampleErrorNew() {
	fmt.Printf(
		"%v\n%#v\n%#v\n%+v\n",
		err.ErrorNew("fmt"),
		err.ErrorNew("fmt"),
		err.ErrorNew("errors"),
		err.ErrorNew("🤷"),
	)
	// Output:
	// we can use fmt to have formatting verbs: "fmt"
	// &errors.errorString{s:"we can use fmt to have formatting verbs: \"fmt\""}
	// &errors.errorString{s:"an error has occured"}
	// <nil>
}

func ExampleErrorCustom() {
	if err := err.ErrorCustom(); err != nil {
		fmt.Println(err)
	}
	// Output: this is a real error that can be returned if something goes wrong
}

func ExampleErrorManyCustoms() {
	if _, err := err.ErrorManyCustoms(
		uint32(math.Pow(2, 32)-1), "(555)867-5309", 'A'); err != nil {
		fmt.Println(err)
	}
	if _, err := err.ErrorManyCustoms(0xff, "(555)67-5309", 'z'); err != nil {
		fmt.Println(err)
	}
	if _, err := err.ErrorManyCustoms(0b1, "(555)867-5309", '🤪'); err != nil {
		fmt.Println(err)
	}
	bearer, err := err.ErrorManyCustoms(0o7, "(555)867-5309", 'G')
	if err != nil {
		panic(err)
	}
	fmt.Println(bearer.Bearer())
	// Output:
	// number too big: 4294967295
	// phone number must have 10 digits: (555)67-5309
	// input rune is not a valid english letter: "🤪"
	// Rise if you would. For that is our curse.
}

func ExampleErrorExtendBasic() {
	if err := err.ErrorExtendBasic("555-212-4958").(err.ConnectionError); err != nil {
		fmt.Printf("%#v\n%s\n", err, err)
	}
	if err := err.ErrorExtendBasic("777-390-9911").(err.ConnectionError); err != nil {
		fmt.Printf("%#v\n%v\n", err, err)
		if err.Miss() {
			fmt.Println("Call again...")
		}
	}
	// Output:
	// err.CallError{Number:"555-212-4958"}
	// the number you dialed could not be reached: it has been disconnected
	//
	// err.CallError{Number:"777-390-9911"}
	// the number you dialed could not be reached: no one picked up the phone
	//
	// Call again...
}

func ExampleErrorWrapOtherErrors() {
	if err := err.ErrorWrapOtherErrors(); err != nil {
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

func ExampleErrorNotNil() {
	if err := err.ErrorNotNil(true); err != nil {
		fmt.Printf("YAH GOOFED! %#v", err)
	}

	if err := err.ErrorNotNil(false); err != nil {
		fmt.Println("Never going to see this")
	}
	// Output:
	// YAH GOOFED! (*err.CallError)(nil)
}
