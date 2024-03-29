package sns

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNS struct {
	svc *sns.SNS
}

func NewSNS(sess *session.Session) *SNS {

  return &SNS{
    svc: sns.New(sess),
  }
}

func (s *SNS) Publish(topicArn string, message string) error {
	fmt.Print(message)
	input := &sns.PublishInput{
    Message:  aws.String(message),
    TopicArn: aws.String(topicArn),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"ContentType": {
				DataType:    aws.String("String"),
				StringValue: aws.String("application/json"),
			},
		},
  }

  _, err := s.svc.Publish(input)
  if err!= nil {
    return fmt.Errorf("failed to publish message: %w", err)
  }

  return nil
}
