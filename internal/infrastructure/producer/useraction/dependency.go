package useraction

import (
	"context"
	"github/user-manager/pkg/event"
)

type eventProducer interface {
	Produce(ctx context.Context, event event.Event) error
}
