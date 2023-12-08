package notes

import (
	"context"
	"fmt"
	"log/slog"
	pb "sgsg/proto"
	"sgsg/users"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetNotesByUserId(stream pb.Service_GetNotesByUserIdServer) error {
	start := time.Now()
	userId, err := users.UserCheck(stream.Context())
	if err != nil {
		slog.Error("Error authorizing user", "users.UserCheck", err)
		return status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	notesCh := make(chan *pb.Note)
	errCh := make(chan error, 1)

	go selectNotesByUserId(notesCh, errCh, userId)

	go func() {
		for note := range notesCh {
			err := stream.Send(note)
			if err != nil {
				errCh <- fmt.Errorf("stream.Send: %w", err)
			}
		}
		errCh <- nil
	}()

	err = <-errCh
	if err != nil {
		slog.Error("Error getting notes", "notes.GetNotes", err)
		return status.Error(codes.Internal, "Internal error")
	}

	slog.Info("GetNotesByUserId", "time", time.Since(start))
	return nil
}

func GetNoteById(ctx context.Context, id string) (*pb.Note, error) {
	start := time.Now()
	userId, err := users.UserCheck(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "users.UserCheck", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	note, err := selectNoteById(id, userId)
	if err != nil {
		slog.Error("Error getting note", "notes.GetNoteById", err)
		return nil, err
	}
	slog.Info("GetNoteById", "time", time.Since(start))
	return note, nil
}

func CreateNote(ctx context.Context, in *pb.Note) (*pb.Note, error) {
	start := time.Now()
	userId, err := users.UserCheck(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "users.UserCheck", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	in.UserId = userId

    err = validateNote(in)
	if err != nil {
        slog.Error("Error validating note", "utils.ValidateStruct", err)
		return nil, err
	}

	var note *pb.Note
	if in.Id == "" {
		note, err = insertNote(in)
	} else {
		note, err = updateNote(in)
	}
	if err != nil {
		slog.Error("Error creating note", "createNote", err)
		return nil, err
	}
	slog.Info("CreateNote", "time", time.Since(start))
	return note, nil

}

func DeleteNoteById(ctx context.Context, id string) (*pb.Empty, error) {
	start := time.Now()
	_, err := users.UserCheck(ctx)
	if err != nil {
		slog.Error("Error authorizing user", "users.UserCheck", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	err = deleteNoteById(id)
	if err != nil {
		slog.Error("Error deleting note", "deleteNoteById", err)
		return nil, err
	}

	slog.Info("DeleteNoteById", "time", time.Since(start))
	return &pb.Empty{}, nil
}
