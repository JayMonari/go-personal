package message

import "observer"

type Message struct {
	observers map[string]observer.Observer
	Data      string
}

func (m *Message) AddObserver(name string, o observer.Observer) {
  if m.observers == nil {
    m.observers = make(map[string]observer.Observer)
  }
	m.observers[name] = o
}

func (m *Message) RemoveObserver(name string) {
	delete(m.observers, name)
}

func (m *Message) NotifyObservers() {
	for _, v := range m.observers {
		v.Notify(m.Data)
	}
}
