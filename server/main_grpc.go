package main

import (
	"context"
	"log/slog"
	"time"

	"sgsg/profiles"
    "sgsg/notes"
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

func (s *server) GetProfile(ctx context.Context, in *pb.Empty) (*pb.Profile, error) {
    start := time.Now()
    userId, err := users.UserCheck(ctx)
    if err != nil {
        slog.Error("Error authorizing user", "users.UserCheck", err)
        return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
    }
    profile, err := profiles.GetProfileByUserId(userId)
    if err != nil {
        slog.Error("Error getting profile", "profiles.GetProfileByUserId", err)
        return nil, err
    }
    slog.Info("GetProfileByUserId", "time", time.Since(start))
    return profile, nil
}

func (s *server) CreateProfile(ctx context.Context, in *pb.Profile) (*pb.Profile, error) {
	start := time.Now()
	userId, err := users.UserCheck(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "users.UserCheck", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	in.UserId = userId
	profile, err := profiles.CreateProfile(in)
	if err != nil {
		slog.Error("Error creating profile", "profiles.CreateProfile", err)
		return nil, err
	}
	slog.Info("CreateProfile", "time", time.Since(start))
	return profile, nil
}

func (s *server) DeleteProfile(ctx context.Context, in *pb.Id) (*pb.Empty, error) {
	start := time.Now()
	_, err := users.UserCheck(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "users.UserCheck", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	err = profiles.DeleteProfile(in.Id)
	if err != nil {
		slog.Error("Error deleting profile", "profiles.DeleteProfile", err)
		return nil, err
	}

	slog.Info("DeleteProfile", "time", time.Since(start))
	return &pb.Empty{}, nil
}

func (s *server) GetNotes(in *pb.Empty, stream pb.Service_GetNotesServer) error {
	start := time.Now()
	userId, err := users.UserCheck(stream.Context())
	if err != nil {
		slog.Error("Error authorizing user", "users.UserCheck", err)
		return status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	err = notes.GetNotesStream(stream, userId)
	if err != nil {
		slog.Error("Error getting notes", "notes.GetNotesStream", err)
		return err
	}
	slog.Info("GetNotes", "time", time.Since(start))
	return nil
}

func (s *server) GetNoteById(ctx context.Context, in *pb.Id) (*pb.Note, error) {
    start := time.Now()
    userId, err := users.UserCheck(ctx)
    if err != nil {
        slog.Error("Error authorizing user", "users.UserCheck", err)
        return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
    }
    note, err := notes.GetNoteById(in.Id, userId)
    if err != nil {
        slog.Error("Error getting note", "notes.GetNoteById", err)
        return nil, err
    }
    slog.Info("GetNoteById", "time", time.Since(start))
    return note, nil
}

func (s *server) CreateNote(ctx context.Context, in *pb.Note) (*pb.Note, error) {
	start := time.Now()
	userId, err := users.UserCheck(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "users.UserCheck", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	in.UserId = userId
	note, err := notes.CreateNote(in)
	if err != nil {
		slog.Error("Error creating note", "notes.CreateNote", err)
		return nil, err
	}
	slog.Info("CreateNote", "time", time.Since(start))
	return note, nil
}

func (s *server) DeleteNote(ctx context.Context, in *pb.Id) (*pb.Empty, error) {
	start := time.Now()
	_, err := users.UserCheck(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "users.UserCheck", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	err = notes.DeleteNote(in.Id)
	if err != nil {
		slog.Error("Error deleting note", "notes.DeleteNote", err)
		return nil, err
	}

	slog.Info("DeleteNote", "time", time.Since(start))
	return &pb.Empty{}, nil
}
