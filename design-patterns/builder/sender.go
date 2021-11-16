package builder

type Sender struct {
	mb MessageBuilder
}

func (s *Sender) SetBuilder(mb MessageBuilder) { s.mb = mb }

// BuildMessage takes in a recipient and text to construct a represented
// message. If no previous builder was set it defaults to JSON.
func (s *Sender) BuildMessage(recpt, txt string) (*MessageRepresented, error) {
	if s.mb == nil {
		s.mb = &JSONMessageBuilder{}
	}
	return s.mb.SetRecipient(recpt).SetText(txt).Build()
}
