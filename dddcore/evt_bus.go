package dddcore

import "context"

type EventHandler func(ctx context.Context, event Event)

// EventBus allow to publis/subscribe to events
type EventBus interface {
	Publish(ctx context.Context, eventType string, event Event)
	Subscribe(eventType string, fn EventHandler) error
	Unsubscribe(eventType string, fn EventHandler) error
}
