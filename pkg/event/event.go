package event

import "github.com/Shopify/sarama"

type Event interface {
	GetKey() sarama.StringEncoder
	GetValue() sarama.ByteEncoder
	GetTopic() string
	Validate() error
}
