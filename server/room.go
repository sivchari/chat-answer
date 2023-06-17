package server

import (
	"context"

	"github.com/bufbuild/connect-go"

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
