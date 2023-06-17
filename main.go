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

	"github.com/sivchari/chat-answer/proto/proto/protoconnect"
	"github.com/sivchari/chat-answer/repository"
	"github.com/sivchari/chat-answer/server"
	"github.com/sivchari/chat-answer/usecase"
)

func main() {
	os.Exit(run())
}

func run() int {
	const (
		ok = 0
		ng = 1
	)

	healthzService := &server.HealthzService{}
	roomService := &server.RoomService{
		RoomUC: &usecase.RoomUCImpl{
			RoomRepo: &repository.RoomRepositoryImpl{},
		},
	}

	mux := http.NewServeMux()
	mux.Handle(protoconnect.NewHealthzHandler(healthzService))
	mux.Handle(protoconnect.NewRoomServiceHandler(roomService))
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
