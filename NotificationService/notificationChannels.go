package main

import "fmt"

type NotificationChannel interface {
	Send(n *Notification) error
}

type EmailChannel struct{}

func (e *EmailChannel) Send(n *Notification) error {
	fmt.Println("Sending email to user: ", n)
	return nil
}

type SmsChannel struct{}

func (e *SmsChannel) Send(n *Notification) error {
	fmt.Println("Sending sms to user: ", n)
	return nil
}
