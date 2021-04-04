package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"log"
	proto "login/proto/login"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.login.client"),
		micro.Version("latest"),
	)

	client := proto.NewLoginService("go.micro.service.login", service.Client())

	req := proto.Request{Name: "yangan"}
	rsp, err := client.Call(context.Background(), &req)
	if err != nil {
		log.Fatal("call failed, ", err)
	}
	fmt.Println("success! msg: ", rsp.Msg)
}
