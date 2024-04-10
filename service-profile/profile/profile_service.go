package profile

import (
	"context"
	"log/slog"
	pb "service-profile/proto"
	"service-profile/system"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProfileService interface {
	GetProfile(ctx context.Context) (*pb.Profile, error)
	UpdateProfile(ctx context.Context, profile *pb.Profile) (*pb.Profile, error)
}

type profileService struct {
	ProfileDB
}

func NewProfileService(db ProfileDB) ProfileService {
	return &profileService{db}
}

func (s *profileService) GetProfile(ctx context.Context) (*pb.Profile, error) {
	defer system.Perf("get_profile", time.Now())
	claims, err := system.ExtractToken(ctx)
	if err != nil {
		slog.Error("Error extracting token", "err", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	profile, exists, err := s.selectProfileByUserId(claims.Id)
	if !exists {
		profile, err = s.insertProfile(&pb.Profile{UserId: claims.Id, Active: false})
	}
	if err != nil {
		slog.Error("Error selecting profile by user id", "err", err)
		return nil, status.Error(codes.NotFound, "Profile not found")
	}
	return profile, nil
}

func (s *profileService) UpdateProfile(ctx context.Context, profile *pb.Profile) (*pb.Profile, error) {
	defer system.Perf("update_profile", time.Now())
	claims, err := system.ExtractToken(ctx)
	if err != nil {
		slog.Error("Error extracting token", "err", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	validationErrors := validateProfile(profile)
	if len(validationErrors) > 0 {
		return nil, system.CreateErrors(validationErrors)
	}

	profile.UserId = claims.Id
	profile, err = s.updateProfile(profile)
	if err != nil {
		slog.Error("Error updating profile", "err", err)
		return nil, status.Error(codes.Internal, "Error updating profile")
	}
	return profile, nil
}
