package chat

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/sivchari/chat-answer/pkg/codes"
	"github.com/sivchari/chat-answer/pkg/domain/entity"
	"github.com/sivchari/chat-answer/pkg/domain/repository/message/mock_message"
	"github.com/sivchari/chat-answer/pkg/domain/repository/room/mock_room"
	"github.com/sivchari/chat-answer/pkg/errors"
	"github.com/sivchari/chat-answer/pkg/ulid/mock_ulid"
)

type mocks struct {
	ulidGenerator     *mock_ulid.MockULIDGenerator
	roomRepository    *mock_room.MockRepository
	messageRepository *mock_message.MockRepository
}

func newWithMocks(t *testing.T) (context.Context, Interactor, *mocks) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	ulidGenerator := mock_ulid.NewMockULIDGenerator(ctrl)
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
	t.Run("【正常系】", func(t *testing.T) {
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
	})

	t.Run("【異常系】nameが空文字列", func(t *testing.T) {
		name := ""

		ctx, i, _ := newWithMocks(t)

		res, err := i.CreateRoom(ctx, name)
		assert.Nil(t, res)
		assert.Equal(t, errors.Code(err), codes.CodeInvalidArgument)
		assert.Equal(t, errors.Message(err), "name is required")
	})
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

func TestInteractor_GetRoom(t *testing.T) {
	t.Run("【正常系】", func(t *testing.T) {
		roomID := "roomID"

		ctx, i, m := newWithMocks(t)
		m.roomRepository.EXPECT().Select(ctx, roomID).Return(&entity.Room{
			ID: roomID,
		}, nil).Times(1)

		res, err := i.GetRoom(ctx, roomID)
		assert.Equal(t, &entity.Room{
			ID: roomID,
		}, res)
		assert.NoError(t, err)
	})

	t.Run("【異常系】roomIDが空", func(t *testing.T) {
		roomID := ""

		ctx, i, _ := newWithMocks(t)

		res, err := i.GetRoom(ctx, roomID)
		assert.Nil(t, res)
		assert.Equal(t, errors.Code(err), codes.CodeInvalidArgument)
		assert.Equal(t, errors.Message(err), "roomID is required")
	})
}

func TestInteractor_GetPass(t *testing.T) {
	id := "id"

	ctx, i, m := newWithMocks(t)
	m.ulidGenerator.EXPECT().Generate().Return(id, nil).Times(1)

	res, err := i.GetPass(ctx)
	assert.Equal(t, id, res)
	assert.NoError(t, err)
}

func TestInteractor_SendMessage(t *testing.T) {
	t.Run("【正常系】", func(t *testing.T) {
		roomID := "roomID"
		text := "text"

		ctx, i, m := newWithMocks(t)
		m.messageRepository.EXPECT().Insert(ctx, &entity.Message{
			RoomID: roomID,
			Text:   text,
		}).Return(nil).Times(1)

		err := i.SendMessage(ctx, roomID, text)
		assert.NoError(t, err)
	})

	t.Run("【異常系】roomIDが空文字列", func(t *testing.T) {
		roomID := ""
		text := "text"

		ctx, i, _ := newWithMocks(t)

		err := i.SendMessage(ctx, roomID, text)
		assert.Equal(t, errors.Code(err), codes.CodeInvalidArgument)
		assert.Equal(t, errors.Message(err), "roomID is required")
	})

	t.Run("【異常系】textが空文字列", func(t *testing.T) {
		roomID := "roomID"
		text := ""

		ctx, i, _ := newWithMocks(t)

		err := i.SendMessage(ctx, roomID, text)
		assert.Equal(t, errors.Code(err), codes.CodeInvalidArgument)
		assert.Equal(t, errors.Message(err), "text is required")
	})
}

func TestInteractor_ListMessage(t *testing.T) {
	t.Run("【正常系】", func(t *testing.T) {
		roomID := "roomID"
		text := "text"

		ctx, i, m := newWithMocks(t)
		m.messageRepository.EXPECT().SelectByRoomID(ctx, roomID).Return([]*entity.Message{{
			RoomID: roomID,
			Text:   text,
		}}, nil).Times(1)

		res, err := i.ListMessage(ctx, roomID)
		assert.Equal(t, []*entity.Message{{
			RoomID: roomID,
			Text:   text,
		}}, res)
		assert.NoError(t, err)
	})

	t.Run("【異常系】roomIDが空文字列", func(t *testing.T) {
		roomID := ""

		ctx, i, _ := newWithMocks(t)

		res, err := i.ListMessage(ctx, roomID)
		assert.Nil(t, res)
		assert.Equal(t, errors.Code(err), codes.CodeInvalidArgument)
		assert.Equal(t, errors.Message(err), "roomID is required")
	})
}
