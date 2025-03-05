package req_ping

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"connectrpc.com/connect"
	next_go_proto_v1 "github.com/yiritani/next_go_proto/sdk/golang/_generated/protos/next_go/v1"
	next_go_proto_v1connect "github.com/yiritani/next_go_proto/sdk/golang/_generated/protos/next_go/v1/next_go_proto_v1connect"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	client := next_go_proto_v1connect.NewPingServiceClient(
		http.DefaultClient,
		"https://grpc-backend-gateway-9uia1nx.uc.gateway.dev",
	)
	res, err := client.Ping(
		context.Background(), 
		connect.NewRequest(&next_go_proto_v1.PingRequest{
			Message: "hello",
		}),
	)
	if err != nil {
		log.Fatalf("Failed to ping: %v", err)
	}

	log.Printf("Ping response: %v", res.Msg)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": res.Msg.Message,
	})
	w.WriteHeader(http.StatusOK)
}
