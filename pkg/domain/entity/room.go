package entity

import (
	"github.com/bufbuild/connect-go"
	"github.com/sivchari/chat-answer/proto/proto"
)

type Room struct {
	ID      string
	Name    string
	Streams map[string]*Stream
}

type Stream struct {
	ID       string
	PbStream *connect.BidiStream[proto.ChatRequest, proto.ChatResponse]
}
