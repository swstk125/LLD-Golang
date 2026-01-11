package main

type EventManager struct {
	Listners []EventListener
}

func NewEventManager() *EventManager {
	return &EventManager{
		Listners: make([]EventListener, 0),
	}
}

func (e *EventManager) Subscribe(l EventListener) {
	e.Listners = append(e.Listners, l)
}

func (e *EventManager) ProcessEvent(message string) error {
	newEvent := &Event{
		message: message,
	}
	for _, l := range e.Listners {
		l.OnEvent(newEvent)
	}
	return nil
}
