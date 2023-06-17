package chat

import (
	"context"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/sivchari/chat-answer/pkg/domain/entity"
	"github.com/sivchari/chat-answer/pkg/usecase/chat"

	"github.com/sivchari/chat-answer/proto/proto"
	"github.com/sivchari/chat-answer/proto/proto/protoconnect"
)

type server struct {
	chatInteracter chat.Interactor
}

func NewServer(chatInteracter chat.Interactor) protoconnect.ChatServiceHandler {
	return &server{
		chatInteracter,
	}
}

func (s *server) CreateRoom(ctx context.Context, req *connect.Request[proto.CreateRoomRequest]) (*connect.Response[proto.CreateRoomResponse], error) {
	room, err := s.chatInteracter.CreateRoom(ctx, req.Msg.GetName())
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&proto.CreateRoomResponse{
		Id: room.ID,
	}), nil
}

func (s *server) ListRoom(ctx context.Context, _ *connect.Request[emptypb.Empty]) (*connect.Response[proto.ListRoomResponse], error) {
	rooms, err := s.chatInteracter.ListRoom(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&proto.ListRoomResponse{
		Rooms: toProtoRooms(rooms),
	}), nil
}

func (s *server) SendMessage(ctx context.Context, req *connect.Request[proto.SendMessageRequest]) (*connect.Response[emptypb.Empty], error) {
	if err := s.chatInteracter.SendMessage(ctx, req.Msg.GetMessage().GetRoomId(), req.Msg.GetMessage().GetText()); err != nil {
		return nil, err
	}
	return nil, nil
}

func toProtoRoom(room *entity.Room) *proto.Room {
	if room == nil {
		return nil
	}
	return &proto.Room{
		Id:   room.ID,
		Name: room.Name,
	}
}

func toProtoRooms(rooms []*entity.Room) []*proto.Room {
	ret := make([]*proto.Room, 0, len(rooms))
	for _, room := range rooms {
		ret = append(ret, toProtoRoom(room))
	}
	return ret
}
