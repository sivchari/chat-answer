package healthz

import (
	"context"
	"fmt"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"

	"github.com/sivchari/chat-answer/proto/proto"
)

func TestServer_Check(t *testing.T) {
	name := "name"

	s := NewServer()
	res, err := s.Check(context.Background(), connect.NewRequest(&proto.CheckRequest{
		Name: name,
	}))
	assert.Equal(t, connect.NewResponse(&proto.CheckResponse{
		Msg: fmt.Sprintf("Hello %s", name),
	}), res)
	assert.NoError(t, err)
}
