// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/healthz.proto

package protoconnect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect_go "github.com/bufbuild/connect-go"

	proto "github.com/sivchari/chat-answer/proto/proto"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// HealthzName is the fully-qualified name of the Healthz service.
	HealthzName = "api.Healthz"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// HealthzCheckProcedure is the fully-qualified name of the Healthz's Check RPC.
	HealthzCheckProcedure = "/api.Healthz/Check"
)

// HealthzClient is a client for the api.Healthz service.
type HealthzClient interface {
	Check(context.Context, *connect_go.Request[proto.CheckRequest]) (*connect_go.Response[proto.CheckResponse], error)
}

// NewHealthzClient constructs a client for the api.Healthz service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHealthzClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) HealthzClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &healthzClient{
		check: connect_go.NewClient[proto.CheckRequest, proto.CheckResponse](
			httpClient,
			baseURL+HealthzCheckProcedure,
			opts...,
		),
	}
}

// healthzClient implements HealthzClient.
type healthzClient struct {
	check *connect_go.Client[proto.CheckRequest, proto.CheckResponse]
}

// Check calls api.Healthz.Check.
func (c *healthzClient) Check(ctx context.Context, req *connect_go.Request[proto.CheckRequest]) (*connect_go.Response[proto.CheckResponse], error) {
	return c.check.CallUnary(ctx, req)
}

// HealthzHandler is an implementation of the api.Healthz service.
type HealthzHandler interface {
	Check(context.Context, *connect_go.Request[proto.CheckRequest]) (*connect_go.Response[proto.CheckResponse], error)
}

// NewHealthzHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHealthzHandler(svc HealthzHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(HealthzCheckProcedure, connect_go.NewUnaryHandler(
		HealthzCheckProcedure,
		svc.Check,
		opts...,
	))
	return "/api.Healthz/", mux
}

// UnimplementedHealthzHandler returns CodeUnimplemented from all methods.
type UnimplementedHealthzHandler struct{}

func (UnimplementedHealthzHandler) Check(context.Context, *connect_go.Request[proto.CheckRequest]) (*connect_go.Response[proto.CheckResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.Healthz.Check is not implemented"))
}
