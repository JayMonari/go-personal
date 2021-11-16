package builder

import "encoding/xml"

type XMLMessageBuilder struct {
	msg Message
}

func (b *XMLMessageBuilder) SetRecipient(r string) MessageBuilder {
	b.msg.Recipient = r
	return b
}

func (b *XMLMessageBuilder) SetText(t string) MessageBuilder {
	b.msg.Text = t
	return b
}

func (b *XMLMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := xml.Marshal(b.msg)
	if err != nil {
		return nil, err
	}
	return &MessageRepresented{Body: data, Format: "XML"}, nil
}
