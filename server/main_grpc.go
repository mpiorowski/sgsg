package main

import (
	"context"

	"sgsg/auth"
	"sgsg/notes"
	"sgsg/profiles"
	pb "sgsg/proto"
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

func (s *server) GetProfileByUserId(ctx context.Context, in *pb.Empty) (*pb.Profile, error) {
	var authDb = auth.NewAuthDB(&s.storage)
	var auth = auth.NewAuthService(authDb)
	var profileDb = profiles.NewProfileDB(&s.storage)
	var profiles = profiles.NewProfileService(profileDb, auth)
	return profiles.GetProfileByUserId(ctx)
}

func (s *server) CreateProfile(ctx context.Context, in *pb.Profile) (*pb.Profile, error) {
	var authDb = auth.NewAuthDB(&s.storage)
	var auth = auth.NewAuthService(authDb)
	var profileDb = profiles.NewProfileDB(&s.storage)
	var profiles = profiles.NewProfileService(profileDb, auth)
	return profiles.CreateProfile(ctx, in)
}

func (s *server) GetNotesByUserId(in *pb.Empty, stream pb.Service_GetNotesByUserIdServer) error {
	var authDb = auth.NewAuthDB(&s.storage)
	var auth = auth.NewAuthService(authDb)
	var noteDb = notes.NewNoteDB(&s.storage)
	var notes = notes.NewNoteService(noteDb, auth)
	return notes.GetNotesByUserId(stream)

}

func (s *server) GetNoteById(ctx context.Context, in *pb.Id) (*pb.Note, error) {
	var authDb = auth.NewAuthDB(&s.storage)
	var auth = auth.NewAuthService(authDb)
	var noteDb = notes.NewNoteDB(&s.storage)
	var notes = notes.NewNoteService(noteDb, auth)
	return notes.GetNoteById(ctx, in.Id)
}

func (s *server) CreateNote(ctx context.Context, in *pb.Note) (*pb.Note, error) {
	var authDb = auth.NewAuthDB(&s.storage)
	var auth = auth.NewAuthService(authDb)
	var noteDb = notes.NewNoteDB(&s.storage)
	var notes = notes.NewNoteService(noteDb, auth)
	return notes.CreateNote(ctx, in)
}

func (s *server) DeleteNoteById(ctx context.Context, in *pb.Id) (*pb.Empty, error) {
	var authDb = auth.NewAuthDB(&s.storage)
	var auth = auth.NewAuthService(authDb)
	var noteDb = notes.NewNoteDB(&s.storage)
	var notes = notes.NewNoteService(noteDb, auth)
	return notes.DeleteNoteById(ctx, in.Id)
}
