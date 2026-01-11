package main

func main() {
	userService := NewUserService()

	user1, user2 := userService.Create(), userService.Create()
	userService.AddPreference(user1, "EMAIL")
	userService.AddPreference(user1, "SMS")

	userService.AddPreference(user2, "SMS")

	notificationService := NewNotificationService(userService)
	notificationListener := NewNotificationListener(notificationService)

	eventMgr := NewEventManager()
	eventMgr.Subscribe(notificationListener)

	eventMgr.ProcessEvent("is it happening?")
}
