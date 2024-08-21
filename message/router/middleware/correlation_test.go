package middleware_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/exasyvetde/watermill/message"
	"github.com/exasyvetde/watermill/message/router/middleware"
)

func TestCorrelationID(t *testing.T) {
	handlerErr := errors.New("foo")

	handler := middleware.CorrelationID(func(msg *message.Message) ([]*message.Message, error) {
		return message.Messages{message.NewMessage("2", nil)}, handlerErr
	})

	msg := message.NewMessage("1", nil)
	middleware.SetCorrelationID("correlation_id", msg)

	producedMsgs, err := handler(msg)

	assert.Equal(t, "2", producedMsgs[0].UUID)
	assert.Equal(t, middleware.MessageCorrelationID(producedMsgs[0]), "correlation_id")
	assert.Equal(t, handlerErr, err)
}

func TestSetCorrelationID_already_set(t *testing.T) {
	msg := message.NewMessage("", nil)

	middleware.SetCorrelationID("foo", msg)
	middleware.SetCorrelationID("bar", msg)

	assert.Equal(t, "foo", middleware.MessageCorrelationID(msg))
}
