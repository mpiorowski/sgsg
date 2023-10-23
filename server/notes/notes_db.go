package notes

import (
	"database/sql"
	"fmt"
	"sgsg/db"
	pb "sgsg/proto"

	"github.com/google/uuid"
)

func selectNotesStream(userId string) (*sql.Rows, error) {
	return db.Db.Query("select * from notes where user_id = $1", userId)
}

func selectNoteById(id string, userId string) (*pb.Note, error) {
    row := db.Db.QueryRow("select * from notes where id = $1 and user_id = $2", id, userId)
    note, err := scanNote(nil, row)
    if err != nil {
        return nil, fmt.Errorf("scanNote: %w", err)
    }
    return note, nil
}

func insertNote(note *pb.Note) (*pb.Note, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("uuid.NewRandom: %w", err)
	}
	row := db.Db.QueryRow(
		"insert into notes (id, user_id, title, content) values ($1, $2, $3, $4) returning *",
		id,
		note.UserId,
		note.Title,
		note.Content,
	)
	note, err = scanNote(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanNote: %w", err)
	}
	return note, nil
}

func updateNote(note *pb.Note) (*pb.Note, error) {
	row := db.Db.QueryRow(
		"update notes set title = $1, content = $2 where id = $3 and user_id = $4 returning *",
		note.Title,
		note.Content,
		note.Id,
        note.UserId,
	)
	note, err := scanNote(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanNote: %w", err)
	}
	return note, nil
}

func deleteNoteById(id string) error {
	_, err := db.Db.Exec("delete from notes where id = $1", id)
	if err != nil {
		return fmt.Errorf("db.Exec: %w", err)
	}
	return nil
}
