package interceptors

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/sivchari/chat-answer/pkg/ulid"
	"github.com/sivchari/chat-answer/pkg/xcontext"
)

type ulidInterceptor struct {
	generator ulid.ULIDGenerator
}

func NewULIDInterceptor(ulidGenerator ulid.ULIDGenerator) *ulidInterceptor {
	return &ulidInterceptor{
		generator: ulidGenerator,
	}
}

func (i *ulidInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(
		ctx context.Context,
		req connect.AnyRequest,
	) (connect.AnyResponse, error) {
		id, err := i.generator.Generate()
		if err != nil {
			// TODO: error handling (toConnectError)
			return nil, err
		}
		valuectx := xcontext.WithValue[xcontext.ULID](ctx, id)
		res, err := next(valuectx, req)
		if err != nil {
			return nil, toConnectError(err)
		}
		return res, nil
	})
}

func (i *ulidInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return connect.StreamingClientFunc(func(
		ctx context.Context,
		spec connect.Spec,
	) connect.StreamingClientConn {
		return next(ctx, spec)
	})
}

func (i *ulidInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return connect.StreamingHandlerFunc(func(
		ctx context.Context,
		conn connect.StreamingHandlerConn,
	) error {
		id, err := i.generator.Generate()
		if err != nil {
			// TODO: error handling (toConnectError)
			return err
		}
		valuectx := xcontext.WithValue[xcontext.ULID](ctx, id)
		if err := next(valuectx, conn); err != nil {
			return toConnectError(err)
		}
		return nil
	})
}
