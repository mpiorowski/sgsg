package notes

import (
	"database/sql"
	"fmt"
	pb "sgsg/proto"
	"sgsg/system"

	"github.com/google/uuid"
)

type NoteDB interface {
	SelectNotesByUserID(notesCh chan<- *pb.Note, errCh chan<- error, userId string)
	SelectNoteByID(id string, userId string) (*pb.Note, error)
	InsertNote(in *pb.Note) (*pb.Note, error)
	UpdateNote(in *pb.Note) (*pb.Note, error)
	DeleteNoteByID(id string) error
}

type noteDB struct {
	*system.Storage
}

func NewNoteDB(storage *system.Storage) NoteDB {
	return &noteDB{storage}
}

func (db *noteDB) SelectNotesByUserID(notesCh chan<- *pb.Note, errCh chan<- error, userId string) {
	defer close(notesCh)
	rows, err := db.Conn.Query("select * from notes where user_id = $1", userId)
	if err != nil {
		errCh <- fmt.Errorf("db.Query: %w", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var note pb.Note
		err := rows.Scan(dest(&note)...)
		if err != nil {
			errCh <- fmt.Errorf("scanNote: %w", err)
			return
		}
		notesCh <- &note
	}
	err = rows.Err()
	if err != nil {
		errCh <- fmt.Errorf("rows.Err: %w", err)
	}
}

func (db *noteDB) SelectNoteByID(id string, userId string) (*pb.Note, error) {
	row := db.Conn.QueryRow("select * from notes where id = $1 and user_id = $2", id, userId)
	var note pb.Note
	err := row.Scan(dest(&note)...)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("scanNote: %w", err)
	}
	if err == sql.ErrNoRows {
		return &pb.Note{}, nil
	}
	return &note, nil
}

func (db *noteDB) InsertNote(in *pb.Note) (*pb.Note, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("uuid.NewV7: %w", err)
	}
	row := db.Conn.QueryRow(
		"insert into notes (id, user_id, title, content) values ($1, $2, $3, $4) returning *",
		id,
		in.UserId,
		in.Title,
		in.Content,
	)
	var note pb.Note
	err = row.Scan(dest(&note)...)
	if err != nil {
		return nil, fmt.Errorf("scanNote: %w", err)
	}
	return &note, nil
}

func (db *noteDB) UpdateNote(in *pb.Note) (*pb.Note, error) {
	row := db.Conn.QueryRow(
		"update notes set title = $1, content = $2 where id = $3 and user_id = $4 returning *",
		in.Title,
		in.Content,
		in.Id,
		in.UserId,
	)
	var note pb.Note
	err := row.Scan(dest(&note)...)
	if err != nil {
		return nil, fmt.Errorf("scanNote: %w", err)
	}
	return &note, nil
}

func (db *noteDB) DeleteNoteByID(id string) error {
	err := db.Conn.QueryRow("delete from notes where id = $1 returning id", id).Scan(&id)
	if err != nil {
		return fmt.Errorf("db.Exec: %w", err)
	}
	return nil
}

func dest(note *pb.Note) []interface{} {
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
