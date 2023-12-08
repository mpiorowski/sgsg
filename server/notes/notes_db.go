package notes

import (
	"database/sql"
	"fmt"
	"sgsg/db"
	pb "sgsg/proto"

	"github.com/google/uuid"
)

func scanNote(rows *sql.Rows, row *sql.Row) (*pb.Note, error) {
	note := pb.Note{}
	if rows != nil {
		err := rows.Scan(&note.Id, &note.Created, &note.Updated, &note.Deleted, &note.UserId, &note.Title, &note.Content)
		if err != nil {
			return nil, err
		}
	} else {
		err := row.Scan(&note.Id, &note.Created, &note.Updated, &note.Deleted, &note.UserId, &note.Title, &note.Content)
		if err != nil {
			return nil, err
		}
	}
	return &note, nil
}

func selectNotes(notesCh chan<- *pb.Note, errCh chan<- error, userId string) {
    rows, err := db.Db.Query("select * from notes where user_id = $1", userId)
    if err != nil {
        errCh <- fmt.Errorf("db.Query: %w", err)
        return
    }
    defer rows.Close()
    for rows.Next() {
        note, err := scanNote(rows, nil)
        if err != nil {
            errCh <- fmt.Errorf("scanNote: %w", err)
            return
        }
        notesCh <- note
    }
    if err := rows.Err(); err != nil {
        errCh <- fmt.Errorf("rows.Err: %w", err)
        return
    }
    close(notesCh)
}

func selectNoteById(id string, userId string) (*pb.Note, error) {
	row := db.Db.QueryRow("select * from notes where id = $1 and user_id = $2", id, userId)
	note, err := scanNote(nil, row)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("scanNote: %w", err)
	}
    if err == sql.ErrNoRows {
        return &pb.Note{}, nil
    }
	return note, nil
}

func insertNote(in *pb.Note) (*pb.Note, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("uuid.NewRandom: %w", err)
	}
	row := db.Db.QueryRow(
		"insert into notes (id, user_id, title, content) values ($1, $2, $3, $4) returning *",
		id,
		in.UserId,
		in.Title,
		in.Content,
	)
	note, err := scanNote(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanNote: %w", err)
	}
	return note, nil
}

func updateNote(in *pb.Note) (*pb.Note, error) {
	row := db.Db.QueryRow(
		"update notes set title = $1, content = $2 where id = $3 and user_id = $4 returning *",
		in.Title,
		in.Content,
		in.Id,
		in.UserId,
	)
	note, err := scanNote(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanNote: %w", err)
	}
	return note, nil
}

func deleteNoteById(id string) error {
	err := db.Db.QueryRow("delete from notes where id = $1 returning id", id).Scan(&id)
	if err != nil {
		return fmt.Errorf("db.Exec: %w", err)
	}
	return nil
}
