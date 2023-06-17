package chat

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/sivchari/chat-answer/pkg/domain/entity"
	"github.com/sivchari/chat-answer/pkg/domain/repository/message"
	"github.com/sivchari/chat-answer/pkg/domain/repository/room"
	"github.com/sivchari/chat-answer/pkg/util"
	"github.com/sivchari/chat-answer/proto/proto"
)

type Interactor interface {
	CreateRoom(ctx context.Context, name string) (*entity.Room, error)
	GetRoom(ctx context.Context, id string) (*entity.Room, error)
	ListRoom(ctx context.Context) ([]*entity.Room, error)
	SendMessage(ctx context.Context, roomID, text string) error
	ListMessage(ctx context.Context, roomID string) ([]*entity.Message, error)
	AddStream(ctx context.Context, roomID string, stream *connect.BidiStream[proto.ChatRequest, proto.ChatResponse]) (string, error)
	DeleteStream(ctx context.Context, roomID, streamID string) error
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
		ID:      id,
		Name:    name,
		Streams: make(map[string]*entity.Stream, 0),
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

func (i *interactor) GetRoom(ctx context.Context, id string) (*entity.Room, error) {
	room, err := i.roomRepository.Select(ctx, id)
	if err != nil {
		return nil, err
	}
	return room, nil
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

func (i *interactor) ListMessage(ctx context.Context, roomID string) ([]*entity.Message, error) {
	messages, err := i.messageRepository.SelectByRoomID(ctx, roomID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (i *interactor) AddStream(ctx context.Context, roomID string, stream *connect.BidiStream[proto.ChatRequest, proto.ChatResponse]) (string, error) {
	id, err := i.ulidGenerator.Generate()
	if err != nil {
		return "", err
	}
	s := &entity.Stream{
		ID:       id,
		PbStream: stream,
	}
	if err := i.roomRepository.InsertStream(ctx, roomID, s); err != nil {
		return "", err
	}
	return id, nil
}

func (i *interactor) DeleteStream(ctx context.Context, roomID, streamID string) error {
	return i.roomRepository.DeleteStream(ctx, roomID, streamID)
}
