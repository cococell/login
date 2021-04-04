package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	login "login/proto/login"
)

type Login struct{}

func (e *Login) Handle(ctx context.Context, msg *login.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *login.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
