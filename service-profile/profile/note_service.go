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

type NoteService interface {
	CountNotesByUserId(ctx context.Context, empty *pb.Empty) (*pb.Count, error)
	GetNotesByUserId(ctx context.Context, page *pb.Page, stream pb.ProfileService_GetNotesByUserIdServer) error
	GetNoteById(ctx context.Context, id *pb.Id) (*pb.Note, error)
	CreateNote(ctx context.Context, note *pb.Note) (*pb.Note, error)
	DeleteNoteById(ctx context.Context, id *pb.Id) (*pb.Empty, error)
}

type NoteServiceImpl struct {
	db NoteDB
}

func NewNoteServiceImpl(db NoteDB) NoteService {
	return &NoteServiceImpl{db: db}
}

func (n *NoteServiceImpl) CountNotesByUserId(ctx context.Context, empty *pb.Empty) (*pb.Count, error) {
	defer system.Perf("count_notes_by_user_id", time.Now())
	user, err := system.ExtractToken(ctx)
	if err != nil {
		slog.Error("Error extracting token", "error", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	count, err := n.db.CountNotesByUserId(ctx, user.Id)
	if err != nil {
		slog.Error("Error counting notes", "user_id", user.Id, "error", err)
		return nil, status.Error(codes.Internal, "Error counting notes")
	}
	return &pb.Count{Count: int32(count)}, nil
}

func (n *NoteServiceImpl) GetNotesByUserId(ctx context.Context, page *pb.Page, stream pb.ProfileService_GetNotesByUserIdServer) error {
	defer system.Perf("get_notes_by_user_id", time.Now())
	user, err := system.ExtractToken(stream.Context())
	if err != nil {
		slog.Error("Error extracting token", "error", err)
		return status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	noteChan := make(chan *pb.Note)
	errChan := make(chan error, 1)
	go func() {
		defer close(noteChan)
		defer close(errChan)
		n.db.GetNotesByUserId(ctx, noteChan, errChan, user.Id, int64(page.Limit), int64(page.Offset))
	}()

	for note := range noteChan {
		select {
		case <-ctx.Done():
			slog.Error("Context canceled", "user_id", user.Id, "error", ctx.Err())
			return status.Error(codes.Internal, "Context canceled")
		default:
		}
		err = stream.Send(note)
		if err != nil {
			slog.Error("Error sending note", "error", err)
			return status.Error(codes.Internal, "Error sending note")
		}
	}
	err = <-errChan
	if err != nil {
		slog.Error("Error getting notes", "user_id", user.Id, "error", err)
		return status.Error(codes.Internal, "Error getting notes")
	}
	return nil
}

func (n *NoteServiceImpl) GetNoteById(ctx context.Context, id *pb.Id) (*pb.Note, error) {
	defer system.Perf("get_note_by_id", time.Now())
	note, err := n.db.GetNoteById(ctx, id)
	if err != nil {
		slog.Error("Error getting note", "id", id, "error", err)
		return nil, status.Error(codes.Internal, "Error getting note")
	}
	return note, err
}

func (n *NoteServiceImpl) CreateNote(ctx context.Context, note *pb.Note) (*pb.Note, error) {
	defer system.Perf("create_note", time.Now())
	user, err := system.ExtractToken(ctx)
	if err != nil {
		slog.Error("Error extracting token", "error", err)
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}
	note.UserId = user.Id
	if note.Id == "" {
		note, err = n.db.InsertNote(ctx, note)
	} else {
		note, err = n.db.UpdateNoteById(ctx, note)
	}
	if err != nil {
		slog.Error("Error creating note", "error", err)
		return nil, status.Error(codes.Internal, "Error creating note")
	}
	return note, err
}

func (n *NoteServiceImpl) DeleteNoteById(ctx context.Context, id *pb.Id) (*pb.Empty, error) {
	defer system.Perf("delete_note_by_id", time.Now())
	err := n.db.DeleteNoteById(ctx, id)
	if err != nil {
		slog.Error("Error deleting note", "id", id, "error", err)
		return nil, status.Error(codes.Internal, "Error deleting note")
	}
	return &pb.Empty{}, nil
}
