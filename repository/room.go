package repository

import (
	"sync"

	"github.com/sivchari/chat-answer/entity"
)

type RoomRepository interface {
	CreateRoom(room *entity.Room) (*entity.Room, error)
	ListRooms() ([]*entity.Room, error)
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

func (r *RoomRepositoryImpl) ListRooms() ([]*entity.Room, error) {
	mu.Lock()
	defer mu.Unlock()

	var rooms []*entity.Room
	for _, room := range store {
		rooms = append(rooms, room)
	}
	return rooms, nil
}
