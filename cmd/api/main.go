package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/bufbuild/connect-go"
	"github.com/rs/cors"

	"github.com/sivchari/chat-answer/pkg/handler/chat"
	"github.com/sivchari/chat-answer/pkg/handler/healthz"
	"github.com/sivchari/chat-answer/pkg/handler/interceptors"
	messagerepository "github.com/sivchari/chat-answer/pkg/infra/repository/message"
	roomrepository "github.com/sivchari/chat-answer/pkg/infra/repository/room"
	"github.com/sivchari/chat-answer/pkg/log"
	"github.com/sivchari/chat-answer/pkg/ulid"
	chatinteractor "github.com/sivchari/chat-answer/pkg/usecase/chat"
	"github.com/sivchari/chat-answer/proto/proto/protoconnect"
)

func main() {
	os.Exit(run())
}

func run() int {
	const (
		ok = 0
		ng = 1
	)

	// DI
	logger := log.NewHandler(log.LevelInfo, log.WithJSONFormat())
	ulidGenerator := ulid.NewUILDGenerator()
	roomRepository := roomrepository.New()
	messageRepository := messagerepository.New()
	chatInteractor := chatinteractor.NewInteractor(ulidGenerator, roomRepository, messageRepository)
	healthzServer := healthz.New(logger)
	chatServer := chat.New(logger, chatInteractor)

	mux := http.NewServeMux()
	mux.Handle(protoconnect.NewHealthzHandler(healthzServer))
	mux.Handle(protoconnect.NewChatServiceHandler(chatServer, connect.WithInterceptors(
		interceptors.NewErrorInterceptor(),
	)))
	handler := cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{}))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorCtx(ctx, "failed to ListenAndServe", "err", err)
		}
	}()

	<-ctx.Done()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(timeoutCtx); err != nil {
		return ng
	}
	return ok
}
