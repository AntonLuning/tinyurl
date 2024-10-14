package api

import (
	"context"
	"log/slog"
	"net"

	"github.com/AntonLuning/tiny-url/proto"
	"github.com/AntonLuning/tiny-url/service"
	"github.com/AntonLuning/tiny-url/utils"
	"google.golang.org/grpc"
)

type GRPCAPIServer struct {
	addr    string
	service service.Service
	proto.UnimplementedServiceServer
}

func NewGRPCAPIServer(addr string, service service.Service) *GRPCAPIServer {
	return &GRPCAPIServer{
		addr:    addr,
		service: service,
	}
}

func (s *GRPCAPIServer) Run() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	proto.RegisterServiceServer(server, s)

	slog.Info("gRPC API server starting", "address", s.addr)

	return server.Serve(ln)
}

func (s *GRPCAPIServer) CreateShortenURL(ctx context.Context, r *proto.ShortenURLRequest) (*proto.ShortenURLResponse, error) {
	ctx = utils.SetContextValues(ctx, "gRPC")

	shortenURL, err := s.service.CreateShortenURL(ctx, r.Original)
	if err != nil {
		return nil, err
	}

	resp := proto.ShortenURLResponse{
		Original: r.Original,
		Shorten:  *shortenURL,
	}

	return &resp, nil
}

func (s *GRPCAPIServer) GetOriginalURL(ctx context.Context, r *proto.OriginalURLRequest) (*proto.OriginalURLResponse, error) {
	ctx = utils.SetContextValues(ctx, "gRPC")

	originalURL, err := s.service.GetOriginalURL(ctx, r.Shorten)
	if err != nil {
		return nil, err
	}

	resp := proto.OriginalURLResponse{
		Shorten:  r.Shorten,
		Original: *originalURL,
	}

	return &resp, nil
}
