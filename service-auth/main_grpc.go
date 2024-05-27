package main

import (
	"context"

	pb "service-auth/proto"
	"service-auth/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Auth(ctx context.Context, in *pb.Empty) (*pb.AuthResponse, error) {
	r, err := service.Auth(ctx, s.storage)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	return r, nil
}

func (s *server) CreateStripeCheckout(ctx context.Context, in *pb.Empty) (*pb.StripeUrlResponse, error) {
	r, err := service.CreateStripeCheckout(ctx, s.storage)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error creating stripe user")
	}
	return r, nil
}

func (s *server) CreateStripePortal(ctx context.Context, in *pb.Empty) (*pb.StripeUrlResponse, error) {
	r, err := service.CreateStripePortal(ctx, s.storage)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error creating stripe portal")
	}
	return r, nil
}
