package chat

import (
	"context"
	"errors"
	"io"
	"sync"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/sivchari/chat-answer/pkg/domain/entity"
	"github.com/sivchari/chat-answer/pkg/usecase/chat"

	"github.com/sivchari/chat-answer/proto/proto"
	"github.com/sivchari/chat-answer/proto/proto/protoconnect"
)

type server struct {
	chatInteracter  chat.Interactor
	streamsByRoomID map[string]Streams
	mu              sync.RWMutex
}

type Streams map[string]*connect.BidiStream[proto.ChatRequest, proto.ChatResponse]

func NewServer(chatInteracter chat.Interactor) protoconnect.ChatServiceHandler {
	return &server{
		chatInteracter,
		make(map[string]Streams, 0),
		sync.RWMutex{},
	}
}

func (s *server) initStreams(roomID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.streamsByRoomID[roomID] = make(Streams, 0)
}

func (s *server) addStream(roomID string, stream *connect.BidiStream[proto.ChatRequest, proto.ChatResponse]) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	streams, ok := s.streamsByRoomID[roomID]
	if !ok {
		return ""
	}

	key := stream.Peer().Addr
	streams[key] = stream
	s.streamsByRoomID[roomID] = streams

	return key
}

func (s *server) deleteStream(roomID, key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	streams, ok := s.streamsByRoomID[roomID]
	if !ok {
		return
	}
	delete(streams, key)
	s.streamsByRoomID[roomID] = streams
	return
}

func (s *server) getStreams(roomID string) Streams {
	s.mu.RLock()
	defer s.mu.RUnlock()

	streams, ok := s.streamsByRoomID[roomID]
	if !ok {
		return nil
	}
	return streams
}

func (s *server) existStream(roomID string, key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	streams, ok := s.streamsByRoomID[roomID]
	if !ok {
		return false
	}
	_, ok = streams[key]
	return ok
}

func (s *server) CreateRoom(ctx context.Context, req *connect.Request[proto.CreateRoomRequest]) (*connect.Response[proto.CreateRoomResponse], error) {
	room, err := s.chatInteracter.CreateRoom(ctx, req.Msg.GetName())
	if err != nil {
		return nil, err
	}
	s.initStreams(room.ID)
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

func (s *server) ListMessage(ctx context.Context, req *connect.Request[proto.ListMessageRequest]) (*connect.Response[proto.ListMessageResponse], error) {
	messages, err := s.chatInteracter.ListMessage(ctx, req.Msg.GetRoomId())
	if err != nil {
		return nil, err
	}

	return &connect.Response[proto.ListMessageResponse]{Msg: &proto.ListMessageResponse{
		Messages: toProtoMessages(messages),
	}}, nil
}

func (s *server) Chat(ctx context.Context, stream *connect.BidiStream[proto.ChatRequest, proto.ChatResponse]) error {
	for {
		req, err := stream.Receive()
		if err != nil {
			// stream closed
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		if req.GetMessage() == nil {
			continue
		}

		room, err := s.chatInteracter.GetRoom(ctx, req.GetMessage().GetRoomId())
		if err != nil {
			return err
		}

		joined := s.existStream(room.ID, stream.Peer().Addr)
		if !joined {
			sID := s.addStream(room.ID, stream)
			defer func() {
				s.deleteStream(room.ID, sID)
			}()
		}

		streams := s.getStreams(room.ID)
		for _, st := range streams {
			if err := s.chatInteracter.SendMessage(ctx, req.GetMessage().GetRoomId(), req.GetMessage().GetText()); err != nil {
				return err
			}

			if err := st.Send(&proto.ChatResponse{
				Message: req.GetMessage(),
			}); err != nil {
				return err
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

func toProtoMessages(messages []*entity.Message) []*proto.Message {
	ret := make([]*proto.Message, 0, len(messages))
	for _, message := range messages {
		ret = append(ret, toProtoMessage(message))
	}
	return ret
}
