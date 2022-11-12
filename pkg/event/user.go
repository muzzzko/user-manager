package event

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/go-playground/validator"
	"time"
)

const (
	CreateAction = "create"
	DeleteAction = "delete"
	UpdateAction = "update"

	userEventTopic = "user_manager.user.v1"
)

type UserActionEvent struct {
	Meta    Meta                   `json:"meta"`
	Payload UserActionEventPayload `json:"payload"`
	topic   string
}

type UserActionEventPayload struct {
	UserID string `json:"user_id" validate:"required,uuid4"`
	Action string `json:"action" validate:"required,oneof=create delete update"`
}

func NewUserEvent() *UserActionEvent {
	return &UserActionEvent{
		Meta: Meta{
			CreatedAt: time.Now().UTC(),
		},
		topic: userEventTopic,
	}
}

func (e *UserActionEvent) GetKey() sarama.StringEncoder {
	return sarama.StringEncoder(e.Payload.UserID)
}

func (e *UserActionEvent) GetValue() sarama.ByteEncoder {
	data, _ := json.Marshal(e)

	return data
}

func (e *UserActionEvent) GetTopic() string {
	return e.topic
}

func (e *UserActionEvent) Validate() error {
	validate := validator.New()

	//TODO check topic

	return validate.Struct(e)
}
