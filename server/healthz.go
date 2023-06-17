package server

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/sivchari/chat-answer/proto/proto"
)

type HealthzService struct{}

func (h *HealthzService) Check(ctx context.Context, req *connect.Request[proto.CheckRequest]) (*connect.Response[proto.CheckResponse], error) {
	resp := &proto.CheckResponse{
		Msg: fmt.Sprintf("Hello %s", req.Msg.Name),
	}
	return connect.NewResponse(resp), nil
}
