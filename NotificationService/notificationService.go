package main

type EventListener interface {
	OnEvent(userId string, event string, message string) error
}

type Notification struct {
	UserId  string
	Event   string
	Message string
}

type NotificationService struct {
	userService *UserService
}

func NewNotificationService(u *UserService) *NotificationService {
	return &NotificationService{
		userService: u,
	}
}

func (s *NotificationService) OnEvent(userId, event, message string) {
	newNotif := &Notification{
		UserId:  userId,
		Event:   event,
		Message: message,
	}

	preferences := s.userService.Get(userId).notificationPreferences

	for _, pref := range preferences {
		channel := GetChannel(pref)
		channel.Send(newNotif)
	}
}
