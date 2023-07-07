package message

import (
	"context"
	"sync"

	"github.com/sivchari/chat-answer/pkg/domain/entity"
	"github.com/sivchari/chat-answer/pkg/domain/repository/message"
)

type repository struct {
	mapByRoomID map[string][]*entity.Message
	mu          sync.RWMutex
}

func New() message.Repository {
	return &repository{
		mapByRoomID: make(map[string][]*entity.Message, 0),
	}
}

func (r *repository) Insert(_ context.Context, message *entity.Message) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.mapByRoomID[message.RoomID] = append(r.mapByRoomID[message.RoomID], message)
	return nil
}

func (r *repository) SelectByRoomID(_ context.Context, roomID string) ([]*entity.Message, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	messages, ok := r.mapByRoomID[roomID]
	if !ok {
		return []*entity.Message{}, nil
	}

	return messages, nil
}
