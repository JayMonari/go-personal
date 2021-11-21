package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

// Number sanitizes and returns a 10 digit phone number. If there are any
// invalid characters or incorrect values for area or exchange codes an error
// is returned.
func Number(phoneNo string) (string, error) {
	d := regexp.MustCompile("[^0-9]").ReplaceAllString(phoneNo, "")
	// Validate number length and remove country code if exists.
	if len(d) < 10 || len(d) > 11 {
		return "", errors.New("Invalid number.")
	} else if len(d) == 11 && d[0] != '1' {
		return "", errors.New("Invalid number.")
	} else if len(d) == 11 {
		d = d[1:]
	}

	// Validate area code and exchange code values.
	if int(d[0]-'0') < 2 {
		return "", errors.New("Invalid number.")
	} else if int(d[3]-'0') < 2 {
		return "", errors.New("Invalid number.")
	}

	return d, nil
}

// AreaCode returns the area code of a given phone number. See Number for error
// details.
func AreaCode(phoneNumber string) (string, error) {
	d, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return d[0:3], nil
}

// Format returns a phone number to the following format: (xxx) xxx-xxxx
// see Number for error details.
func Format(phoneNumber string) (string, error) {
	d, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", d[0:3], d[3:6], d[6:]), nil
}
