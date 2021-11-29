package message

import (
	"github.com/jucardi/ms-sample/api/helloworld"
	"github.com/jucardi/ms-sample/server/repository/message"
)

func toDbe(dto *helloworld.Message) *message.MessageDbe {
	return &message.MessageDbe{
		ID:      dto.ID,
		Name:    dto.Name,
		Message: dto.Message,
	}
}

func fromDbe(dbe *message.MessageDbe) *helloworld.Message {
	return &helloworld.Message{
		ID:      dbe.ID,
		Name:    dbe.Name,
		Message: dbe.Message,
	}
}
