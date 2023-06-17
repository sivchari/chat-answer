package room

import (
	"context"

	"github.com/sivchari/chat-answer/pkg/domain/entity"
)

type Repository interface {
	Insert(ctx context.Context, room *entity.Room) error
	SelectAll(ctx context.Context) ([]*entity.Room, error)
}
