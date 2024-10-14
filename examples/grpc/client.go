package grpc

import (
	"github.com/AntonLuning/tiny-url/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(endpointAddr string) (*proto.ServiceClient, error) {
	conn, err := grpc.NewClient(endpointAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := proto.NewServiceClient(conn)

	return &client, nil
}
