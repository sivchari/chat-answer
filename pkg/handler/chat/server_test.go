package chat

import (
	"context"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/sivchari/chat-answer/pkg/domain/entity"
	"github.com/sivchari/chat-answer/pkg/log/mock_log"
	"github.com/sivchari/chat-answer/pkg/usecase/chat/mock_chat"
	"github.com/sivchari/chat-answer/proto/proto"
	"github.com/sivchari/chat-answer/proto/proto/protoconnect"
)

type mocks struct {
	logger         *mock_log.MockHandler
	chatInteractor *mock_chat.MockInteractor
}

func newWithMocks(t *testing.T) (context.Context, protoconnect.ChatServiceHandler, *mocks) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	logger := mock_log.NewMockHandler(ctrl)
	chatInteractor := mock_chat.NewMockInteractor(ctrl)
	return ctx, New(
			logger,
			chatInteractor,
		), &mocks{
			logger,
			chatInteractor,
		}
}

func TestServer_CreateRoom(t *testing.T) {
	id := "id"
	name := "name"

	ctx, s, m := newWithMocks(t)
	m.chatInteractor.EXPECT().CreateRoom(ctx, name).Return(&entity.Room{
		ID:   id,
		Name: name,
	}, nil).Times(1)

	res, err := s.CreateRoom(ctx, connect.NewRequest(&proto.CreateRoomRequest{
		Name: name,
	}))
	assert.Equal(t, connect.NewResponse(&proto.CreateRoomResponse{
		Id: id,
	}), res)
	assert.NoError(t, err)
}

func TestService_GetRoom(t *testing.T) {
	id := "roomID"
	name := "name"

	ctx, s, m := newWithMocks(t)
	m.chatInteractor.EXPECT().GetRoom(ctx, id).Return(&entity.Room{
		ID:   id,
		Name: name,
	}, nil).Times(1)

	res, err := s.GetRoom(ctx, connect.NewRequest(&proto.GetRoomRequest{
		Id: id,
	}))
	assert.Equal(t, connect.NewResponse(&proto.GetRoomResponse{
		Room: &proto.Room{
			Id:   id,
			Name: name,
		},
	}), res)
	assert.NoError(t, err)
}

func TestServer_ListRoom(t *testing.T) {
	id := "id"
	name := "name"

	ctx, s, m := newWithMocks(t)
	m.chatInteractor.EXPECT().ListRoom(ctx).Return([]*entity.Room{{
		ID:   id,
		Name: name,
	}}, nil).Times(1)

	res, err := s.ListRoom(ctx, connect.NewRequest(&emptypb.Empty{}))
	assert.Equal(t, connect.NewResponse(&proto.ListRoomResponse{
		Rooms: []*proto.Room{{
			Id:   id,
			Name: name,
		}},
	}), res)
	assert.NoError(t, err)
}

func TestServer_GetPass(t *testing.T) {
	pass := "pass"

	ctx, s, m := newWithMocks(t)
	m.chatInteractor.EXPECT().GetPass(ctx).Return(pass, nil).Times(1)

	res, err := s.GetPass(ctx, connect.NewRequest(&emptypb.Empty{}))
	assert.Equal(t, connect.NewResponse(&proto.GetPassResponse{
		Pass: pass,
	}), res)
	assert.NoError(t, err)
}

func TestServer_ListMessage(t *testing.T) {
	roomID := "roomID"
	message := &entity.Message{
		RoomID: roomID,
		Text:   "text",
	}

	ctx, s, m := newWithMocks(t)
	m.chatInteractor.EXPECT().ListMessage(ctx, roomID).Return([]*entity.Message{message}, nil).Times(1)

	res, err := s.ListMessage(ctx, connect.NewRequest(&proto.ListMessageRequest{
		RoomId: roomID,
	}))
	assert.Equal(t, connect.NewResponse(&proto.ListMessageResponse{
		Messages: []*proto.Message{
			{
				RoomId: message.RoomID,
				Text:   message.Text,
			},
		},
	}), res)
	assert.NoError(t, err)
}
