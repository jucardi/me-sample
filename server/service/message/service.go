package message

import (
	"context"
	"{{.golang.module_path}}/{{.service_name}}/api/helloworld"
	"{{.golang.module_path}}/{{.service_name}}/server/repository/message"
)

var (
	// To ensure *IService implements IService on build
	_ IService = (*service)(nil)
)

type service struct {
	ctx context.Context
}

func (s *service) init() *service {
	return s
}

func (s *service) WithCtx(ctx context.Context) IService {
	return &service{ctx: ctx}
}

func (s *service) Create(name, msg string) error {
	ret := &message.MessageDbe{
		Name:    name,
		Message: msg,
	}

	return message.Repository().WithCtx(s.context()).Create(ret)
}

func (s *service) Get(name string) (*helloworld.Message, error) {
	obj, err := message.Repository().WithCtx(s.context()).First(name)

	if err != nil {
		return nil, err
	}
	return fromDbe(obj), nil
}

func (s *service) Update(name, msg string) error {
	m, err := s.Get(name)
	if err != nil {
		return err
	}
	m.Message = msg

	return message.Repository().WithCtx(s.context()).Update(toDbe(m))
}

func (s *service) Exists(name string) bool {
	obj, err := message.Repository().First(name)
	return obj != nil && err == nil
}

func (s *service) context() context.Context {
	if s.ctx == nil {
		return context.Background()
	}
	return s.ctx
}
