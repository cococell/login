package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	login "login/proto/login"
)

type Login struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Login) Call(ctx context.Context, req *login.Request, rsp *login.Response) error {
	log.Info("Received Login.Call request", req.Name)
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Login) Stream(ctx context.Context, req *login.StreamingRequest, stream login.Login_StreamStream) error {
	log.Infof("Received Login.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&login.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Login) PingPong(ctx context.Context, stream login.Login_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&login.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
