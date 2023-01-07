package main

import (
	"database/sql"
	pb "go-svelte-grpc/proto"
)

func mapNote(rows *sql.Rows, row *sql.Row) (*pb.Note, error) {
	var note pb.Note
	var err error
	if rows != nil {
		err = rows.Scan(
			&note.Id, &note.Created, &note.Updated, &note.Deleted, &note.UserId, &note.Title, &note.Content,
		)
	} else if row != nil {
		err = row.Scan(
			&note.Id, &note.Created, &note.Updated, &note.Deleted, &note.UserId, &note.Title, &note.Content,
		)
	}
	if err != nil {
		return nil, err
	}
	return &note, nil
}
