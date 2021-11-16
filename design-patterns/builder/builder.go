package builder

type MessageBuilder interface {
	SetRecipient(r string) MessageBuilder
	SetText(t string) MessageBuilder
	Build() (*MessageRepresented, error)
}
