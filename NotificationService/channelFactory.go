package main

type ChannelType string

const (
	EMAIL ChannelType = "EMAIL"
	SMS   ChannelType = "SMS"
)

type ChannelFactory struct{}

func GetChannel(preferredChanel ChannelType) NotificationChannel {
	switch preferredChanel {
	case EMAIL:
		return &EmailChannel{}
	case SMS:
		return &SmsChannel{}
	default:
		return nil
	}
}
