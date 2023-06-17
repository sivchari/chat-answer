package repository

import (
	"context"
	"sync"

	"github.com/sivchari/chat-answer/entity"
)

type MessageRepository interface {
	SendMessage(ctx context.Context, message *entity.Message) (*entity.Message, error)
}

var mStore []*entity.Message = make([]*entity.Message, 0)
var mMu sync.Mutex

type MessageRepositoryImpl struct {
}

func (r *MessageRepositoryImpl) SendMessage(ctx context.Context, message *entity.Message) (*entity.Message, error) {
	mMu.Lock()
	defer mMu.Unlock()

	mStore = append(mStore, message)
	return message, nil
}
