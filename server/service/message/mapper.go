package message

import (
	"{{.golang.module_path}}/{{.service_name}}/api/helloworld"
	"{{.golang.module_path}}/{{.service_name}}/server/repository/message"
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
