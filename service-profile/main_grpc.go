package main

import (
	"context"
	"service-profile/profile"
	pb "service-profile/proto"
	"time"
)

func (s *server) GetProfile(ctx context.Context, in *pb.Empty) (*pb.Profile, error) {
	profileDB := profile.NewProfileDB(&s.storage)
	profileService := profile.NewProfileService(profileDB)
	return profileService.GetProfile(ctx)
}
func (s *server) UpdateProfile(ctx context.Context, in *pb.Profile) (*pb.Profile, error) {
	profileDB := profile.NewProfileDB(&s.storage)
	profileService := profile.NewProfileService(profileDB)
	return profileService.UpdateProfile(ctx, in)
}

func (s *server) CountNotesByUserId(ctx context.Context, in *pb.Empty) (*pb.Count, error) {
    noteDB := profile.NewNoteDBImpl(&s.storage)
    noteService := profile.NewNoteServiceImpl(noteDB)
    return noteService.CountNotesByUserId(ctx, in)
}

func (s *server) GetNotesByUserId(page *pb.Page, stream pb.ProfileService_GetNotesByUserIdServer) error {
	noteDB := profile.NewNoteDBImpl(&s.storage)
	noteService := profile.NewNoteServiceImpl(noteDB)
    contextWithTimeout, cancel := context.WithTimeout(stream.Context(), 5*time.Second)
    defer cancel()
	return noteService.GetNotesByUserId(contextWithTimeout, page, stream)
}
func (s *server) GetNoteById(ctx context.Context, in *pb.Id) (*pb.Note, error) {
	noteDB := profile.NewNoteDBImpl(&s.storage)
	noteService := profile.NewNoteServiceImpl(noteDB)
	return noteService.GetNoteById(ctx, in)
}
func (s *server) CreateNote(ctx context.Context, in *pb.Note) (*pb.Note, error) {
	noteDB := profile.NewNoteDBImpl(&s.storage)
	noteService := profile.NewNoteServiceImpl(noteDB)
	return noteService.CreateNote(ctx, in)
}
func (s *server) DeleteNoteById(ctx context.Context, in *pb.Id) (*pb.Empty, error) {
	noteDB := profile.NewNoteDBImpl(&s.storage)
	noteService := profile.NewNoteServiceImpl(noteDB)
    return noteService.DeleteNoteById(ctx, in)
}
