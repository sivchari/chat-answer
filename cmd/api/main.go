package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/sivchari/chat-answer/pkg/handler/chat"
	"github.com/sivchari/chat-answer/pkg/handler/healthz"
	messagerepository "github.com/sivchari/chat-answer/pkg/infra/repository/message"
	roomrepository "github.com/sivchari/chat-answer/pkg/infra/repository/room"
	chatinteractor "github.com/sivchari/chat-answer/pkg/usecase/chat"
	"github.com/sivchari/chat-answer/pkg/util"
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
	healthzServer := healthz.NewServer()
	ulidGenerator := util.NewUILDGenerator()
	roomRepository := roomrepository.NewRepository()
	messageRepository := messagerepository.NewRepository()
	chatInteractor := chatinteractor.NewInteractor(ulidGenerator, roomRepository, messageRepository)
	chatServer := chat.NewServer(chatInteractor)

	mux := http.NewServeMux()
	mux.Handle(protoconnect.NewHealthzHandler(healthzServer))
	mux.Handle(protoconnect.NewChatServiceHandler(chatServer))
	handler := h2c.NewHandler(mux, &http2.Server{})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Println(err)
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