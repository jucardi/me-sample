package helloworld

import (
	"context"
	"github.com/jucardi/go-titan/logx"
	. "github.com/jucardi/ms-sample/api/helloworld"
	"github.com/jucardi/ms-sample/server/service/message"
)

const (
	// BeanApiDefault is the name of the bean implementation of helloworld.IApi
	BeanApiDefault = "helloworld-controller-default"
)

var (
	// To validate the interface implementation at compile time.
	_ IApi = (*controller)(nil)
)

type controller struct {
	ctx context.Context
}

func (c *controller) init() *controller {
	// Add any initialization of the component here
	return c
}

func (c *controller) WithCtx(ctx context.Context) IApi {
	return &controller{ctx: ctx}
}

func (c *controller) GetHelloWorld() (*Message, error) {
	logx.Debug("HelloWorld called")
	return &Message{Message: "Hello World! Welcome to a Go Microservice."}, nil
}

func (c *controller) Create(req *Message) error {
	logx.Debug("API create")
	return message.Service().WithCtx(c.context()).Create(req.Name, req.Message)
}

func (c *controller) Get(name string) (*Message, error) {
	logx.Debug("API get")
	return message.Service().WithCtx(c.context()).Get(name)
}

func (c *controller) Update(req *Message) error {
	logx.Debug("API update")
	return message.Service().WithCtx(c.context()).Update(req.Name, req.Message)
}

func (c *controller) context() context.Context {
	if c.ctx == nil {
		return context.Background()
	}
	return c.ctx
}
