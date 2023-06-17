package chat

import (
	"context"

	"github.com/sivchari/chat-answer/pkg/domain/entity"
	"github.com/sivchari/chat-answer/pkg/domain/repository/message"
	"github.com/sivchari/chat-answer/pkg/domain/repository/room"
	"github.com/sivchari/chat-answer/pkg/util"
)

type Interactor interface {
	CreateRoom(ctx context.Context, name string) (*entity.Room, error)
	ListRoom(ctx context.Context) ([]*entity.Room, error)
	SendMessage(ctx context.Context, roomID, text string) error
}

type interactor struct {
	ulidGenerator     util.ULIDGenerator
	roomRepository    room.Repository
	messageRepository message.Repository
}

func NewInteractor(
	ulidGenerator util.ULIDGenerator,
	roomRepository room.Repository,
	messageRepository message.Repository,
) Interactor {
	return &interactor{
		ulidGenerator,
		roomRepository,
		messageRepository,
	}
}

func (i *interactor) CreateRoom(ctx context.Context, name string) (*entity.Room, error) {
	id, err := i.ulidGenerator.Generate()
	if err != nil {
		return nil, err
	}
	room := &entity.Room{
		ID:   id,
		Name: name,
	}
	if err := i.roomRepository.Insert(ctx, room); err != nil {
		return nil, err
	}
	return room, nil
}

func (i *interactor) ListRoom(ctx context.Context) ([]*entity.Room, error) {
	rooms, err := i.roomRepository.SelectAll(ctx)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (i *interactor) SendMessage(ctx context.Context, roomID, text string) error {
	message := &entity.Message{
		RoomID: roomID,
		Text:   text,
	}
	if err := i.messageRepository.Insert(ctx, message); err != nil {
		return err
	}
	return nil
}
