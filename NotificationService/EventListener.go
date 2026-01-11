package main

type Event struct {
	message string
}

type EventListener interface {
	OnEvent(e *Event) error
}

type NotificationEventListener struct {
	notificationService *NotificationService
}

func NewNotificationListener(s *NotificationService) *NotificationEventListener {
	return &NotificationEventListener{
		notificationService: s,
	}
}

func (s *NotificationEventListener) OnEvent(e *Event) error {
	s.notificationService.Notify(e)
	return nil
}
