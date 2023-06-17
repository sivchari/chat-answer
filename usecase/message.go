package usecase

import (
	"context"

	"github.com/sivchari/chat-answer/entity"
	"github.com/sivchari/chat-answer/repository"
)

type MessageUC interface {
	SendMessage(ctx context.Context, request *SendMessageRequest) (*SendMessageResponse, error)
}

type SendMessageRequest struct {
	Message *entity.Message
}

type SendMessageResponse struct {
	Message *entity.Message
}

type MessageUCImpl struct {
	MessageRepo repository.MessageRepository
}

func (uc *MessageUCImpl) SendMessage(ctx context.Context, request *SendMessageRequest) (*SendMessageResponse, error) {
	message, err := uc.MessageRepo.SendMessage(ctx, request.Message)
	if err != nil {
		return nil, err
	}
	return &SendMessageResponse{message}, nil
}
