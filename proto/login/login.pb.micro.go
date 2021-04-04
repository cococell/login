// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/login/login.proto

package login

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Login service

func NewLoginEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Login service

type LoginService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Login_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Login_PingPongService, error)
}

type loginService struct {
	c    client.Client
	name string
}

func NewLoginService(name string, c client.Client) LoginService {
	return &loginService{
		c:    c,
		name: name,
	}
}

func (c *loginService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Login.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Login_StreamService, error) {
	req := c.c.NewRequest(c.name, "Login.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &loginServiceStream{stream}, nil
}

type Login_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type loginServiceStream struct {
	stream client.Stream
}

func (x *loginServiceStream) Close() error {
	return x.stream.Close()
}

func (x *loginServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *loginServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *loginServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *loginServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *loginService) PingPong(ctx context.Context, opts ...client.CallOption) (Login_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Login.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &loginServicePingPong{stream}, nil
}

type Login_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type loginServicePingPong struct {
	stream client.Stream
}

func (x *loginServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *loginServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *loginServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *loginServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *loginServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *loginServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Login service

type LoginHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Login_StreamStream) error
	PingPong(context.Context, Login_PingPongStream) error
}

func RegisterLoginHandler(s server.Server, hdlr LoginHandler, opts ...server.HandlerOption) error {
	type login interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Login struct {
		login
	}
	h := &loginHandler{hdlr}
	return s.Handle(s.NewHandler(&Login{h}, opts...))
}

type loginHandler struct {
	LoginHandler
}

func (h *loginHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.LoginHandler.Call(ctx, in, out)
}

func (h *loginHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.LoginHandler.Stream(ctx, m, &loginStreamStream{stream})
}

type Login_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type loginStreamStream struct {
	stream server.Stream
}

func (x *loginStreamStream) Close() error {
	return x.stream.Close()
}

func (x *loginStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *loginStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *loginStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *loginStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *loginHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.LoginHandler.PingPong(ctx, &loginPingPongStream{stream})
}

type Login_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type loginPingPongStream struct {
	stream server.Stream
}

func (x *loginPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *loginPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *loginPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *loginPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *loginPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *loginPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}