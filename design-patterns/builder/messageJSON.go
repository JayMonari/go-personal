package builder

import "encoding/json"

type JSONMessageBuilder struct {
	msg Message
}

func (b *JSONMessageBuilder) SetRecipient(r string) MessageBuilder {
	b.msg.Recipient = r
	return b
}

func (b *JSONMessageBuilder) SetText(t string) MessageBuilder {
	b.msg.Text = t
	return b
}

func (b *JSONMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := json.Marshal(b.msg)
	if err != nil {
		return nil, err
	}
	return &MessageRepresented{Body: data, Format: "JSON"}, nil
}
