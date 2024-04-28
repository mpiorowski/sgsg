package main

import (
	"context"

	pb "service-auth/proto"
	"service-auth/service"
)

func (s *server) Auth(ctx context.Context, in *pb.Empty) (*pb.AuthResponse, error) {
	return service.Auth(ctx, s.storage)
}

func (s *server) CreateStripeCheckout(ctx context.Context, in *pb.Empty) (*pb.StripeUrlResponse, error) {
	return service.CreateStripeCheckout(ctx, s.storage)
}

func (s *server) CreateStripePortal(ctx context.Context, in *pb.Empty) (*pb.StripeUrlResponse, error) {
	return service.CreateStripePortal(ctx, s.storage)
}
