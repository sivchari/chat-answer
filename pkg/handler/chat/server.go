package chat

import (
	"context"
	"fmt"
	"io"

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
	// idを元にroomを取得
	// room, err := s.chatInteracter.GetRoom(ctx, req.Msg.GetMessage().GetRoomId())
	// if err != nil {
	// 	return nil, err
	// }

	// for _, s := range room.Streams {
	// 	// streamに対してsend
	// 	if err := s.PbStream.Send(&proto.SendMessageResponse{
	// 		Message: req.Msg.GetMessage(),
	// 	}); err != nil {
	// 		return nil, err
	// 	}
	// }

	// if err := s.chatInteracter.SendMessage(ctx, req.Msg.GetMessage().GetRoomId(), req.Msg.GetMessage().GetText()); err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (s *server) ListMessages(ctx context.Context, req *connect.Request[proto.ListMessagesRequest], stream *connect.ServerStream[proto.ListMessagesResponse]) error {
	// idを元にroomを取得
	// room, err := s.chatInteracter.GetRoom(ctx, req.Msg.GetRoomId())
	// if err != nil {
	// 	return nil, err
	// }

	// roomに紐づくstreamを取得
	// for _, stream := range room.Streams {
	// 	// streamに対してsend
	// 	if err := stream.Send(&proto.ListMessagesResponse{

	// }

	return nil
}

func (s *server) Chat(ctx context.Context, stream *connect.BidiStream[proto.ChatRequest, proto.ChatResponse]) error {
	for {
		// streamから受け取る
		req, err := stream.Receive()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return nil
		}

		if req.GetMessage() == nil {
			continue
		}

		// idを元にroomを取得
		room, err := s.chatInteracter.GetRoom(ctx, req.GetMessage().GetRoomId())
		if err != nil {
			return err
		}

		if req.IsJoin {
			// roomにstreamを追加
			sID, err := s.chatInteracter.AddStream(ctx, room.ID, stream)
			if err != nil {
				return err
			}

			defer func() {
				if err := s.chatInteracter.DeleteStream(ctx, room.ID, sID); err != nil {
					fmt.Println(err)
				}
			}()

			// roomに存在するすべてのメッセージをsend
			messages, err := s.chatInteracter.ListMessage(ctx, room.ID)
			if err != nil {
				return err
			}

			for _, m := range messages {
				stream.Send(&proto.ChatResponse{
					Message: toProtoMessage(m),
				})
			}

		} else {
			// roomに紐づくstreamに対してsend
			for _, s := range room.Streams {
				if err := s.PbStream.Send(&proto.ChatResponse{
					Message: req.GetMessage(),
				}); err != nil {
					return err
				}
			}
		}
	}
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

func toProtoMessage(message *entity.Message) *proto.Message {
	if message == nil {
		return nil
	}
	return &proto.Message{
		RoomId: message.RoomID,
		Text:   message.Text,
	}
}
