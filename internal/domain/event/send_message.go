package event

type SendMessage interface {
	Publish(topicArn string, message string) error
}
