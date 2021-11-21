package secret

// events are the components of the secret handshake.
var events = [4]string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

// Handshake produces a slice of strings corresponding to the events and the
// given code.
func Handshake(code uint) []string {
	seq := []string{}
	for i, e := range events {
		secret := uint(1 << i)
		if secret&code != 0 {
			seq = append(seq, e)
		}
	}

	if 16&code == 0 {
		return seq
	}
	return reverse(seq)
}

// reverse takes a slice of string and reverses the elements order.
func reverse(sl []string) []string {
	s := make([]string, len(sl))
	for i := 0; i < len(sl); i++ {
		s[i] = sl[len(sl)-i-1]
	}
	return s
}
