package profile

import (
	"context"
	pb "service-profile/proto"
	"service-profile/system"

	"github.com/google/uuid"
)

type NoteDB interface {
	CountNotesByUserId(ctx context.Context, userId string) (int64, error)
	GetNotesByUserId(ctx context.Context, noteChan chan<- *pb.Note, errChan chan<- error, userId string, limit int64, offset int64)
	GetNoteById(ctx context.Context, id *pb.Id) (*pb.Note, error)
	InsertNote(ctx context.Context, note *pb.Note) (*pb.Note, error)
	UpdateNoteById(ctx context.Context, note *pb.Note) (*pb.Note, error)
	DeleteNoteById(ctx context.Context, id *pb.Id) error
}

type NoteDBImpl struct {
	db *system.Storage
}

func NewNoteDBImpl(db *system.Storage) NoteDB {
	return &NoteDBImpl{db: db}
}

func noteMap(note *pb.Note) []interface{} {
	return []interface{}{
		&note.Id,
		&note.Created,
		&note.Updated,
		&note.Deleted,
		&note.UserId,
		&note.Title,
		&note.Content,
	}
}

func (n *NoteDBImpl) CountNotesByUserId(ctx context.Context, userId string) (int64, error) {
	row := n.db.Conn.QueryRowContext(ctx, "select count(*) from notes where user_id = ?", userId)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (n *NoteDBImpl) GetNotesByUserId(ctx context.Context, noteChan chan<- *pb.Note, errChan chan<- error, userId string, limit int64, offset int64) {
	rows, err := n.db.Conn.QueryContext(ctx, "select * from notes where user_id = ? limit ? offset ?", userId, limit, offset)
	if err != nil {
		errChan <- err
		return
	}
	defer rows.Close()
	for rows.Next() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
			return
		default:
		}

		var note pb.Note
		err := rows.Scan(noteMap(&note)...)
		if err != nil {
			errChan <- err
			return
		}
		noteChan <- &note
	}
	err = rows.Err()
	if err != nil {
		errChan <- err
		return
	}
}

func (n *NoteDBImpl) GetNoteById(ctx context.Context, id *pb.Id) (*pb.Note, error) {
	row := n.db.Conn.QueryRowContext(ctx, "select * from notes where id = ?", id.Id)
	var note pb.Note
	err := row.Scan(noteMap(&note)...)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (n *NoteDBImpl) InsertNote(ctx context.Context, note *pb.Note) (*pb.Note, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	row := n.db.Conn.QueryRowContext(ctx, "insert into notes (id, user_id, title, content) values (?, ?, ?, ?) returning *", id, note.UserId, note.Title, note.Content)
	err = row.Scan(noteMap(note)...)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (n *NoteDBImpl) UpdateNoteById(ctx context.Context, note *pb.Note) (*pb.Note, error) {
	row := n.db.Conn.QueryRowContext(ctx, "update notes set title = ?, content = ? where id = ? returning *", note.Title, note.Content, note.Id)
	err := row.Scan(noteMap(note)...)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (n *NoteDBImpl) DeleteNoteById(ctx context.Context, id *pb.Id) error {
	_, err := n.db.Conn.ExecContext(ctx, "delete from notes where id = ?", id.Id)
	return err
}
