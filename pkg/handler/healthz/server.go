package healthz

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"

	"github.com/sivchari/chat-answer/pkg/log"
	"github.com/sivchari/chat-answer/pkg/xcontext"
	"github.com/sivchari/chat-answer/proto/proto"
	"github.com/sivchari/chat-answer/proto/proto/protoconnect"
)

type server struct {
	logger log.Handler
}

func New(logger log.Handler) protoconnect.HealthzHandler {
	return &server{
		logger: logger,
	}
}

func (s *server) Check(ctx context.Context, req *connect.Request[proto.CheckRequest]) (*connect.Response[proto.CheckResponse], error) {
	userID, _ := xcontext.Value[xcontext.UserID, string](ctx)
	s.logger.InfoCtx(ctx, "healthz check", "name", req.Msg.GetName(), "userID", userID)
	return connect.NewResponse(&proto.CheckResponse{
		Msg: fmt.Sprintf("Hello %s", req.Msg.GetName()),
	}), nil
}
