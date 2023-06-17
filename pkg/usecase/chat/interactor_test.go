package chat

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/sivchari/chat-answer/pkg/domain/entity"
	"github.com/sivchari/chat-answer/pkg/domain/repository/message/mock_message"
	"github.com/sivchari/chat-answer/pkg/domain/repository/room/mock_room"
	"github.com/sivchari/chat-answer/pkg/util/mock_util"
)

type mocks struct {
	ulidGenerator     *mock_util.MockULIDGenerator
	roomRepository    *mock_room.MockRepository
	messageRepository *mock_message.MockRepository
}

func newWithMocks(t *testing.T) (context.Context, Interactor, *mocks) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	ulidGenerator := mock_util.NewMockULIDGenerator(ctrl)
	roomRepository := mock_room.NewMockRepository(ctrl)
	messageRepository := mock_message.NewMockRepository(ctrl)
	return ctx, NewInteractor(
			ulidGenerator,
			roomRepository,
			messageRepository,
		), &mocks{
			ulidGenerator,
			roomRepository,
			messageRepository,
		}
}

func TestInteractor_CreateRoom(t *testing.T) {
	id := "id"
	name := "name"

	ctx, i, m := newWithMocks(t)
	m.ulidGenerator.EXPECT().Generate().Return(id, nil).Times(1)
	m.roomRepository.EXPECT().Insert(ctx, &entity.Room{
		ID:   id,
		Name: name,
	}).Return(nil).Times(1)

	res, err := i.CreateRoom(ctx, name)
	assert.Equal(t, &entity.Room{
		ID:   id,
		Name: name,
	}, res)
	assert.NoError(t, err)
}

func TestInteractor_ListRoom(t *testing.T) {
	id := "id"
	name := "name"

	ctx, i, m := newWithMocks(t)
	m.roomRepository.EXPECT().SelectAll(ctx).Return([]*entity.Room{{
		ID:   id,
		Name: name,
	}}, nil).Times(1)

	res, err := i.ListRoom(ctx)
	assert.Equal(t, []*entity.Room{{
		ID:   id,
		Name: name,
	}}, res)
	assert.NoError(t, err)
}
