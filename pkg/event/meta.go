package event

import "time"

type Meta struct {
	CreatedAt time.Time `json:"created_at" validate:"required"`
}
