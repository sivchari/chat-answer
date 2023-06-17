package server

import (
	"context"

	"github.com/bufbuild/connect-go"

	"github.com/sivchari/chat-answer/proto/proto"
	"github.com/sivchari/chat-answer/usecase"
)

type MessageService struct {
	MessageUC usecase.MessageUC
}

func (r *MessageService) SendMessage(ctx context.Context, req *connect.Request[proto.SendMessageRequest]) (*connect.Response[proto.SendMessageResponse], error) {
	resp, err := r.MessageUC.SendMessage(ctx, &usecase.SendMessageRequest{
		Message: req.Msg.Message,
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&proto.SendMessageResponse{
		Message: resp.Message,
	}), nil
}
