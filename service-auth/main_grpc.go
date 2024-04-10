package main

import (
	"context"

	"service-auth/auth"
	pb "service-auth/proto"
)

func (s *server) Auth(ctx context.Context, in *pb.Empty) (*pb.AuthResponse, error) {
    var authDb = auth.NewAuthDB(&s.storage)
	var auth = auth.NewAuthService(authDb)
	return auth.Auth(ctx)
}

func (s *server) CreateStripeCheckout(ctx context.Context, in *pb.Empty) (*pb.StripeUrlResponse, error) {
    var authDb = auth.NewAuthDB(&s.storage)
	var auth = auth.NewAuthService(authDb)
	return auth.CreateStripeCheckout(ctx)
}

func (s *server) CreateStripePortal(ctx context.Context, in *pb.Empty) (*pb.StripeUrlResponse, error) {
    var authDb = auth.NewAuthDB(&s.storage)
	var auth = auth.NewAuthService(authDb)
	return auth.CreateStripePortal(ctx)
}
