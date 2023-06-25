package interceptors

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/sivchari/chat-answer/pkg/errcodes"
)

type errorInterceptor struct{}

func NewErrorInterceptor() *errorInterceptor {
	return &errorInterceptor{}
}

func (i *errorInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return connect.UnaryFunc(func(
		ctx context.Context,
		req connect.AnyRequest,
	) (connect.AnyResponse, error) {
		res, err := next(ctx, req)
		if err != nil {
			return nil, toConnectError(err)
		}
		return res, nil
	})
}

func (i *errorInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return connect.StreamingClientFunc(func(
		ctx context.Context,
		spec connect.Spec,
	) connect.StreamingClientConn {
		return next(ctx, spec)
	})
}

func (i *errorInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return connect.StreamingHandlerFunc(func(
		ctx context.Context,
		conn connect.StreamingHandlerConn,
	) error {
		return toConnectError(next(ctx, conn))
	})
}

func toConnectError(err error) *connect.Error {
	if err == nil {
		return nil
	}

	c := errcodes.NewCode(err)
	switch c {
	case errcodes.CodeUnknown:
		return connect.NewError(connect.CodeUnknown, err)
	case errcodes.CodeInvalidArgument:
		return connect.NewError(connect.CodeInvalidArgument, err)
	case errcodes.CodeNotFound:
		return connect.NewError(connect.CodeNotFound, err)
	case errcodes.CodeInternal:
		return connect.NewError(connect.CodeInternal, err)
	case errcodes.CodeOK:
		return nil
	}
	return connect.NewError(connect.CodeUnknown, err)
}
