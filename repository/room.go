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

var rStore map[string]*entity.Room = make(map[string]*entity.Room)
var rMu sync.Mutex

func NewRoomRepository() *RoomRepositoryImpl {
	return &RoomRepositoryImpl{}
}

func (r *RoomRepositoryImpl) CreateRoom(room *entity.Room) (*entity.Room, error) {
	rMu.Lock()
	defer rMu.Unlock()

	rStore[room.GetId()] = room
	return room, nil
}

func (r *RoomRepositoryImpl) ListRooms() ([]*entity.Room, error) {
	rMu.Lock()
	defer rMu.Unlock()

	var rooms []*entity.Room
	for _, room := range rStore {
		rooms = append(rooms, room)
	}
	return rooms, nil
}
