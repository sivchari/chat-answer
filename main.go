package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/sivchari/chat-answer/proto/proto"
	"github.com/sivchari/chat-answer/proto/proto/protoconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type HealthzService struct{}

func (h *HealthzService) Check(ctx context.Context, req *connect.Request[proto.CheckRequest]) (*connect.Response[proto.CheckResponse], error) {
	resp := &proto.CheckResponse{
		Msg: fmt.Sprintf("Hello %s", req.Msg.Name),
	}
	return connect.NewResponse(resp), nil
}

func main() {
	mux := http.NewServeMux()
	healthzService := &HealthzService{}
	mux.Handle(protoconnect.NewHealthzHandler(healthzService))
	log.Println(http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{})))
}
