package message

import (
	"context"

	"github.com/sivchari/chat-answer/pkg/domain/entity"
)

type Repository interface {
	Insert(ctx context.Context, message *entity.Message) error
}