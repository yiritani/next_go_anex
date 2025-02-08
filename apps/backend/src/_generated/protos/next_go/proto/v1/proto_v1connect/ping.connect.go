// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: protos/next_go/proto/v1/ping.proto

package proto_v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "grpc_backend/src/_generated/proto/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// PingServiceName is the fully-qualified name of the PingService service.
	PingServiceName = "next_go.api.v1.ping.PingService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PingServicePingProcedure is the fully-qualified name of the PingService's Ping RPC.
	PingServicePingProcedure = "/next_go.api.v1.ping.PingService/Ping"
	// PingServicePingSampleProcedure is the fully-qualified name of the PingService's PingSample RPC.
	PingServicePingSampleProcedure = "/next_go.api.v1.ping.PingService/PingSample"
)

// PingServiceClient is a client for the next_go.api.v1.ping.PingService service.
type PingServiceClient interface {
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
	PingSample(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
}

// NewPingServiceClient constructs a client for the next_go.api.v1.ping.PingService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPingServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PingServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	pingServiceMethods := v1.File_protos_next_go_proto_v1_ping_proto.Services().ByName("PingService").Methods()
	return &pingServiceClient{
		ping: connect.NewClient[v1.PingRequest, v1.PingResponse](
			httpClient,
			baseURL+PingServicePingProcedure,
			connect.WithSchema(pingServiceMethods.ByName("Ping")),
			connect.WithClientOptions(opts...),
		),
		pingSample: connect.NewClient[v1.PingRequest, v1.PingResponse](
			httpClient,
			baseURL+PingServicePingSampleProcedure,
			connect.WithSchema(pingServiceMethods.ByName("PingSample")),
			connect.WithClientOptions(opts...),
		),
	}
}

// pingServiceClient implements PingServiceClient.
type pingServiceClient struct {
	ping       *connect.Client[v1.PingRequest, v1.PingResponse]
	pingSample *connect.Client[v1.PingRequest, v1.PingResponse]
}

// Ping calls next_go.api.v1.ping.PingService.Ping.
func (c *pingServiceClient) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// PingSample calls next_go.api.v1.ping.PingService.PingSample.
func (c *pingServiceClient) PingSample(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return c.pingSample.CallUnary(ctx, req)
}

// PingServiceHandler is an implementation of the next_go.api.v1.ping.PingService service.
type PingServiceHandler interface {
	Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
	PingSample(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error)
}

// NewPingServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPingServiceHandler(svc PingServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	pingServiceMethods := v1.File_protos_next_go_proto_v1_ping_proto.Services().ByName("PingService").Methods()
	pingServicePingHandler := connect.NewUnaryHandler(
		PingServicePingProcedure,
		svc.Ping,
		connect.WithSchema(pingServiceMethods.ByName("Ping")),
		connect.WithHandlerOptions(opts...),
	)
	pingServicePingSampleHandler := connect.NewUnaryHandler(
		PingServicePingSampleProcedure,
		svc.PingSample,
		connect.WithSchema(pingServiceMethods.ByName("PingSample")),
		connect.WithHandlerOptions(opts...),
	)
	return "/next_go.api.v1.ping.PingService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PingServicePingProcedure:
			pingServicePingHandler.ServeHTTP(w, r)
		case PingServicePingSampleProcedure:
			pingServicePingSampleHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPingServiceHandler struct{}

func (UnimplementedPingServiceHandler) Ping(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("next_go.api.v1.ping.PingService.Ping is not implemented"))
}

func (UnimplementedPingServiceHandler) PingSample(context.Context, *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("next_go.api.v1.ping.PingService.PingSample is not implemented"))
}
