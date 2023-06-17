package room

import (
	"context"

	"github.com/sivchari/chat-answer/pkg/domain/entity"
)

type Repository interface {
	Insert(ctx context.Context, room *entity.Room) error
	InsertStream(ctx context.Context, roomID string, stream *entity.Stream) error
	DeleteStream(ctx context.Context, roomID string, streamID string) error
	Select(ctx context.Context, id string) (*entity.Room, error)
	SelectAll(ctx context.Context) ([]*entity.Room, error)
}
