package main

import (
	"context"
	"log/slog"
	"time"

	"sgsg/notes"
	"sgsg/profiles"
	pb "sgsg/proto"
	"sgsg/users"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) Auth(ctx context.Context, in *pb.Empty) (*pb.AuthResponse, error) {
	start := time.Now()
	user, tokenId, err := users.UserAuth(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "users.UserAuth", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	slog.Info("Auth", "time", time.Since(start))
	return &pb.AuthResponse{
		TokenId: tokenId,
		User:    user,
	}, nil
}

func (s *server) GetProfileByUserId(ctx context.Context, in *pb.Empty) (*pb.Profile, error) {
    return profiles.GetProfileByUserId(ctx, in)
}

func (s *server) CreateProfile(ctx context.Context, in *pb.Profile) (*pb.Profile, error) {
    return profiles.CreateProfile(ctx, in)
}

func (s *server) GetNotesByUserId(in *pb.Empty, stream pb.Service_GetNotesByUserIdServer) error {
	return notes.GetNotesByUserId(stream)

}

func (s *server) GetNoteById(ctx context.Context, in *pb.Id) (*pb.Note, error) {
	return notes.GetNoteById(ctx, in.Id)
}

func (s *server) CreateNote(ctx context.Context, in *pb.Note) (*pb.Note, error) {
    return notes.CreateNote(ctx, in)
}

func (s *server) DeleteNoteById(ctx context.Context, in *pb.Id) (*pb.Empty, error) {
    return notes.DeleteNoteById(ctx, in.Id)
}
