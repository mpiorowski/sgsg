package notes

import (
	"context"
	"fmt"
	"log/slog"
	"sgsg/auth"
	pb "sgsg/proto"
	"sgsg/system"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NoteService interface {
	GetNotesByUserId(stream pb.Service_GetNotesByUserIdServer) error
	GetNoteById(ctx context.Context, id string) (*pb.Note, error)
	CreateNote(ctx context.Context, in *pb.Note) (*pb.Note, error)
	DeleteNoteById(ctx context.Context, id string) (*pb.Empty, error)
}

type noteService struct {
	NoteDB
	auth.AuthService
}

func NewNoteService(db NoteDB, auth auth.AuthService) NoteService {
	return &noteService{db, auth}
}

func (s *noteService) GetNotesByUserId(stream pb.Service_GetNotesByUserIdServer) error {
	defer system.Perf("get_notes_by_user_id", time.Now())
	user, err := s.GetUser(stream.Context())
	if err != nil {
		slog.Error("Error authorizing user", "s.GetUser", err)
		return status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	notesCh := make(chan *pb.Note)
	errCh := make(chan error, 1)
	defer close(errCh)

	go s.SelectNotesByUserID(notesCh, errCh, user.Id)

	for note := range notesCh {
		err := stream.Send(note)
		if err != nil {
			return fmt.Errorf("stream.Send: %w", err)
		}
	}

	if len(errCh) > 0 {
		slog.Error("Error getting notes", "notes.GetNotes", <-errCh)
		return status.Error(codes.Internal, "Internal error")
	}
	return nil
}

func (s *noteService) GetNoteById(ctx context.Context, id string) (*pb.Note, error) {
	defer system.Perf("get_note_by_id", time.Now())
	user, err := s.GetUser(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "s.GetUser", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	note, err := s.SelectNoteByID(id, user.Id)
	if err != nil {
		slog.Error("Error getting note", "s.SelectNoteByID", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return note, nil
}

func (s *noteService) CreateNote(ctx context.Context, in *pb.Note) (*pb.Note, error) {
	defer system.Perf("create_note", time.Now())
	user, err := s.GetUser(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "s.GetUser", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	in.UserId = user.Id

	validationErrors := validateNote(in)
	if len(validationErrors) > 0 {
		return nil, system.CreateErrors(validationErrors)
	}

	var note *pb.Note
	if in.Id == "" {
		note, err = s.InsertNote(in)
	} else {
		note, err = s.UpdateNote(in)
	}
	if err != nil {
		slog.Error("Error creating note", "createNote", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return note, nil

}

func (s *noteService) DeleteNoteById(ctx context.Context, id string) (*pb.Empty, error) {
	defer system.Perf("delete_note_by_id", time.Now())
	_, err := s.GetUser(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "s.GetUser", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	err = s.DeleteNoteByID(id)
	if err != nil {
		slog.Error("Error deleting note", "s.DeleteNoteByID", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &pb.Empty{}, nil
}
