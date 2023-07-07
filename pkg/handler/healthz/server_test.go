package healthz

import (
	"context"
	"fmt"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/sivchari/chat-answer/pkg/log/mock_log"
	"github.com/sivchari/chat-answer/pkg/xcontext"
	"github.com/sivchari/chat-answer/proto/proto"
	"github.com/sivchari/chat-answer/proto/proto/protoconnect"
)

type mocks struct {
	logger *mock_log.MockHandler
}

func newWithMocks(t *testing.T) (context.Context, protoconnect.HealthzHandler, *mocks) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	logger := mock_log.NewMockHandler(ctrl)
	return ctx, New(
			logger,
		), &mocks{
			logger,
		}
}

func TestServer_Check(t *testing.T) {
	userID := "userID"
	name := "name"

	ctx, s, m := newWithMocks(t)
	ctx = xcontext.WithValue[xcontext.UserID, string](ctx, userID)
	m.logger.EXPECT().InfoCtx(ctx, "healthz check", "name", name, "userID", userID)

	res, err := s.Check(ctx, connect.NewRequest(&proto.CheckRequest{
		Name: name,
	}))
	assert.Equal(t, connect.NewResponse(&proto.CheckResponse{
		Msg: fmt.Sprintf("Hello %s", name),
	}), res)
	assert.NoError(t, err)
}
