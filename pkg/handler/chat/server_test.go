package chat

import (
	"context"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/sivchari/chat-answer/pkg/domain/entity"
	"github.com/sivchari/chat-answer/pkg/usecase/chat/mock_chat"
	"github.com/sivchari/chat-answer/proto/proto"
	"github.com/sivchari/chat-answer/proto/proto/protoconnect"
)

type mocks struct {
	chatInteractor *mock_chat.MockInteractor
}

func newWithMocks(t *testing.T) (context.Context, protoconnect.ChatServiceHandler, *mocks) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	chatInteractor := mock_chat.NewMockInteractor(ctrl)
	return ctx, NewServer(
			chatInteractor,
		), &mocks{
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

func TestServer_ListRoom(t *testing.T) {
	ctx, s, m := newWithMocks(t)
	m.chatInteractor.EXPECT().ListRoom(ctx).Return([]*entity.Room{{
		ID:   "id",
		Name: "name",
	}}, nil).Times(1)

	res, err := s.ListRoom(ctx, connect.NewRequest(&emptypb.Empty{}))
	assert.Equal(t, connect.NewResponse(&proto.ListRoomResponse{
		Rooms: []*proto.Room{{
			Id:   "id",
			Name: "name",
		}},
	}), res)
	assert.NoError(t, err)
}
