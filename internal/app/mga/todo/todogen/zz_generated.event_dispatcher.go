// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package todogen

import (
	"context"
	"emperror.dev/errors"
	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo"
)

// EventBus is a generic event bus.
type EventBus interface {
	// Publish sends an event to the underlying message bus.
	Publish(ctx context.Context, event interface{}) error
}

// EventDispatcher dispatches events through the underlying generic event bus.
type EventDispatcher struct {
	bus EventBus
}

// NewEventDispatcher returns a new EventDispatcher instance.
func NewEventDispatcher(bus EventBus) EventDispatcher {
	return EventDispatcher{bus: bus}
}

// MarkedAsDone dispatches a(n) MarkedAsDone event.
func (d EventDispatcher) MarkedAsDone(ctx context.Context, event todo.MarkedAsDone) error {
	err := d.bus.Publish(ctx, event)
	if err != nil {
		return errors.WithDetails(errors.WithMessage(err, "failed to dispatch event"), "event", "MarkedAsDone")
	}

	return nil
}
