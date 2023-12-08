package profiles

import (
	"context"
	"log/slog"
	pb "sgsg/proto"
	"sgsg/users"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetProfileByUserId(ctx context.Context, in *pb.Empty) (*pb.Profile, error) {
	start := time.Now()
	userId, err := users.UserCheck(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "users.UserCheck", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	profile, err := selectProfileByUserId(userId)
	if err != nil {
		slog.Error("Error getting profile", "selectProfileByUserId", err)
		return nil, err
	}
	slog.Info("GetProfileByUserId", "time", time.Since(start))
	return profile, nil
}

func CreateProfile(ctx context.Context, in *pb.Profile) (*pb.Profile, error) {
	start := time.Now()
	userId, err := users.UserCheck(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "users.UserCheck", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	in.UserId = userId

	err = validateProfile(in)
	if err != nil {
        slog.Error("Error validating profile", "validateProfile", err)
		return nil, err
	}

	var profile *pb.Profile
	if in.Id == "" {
		profile, err = insertProfile(in)
	} else {
		profile, err = updateProfile(in)
	}
	if err != nil {
		slog.Error("Error creating profile", "createProfile", err)
		return nil, err
	}
	slog.Info("CreateProfile", "time", time.Since(start))
	return profile, nil
}

func DeleteProfileById(ctx context.Context, id string) (*pb.Empty, error) {
	start := time.Now()
	_, err := users.UserCheck(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "users.UserCheck", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	err = deleteProfileById(id)
	if err != nil {
		slog.Error("Error deleting profile", "deleteProfileById", err)
		return nil, err
	}

	slog.Info("DeleteProfileById", "time", time.Since(start))
	return &pb.Empty{}, nil
}
