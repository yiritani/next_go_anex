package cmd

import (
	"context"
	"log"

	proto "backend/src/_generated/proto/v1"

	"connectrpc.com/connect"
)

type PingServer struct{}

func (p *PingServer) Ping(
	ctx context.Context,
	req *connect.Request[proto.PingRequest],
) (*connect.Response[proto.PingResponse], error) {
	log.Println("ping controller")

	res := connect.NewResponse(&proto.PingResponse{
		Message: "Pong from grpc_anex",
	})
	res.Header().Set("Content-Type", "application/json")
	return res, nil
}
