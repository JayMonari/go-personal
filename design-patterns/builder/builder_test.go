package builder_test

import (
	"builder"
	"testing"
)

func TestSender_BuildMessageXML(t *testing.T) {
	x := &builder.XMLMessageBuilder{}
	s := &builder.Sender{}

	s.SetBuilder(x)
	msg, err := s.BuildMessage("Gopher", "Hola mundo!")
	if err != nil {
		t.Fatalf("BuildMessage(): received error when shouldn't have: %v", err)
	}

	if `<Message><recipient>Gopher</recipient><text>Hola mundo!</text></Message>` != string(msg.Body) {
		t.Fail()
	}

	t.Log(string(msg.Body))
}

func TestSender_BuildMessageJSON(t *testing.T) {
	j := &builder.JSONMessageBuilder{}
	s := &builder.Sender{}

	s.SetBuilder(j)
	msg, err := s.BuildMessage("Gopher", "Hola mundo!")
	if err != nil {
		t.Fatalf("BuildMessage(): received error when shouldn't have: %v", err)
	}

	if `{"recipient":"Gopher","text":"Hola mundo!"}` != string(msg.Body) {
		t.Fail()
	}
	t.Log(string(msg.Body))
}
