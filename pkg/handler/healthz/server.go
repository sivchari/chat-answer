package healthz

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"

	"github.com/sivchari/chat-answer/proto/proto"
	"github.com/sivchari/chat-answer/proto/proto/protoconnect"
)

type server struct{}

func NewServer() protoconnect.HealthzHandler {
	return &server{}
}

func (s *server) Check(_ context.Context, req *connect.Request[proto.CheckRequest]) (*connect.Response[proto.CheckResponse], error) {
	return connect.NewResponse(&proto.CheckResponse{
		Msg: fmt.Sprintf("Hello %s", req.Msg.GetName()),
	}), nil
}
