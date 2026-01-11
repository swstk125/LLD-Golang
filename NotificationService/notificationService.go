package main

type Notification struct {
	message string
}

type NotificationService struct {
	userService *UserService
}

func NewNotificationService(u *UserService) *NotificationService {
	return &NotificationService{
		userService: u,
	}
}

func (s *NotificationService) Notify(e *Event) {
	newNotif := &Notification{
		message: e.message,
	}

	for _, u := range s.userService.GetAll() {
		preferences := u.notificationPreferences

		for _, pref := range preferences {
			channel := GetChannel(pref)
			channel.Send(newNotif)
		}
	}
}
