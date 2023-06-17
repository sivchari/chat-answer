package usecase

import (
	"context"

	"github.com/sivchari/chat-answer/entity"
	"github.com/sivchari/chat-answer/repository"
)

type CreateRoomRequest struct {
	Room *entity.Room
}

type CreateRoomResponse struct {
	Room *entity.Room
}

type ListRoomsResponse struct {
	Rooms []*entity.Room
}

type RoomUC interface {
	CreateRoom(ctx context.Context, request *CreateRoomRequest) (*CreateRoomResponse, error)
	ListRooms(ctx context.Context) (*ListRoomsResponse, error)
}

type RoomUCImpl struct {
	RoomRepo repository.RoomRepository
}

func (uc *RoomUCImpl) CreateRoom(ctx context.Context, request *CreateRoomRequest) (*CreateRoomResponse, error) {
	room, err := uc.RoomRepo.CreateRoom(request.Room)
	if err != nil {
		return nil, err
	}
	return &CreateRoomResponse{room}, nil
}

func (uc *RoomUCImpl) ListRooms(ctx context.Context) (*ListRoomsResponse, error) {
	rooms, err := uc.RoomRepo.ListRooms()
	if err != nil {
		return nil, err
	}
	return &ListRoomsResponse{Rooms: rooms}, nil
}
