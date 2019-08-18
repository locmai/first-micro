package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	blog "github.com/locmai/first-micro/services/blog/proto/blog"
)

type Blog struct{}

func (e *Blog) Handle(ctx context.Context, msg *blog.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *blog.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
