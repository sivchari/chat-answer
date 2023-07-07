package interceptors

import (
	"context"

	"github.com/bufbuild/connect-go"
	"google.golang.org/grpc/metadata"

	"github.com/sivchari/chat-answer/pkg/xcontext"
)

type authInterceptor struct {
}

func NewAuthInterceptor() *authInterceptor {
	return &authInterceptor{}
}

const AuthTokenHeader = "x-auth-token"

func (i *authInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		var authToken string
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if values, ok := md[AuthTokenHeader]; ok {
				authToken = values[0]
			}
		}
		// TODO: トークンから認証サーバへ問い合わせてUserIDを取得
		_ = authToken
		userID := "mockUserID"

		valuectx := xcontext.WithValue[xcontext.UserID, string](ctx, userID)
		res, err := next(valuectx, req)
		if err != nil {
			return nil, toConnectError(err)
		}
		return res, nil
	}
}

func (i *authInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(ctx context.Context, spec connect.Spec) connect.StreamingClientConn {
		return next(ctx, spec)
	}
}

func (i *authInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(ctx context.Context, conn connect.StreamingHandlerConn) error {
		var authToken string
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if values, ok := md[AuthTokenHeader]; ok {
				authToken = values[0]
			}
		}
		// TODO: トークンから認証サーバへ問い合わせてUserIDを取得
		_ = authToken
		userID := "mockUserID"

		valuectx := xcontext.WithValue[xcontext.UserID](ctx, userID)
		if err := next(valuectx, conn); err != nil {
			return toConnectError(err)
		}
		return nil
	}
}
