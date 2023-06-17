package room

import (
	"context"
	"sync"

	"github.com/sivchari/chat-answer/pkg/domain/entity"
	"github.com/sivchari/chat-answer/pkg/domain/repository/room"
)

type repository struct {
	mapByID map[string]*entity.Room
	mu      sync.RWMutex
}

func NewRepository() room.Repository {
	return &repository{
		mapByID: make(map[string]*entity.Room, 0),
	}
}

func (r *repository) Insert(_ context.Context, room *entity.Room) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.mapByID[room.ID] = room
	return nil
}

func (r *repository) SelectAll(_ context.Context) ([]*entity.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	rooms := make([]*entity.Room, 0, len(r.mapByID))
	for _, e := range r.mapByID {
		rooms = append(rooms, e)
	}
	return rooms, nil
}
