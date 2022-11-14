package useraction

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github/user-manager/pkg/event"
)

type Producer struct {
	eventProducer eventProducer
}

func NewProducer(eventProducer eventProducer) *Producer {
	return &Producer{eventProducer: eventProducer}
}

func (p *Producer) Produce(
	ctx context.Context,
	userID strfmt.UUID,
	action string,
) error {
	ue := event.NewUserEvent()
	ue.Payload = event.UserActionEventPayload{
		UserID: userID.String(),
		Action: action,
	}

	return p.eventProducer.Produce(ctx, ue)
}
