package repository

import (
	"sync"

	"github.com/sivchari/chat-answer/entity"
)

type RoomRepository interface {
	CreateRoom(room *entity.Room) (*entity.Room, error)
}

type RoomRepositoryImpl struct {
}

var store map[string]*entity.Room = make(map[string]*entity.Room)
var mu sync.Mutex

func NewRoomRepository() *RoomRepositoryImpl {
	return &RoomRepositoryImpl{}
}

func (r *RoomRepositoryImpl) CreateRoom(room *entity.Room) (*entity.Room, error) {
	mu.Lock()
	defer mu.Unlock()

	store[room.GetId()] = room
	return room, nil
}
