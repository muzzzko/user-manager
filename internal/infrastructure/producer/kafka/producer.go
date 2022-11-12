package kafka

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"

	"github/user-manager/pkg/event"
)

type Producer struct {
	saramaProducer sarama.AsyncProducer
}

func NewProducer(saramaProducer sarama.AsyncProducer) *Producer {
	return &Producer{
		saramaProducer: saramaProducer,
	}
}

func (p *Producer) Produce(ctx context.Context, e event.Event) error {
	if err := e.Validate(); err != nil {
		return err
	}

	select {
	case p.saramaProducer.Input() <- &sarama.ProducerMessage{
		Topic: e.GetTopic(),
		Key:   e.GetKey(),
		Value: e.GetValue(),
	}:
		return nil
	case err := <-p.saramaProducer.Errors():
		return fmt.Errorf("failed to produce message: %w", err)
	case <-ctx.Done():
		return ctx.Err()
	}
}
