// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/chat.proto

package protoconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	proto "github.com/sivchari/chat-answer/proto/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ChatServiceName is the fully-qualified name of the ChatService service.
	ChatServiceName = "api.ChatService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ChatServiceCreateRoomProcedure is the fully-qualified name of the ChatService's CreateRoom RPC.
	ChatServiceCreateRoomProcedure = "/api.ChatService/CreateRoom"
	// ChatServiceGetRoomProcedure is the fully-qualified name of the ChatService's GetRoom RPC.
	ChatServiceGetRoomProcedure = "/api.ChatService/GetRoom"
	// ChatServiceListRoomProcedure is the fully-qualified name of the ChatService's ListRoom RPC.
	ChatServiceListRoomProcedure = "/api.ChatService/ListRoom"
	// ChatServiceListMessageProcedure is the fully-qualified name of the ChatService's ListMessage RPC.
	ChatServiceListMessageProcedure = "/api.ChatService/ListMessage"
	// ChatServiceChatProcedure is the fully-qualified name of the ChatService's Chat RPC.
	ChatServiceChatProcedure = "/api.ChatService/Chat"
)

// ChatServiceClient is a client for the api.ChatService service.
type ChatServiceClient interface {
	CreateRoom(context.Context, *connect_go.Request[proto.CreateRoomRequest]) (*connect_go.Response[proto.CreateRoomResponse], error)
	GetRoom(context.Context, *connect_go.Request[proto.GetRoomRequest]) (*connect_go.Response[proto.GetRoomResponse], error)
	ListRoom(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[proto.ListRoomResponse], error)
	ListMessage(context.Context, *connect_go.Request[proto.ListMessageRequest]) (*connect_go.Response[proto.ListMessageResponse], error)
	Chat(context.Context) *connect_go.BidiStreamForClient[proto.ChatRequest, proto.ChatResponse]
}

// NewChatServiceClient constructs a client for the api.ChatService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewChatServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ChatServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &chatServiceClient{
		createRoom: connect_go.NewClient[proto.CreateRoomRequest, proto.CreateRoomResponse](
			httpClient,
			baseURL+ChatServiceCreateRoomProcedure,
			opts...,
		),
		getRoom: connect_go.NewClient[proto.GetRoomRequest, proto.GetRoomResponse](
			httpClient,
			baseURL+ChatServiceGetRoomProcedure,
			opts...,
		),
		listRoom: connect_go.NewClient[emptypb.Empty, proto.ListRoomResponse](
			httpClient,
			baseURL+ChatServiceListRoomProcedure,
			opts...,
		),
		listMessage: connect_go.NewClient[proto.ListMessageRequest, proto.ListMessageResponse](
			httpClient,
			baseURL+ChatServiceListMessageProcedure,
			opts...,
		),
		chat: connect_go.NewClient[proto.ChatRequest, proto.ChatResponse](
			httpClient,
			baseURL+ChatServiceChatProcedure,
			opts...,
		),
	}
}

// chatServiceClient implements ChatServiceClient.
type chatServiceClient struct {
	createRoom  *connect_go.Client[proto.CreateRoomRequest, proto.CreateRoomResponse]
	getRoom     *connect_go.Client[proto.GetRoomRequest, proto.GetRoomResponse]
	listRoom    *connect_go.Client[emptypb.Empty, proto.ListRoomResponse]
	listMessage *connect_go.Client[proto.ListMessageRequest, proto.ListMessageResponse]
	chat        *connect_go.Client[proto.ChatRequest, proto.ChatResponse]
}

// CreateRoom calls api.ChatService.CreateRoom.
func (c *chatServiceClient) CreateRoom(ctx context.Context, req *connect_go.Request[proto.CreateRoomRequest]) (*connect_go.Response[proto.CreateRoomResponse], error) {
	return c.createRoom.CallUnary(ctx, req)
}

