package notes

import (
	"database/sql"
	pb "sgsg/proto"
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
