package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"login/handler"
	"login/subscriber"

	login "login/proto/login"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.login"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	login.RegisterLoginHandler(service.Server(), new(handler.Login))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.login", service.Server(), new(subscriber.Login))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
