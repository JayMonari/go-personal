package notify

import "fmt"

type Email struct{ to, msg string }

func New() Email { return Email{} }

func (e *Email) Send(to, comment string) {
	e.to = to
	e.msg = comment
	fmt.Printf("Sent email to: %s, message: %s\n", e.to, e.msg)
}
