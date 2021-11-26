package email

import "fmt"

type Email struct{}

func (e *Email) Notify(data string) { sendEmail(data) }

func sendEmail(data string) {
	fmt.Printf("An email was sent with data: %q\n", data)
}
