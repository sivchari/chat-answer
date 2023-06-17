package server

import (
	"context"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/sivchari/chat-answer/proto/proto"
	"github.com/sivchari/chat-answer/usecase"
)

type RoomService struct {
	RoomUC usecase.RoomUC
}

func (r *RoomService) CreateRoom(ctx context.Context, req *connect.Request[proto.CreateRoomRequest]) (*connect.Response[proto.CreateRoomResponse], error) {
	resp, err := r.RoomUC.CreateRoom(ctx, &usecase.CreateRoomRequest{
		Room: req.Msg.Room,
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&proto.CreateRoomResponse{
		Room: resp.Room,
	}), nil
}

func (r *RoomService) ListRooms(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[proto.ListRoomsResponse], error) {
	resp, err := r.RoomUC.ListRooms(ctx)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&proto.ListRoomsResponse{
		Rooms: resp.Rooms,
	}), nil
}
