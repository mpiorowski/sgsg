package profiles

import (
	"context"
	"database/sql"
	"log/slog"
	"sgsg/auth"
	pb "sgsg/proto"
	"sgsg/system"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProfileService interface {
	GetProfileByUserId(ctx context.Context) (*pb.Profile, error)
	CreateProfile(ctx context.Context, in *pb.Profile) (*pb.Profile, error)
}

type ProfileServiceImpl struct {
	ProfileDB
	auth.AuthService
}

func NewProfileService(db ProfileDB, auth auth.AuthService) ProfileService {
	return &ProfileServiceImpl{db, auth}
}

func (s *ProfileServiceImpl) GetProfileByUserId(ctx context.Context) (*pb.Profile, error) {
	defer system.Perf("get_profile_by_user_id", time.Now())
	user, err := s.GetUser(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "auth.GetUser", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	profile, err := s.SelectProfileByUserID(user.Id)
    if err == sql.ErrNoRows {
        return &pb.Profile{}, nil
    }
	if err != nil {
		slog.Error("Error getting profile", "db.SelectProfileByUserID", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return profile, nil
}

func (s *ProfileServiceImpl) CreateProfile(ctx context.Context, in *pb.Profile) (*pb.Profile, error) {
	defer system.Perf("create_profile", time.Now())
	user, err := s.GetUser(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "auth.GetUser", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	in.UserId = user.Id

	validationErrors := validateProfile(in)
	if len(validationErrors) > 0 {
		return nil, system.CreateErrors(validationErrors)
	}

	var profile *pb.Profile
	if in.Id == "" {
		profile, err = s.InsertProfile(in)
	} else {
		profile, err = s.UpdateProfile(in)
	}
	if err != nil {
		slog.Error("Error creating profile", "createProfile", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return profile, nil
}