// GetRoom calls api.ChatService.GetRoom.
func (c *chatServiceClient) GetRoom(ctx context.Context, req *connect_go.Request[proto.GetRoomRequest]) (*connect_go.Response[proto.GetRoomResponse], error) {
	return c.getRoom.CallUnary(ctx, req)
}

// ListRoom calls api.ChatService.ListRoom.
func (c *chatServiceClient) ListRoom(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[proto.ListRoomResponse], error) {
	return c.listRoom.CallUnary(ctx, req)
}

// ListMessage calls api.ChatService.ListMessage.
func (c *chatServiceClient) ListMessage(ctx context.Context, req *connect_go.Request[proto.ListMessageRequest]) (*connect_go.Response[proto.ListMessageResponse], error) {
	return c.listMessage.CallUnary(ctx, req)
}

// Chat calls api.ChatService.Chat.
func (c *chatServiceClient) Chat(ctx context.Context) *connect_go.BidiStreamForClient[proto.ChatRequest, proto.ChatResponse] {
	return c.chat.CallBidiStream(ctx)
}

// ChatServiceHandler is an implementation of the api.ChatService service.
type ChatServiceHandler interface {
	CreateRoom(context.Context, *connect_go.Request[proto.CreateRoomRequest]) (*connect_go.Response[proto.CreateRoomResponse], error)
	GetRoom(context.Context, *connect_go.Request[proto.GetRoomRequest]) (*connect_go.Response[proto.GetRoomResponse], error)
	ListRoom(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[proto.ListRoomResponse], error)
	ListMessage(context.Context, *connect_go.Request[proto.ListMessageRequest]) (*connect_go.Response[proto.ListMessageResponse], error)
	Chat(context.Context, *connect_go.BidiStream[proto.ChatRequest, proto.ChatResponse]) error
}

// NewChatServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewChatServiceHandler(svc ChatServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ChatServiceCreateRoomProcedure, connect_go.NewUnaryHandler(
		ChatServiceCreateRoomProcedure,
		svc.CreateRoom,
		opts...,
	))
	mux.Handle(ChatServiceGetRoomProcedure, connect_go.NewUnaryHandler(
		ChatServiceGetRoomProcedure,
		svc.GetRoom,
		opts...,
	))
	mux.Handle(ChatServiceListRoomProcedure, connect_go.NewUnaryHandler(
		ChatServiceListRoomProcedure,
		svc.ListRoom,
		opts...,
	))
	mux.Handle(ChatServiceListMessageProcedure, connect_go.NewUnaryHandler(
		ChatServiceListMessageProcedure,
		svc.ListMessage,
		opts...,
	))
	mux.Handle(ChatServiceChatProcedure, connect_go.NewBidiStreamHandler(
		ChatServiceChatProcedure,
		svc.Chat,
		opts...,
	))
	return "/api.ChatService/", mux
}

// UnimplementedChatServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedChatServiceHandler struct{}

func (UnimplementedChatServiceHandler) CreateRoom(context.Context, *connect_go.Request[proto.CreateRoomRequest]) (*connect_go.Response[proto.CreateRoomResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.ChatService.CreateRoom is not implemented"))
}

func (UnimplementedChatServiceHandler) GetRoom(context.Context, *connect_go.Request[proto.GetRoomRequest]) (*connect_go.Response[proto.GetRoomResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.ChatService.GetRoom is not implemented"))
}

func (UnimplementedChatServiceHandler) ListRoom(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[proto.ListRoomResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.ChatService.ListRoom is not implemented"))
}

func (UnimplementedChatServiceHandler) ListMessage(context.Context, *connect_go.Request[proto.ListMessageRequest]) (*connect_go.Response[proto.ListMessageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.ChatService.ListMessage is not implemented"))
}

func (UnimplementedChatServiceHandler) Chat(context.Context, *connect_go.BidiStream[proto.ChatRequest, proto.ChatResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.ChatService.Chat is not implemented"))
}
