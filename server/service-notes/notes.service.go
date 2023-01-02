package main

import (
	"context"
	"database/sql"
	"log"

	utils "github.com/mpiorowski/golang"
	pb "go-svelte-grpc/grpc"
)

func (s *server) GetNotes(in *pb.UserId, stream pb.NotesService_GetNotesServer) error {
	rows, err := db.Query(`select * from notes where "userId" = $1 and deleted is null`, in.UserId)
	if err != nil {
		log.Printf("db.Query: %v", err)
		return err
	}
	defer rows.Close()

	// Connect to gRPC server.
	conn, err, ctx, cancel := utils.Connect(ENV, URI_USERS)
	if err != nil {
        log.Printf("utils.Connect: %v", err)
		return err
	}
	defer conn.Close()
	defer cancel()

	// Make a gRPC request.
	service := pb.NewUsersServiceClient(conn)


	for rows.Next() {
		note, err := mapNote(rows, nil)
		if err != nil {
			log.Printf("mapNote: %v", err)
			return err
		}

        // Get user for note
	    user, err := service.GetUser(ctx, &pb.UserId{UserId: in.UserId})
        if err != nil {
            log.Printf("service.GetUser: %v", err)
            return err
        }
        note.User = user

		err = stream.Send(note)
		if err != nil {
			log.Printf("stream.Send: %v", err)
			return err
		}
	}
	if rows.Err() != nil {
		log.Printf("rows.Err: %v", err)
		return err
	}
	return nil
}

func (s *server) CreateNote(ctx context.Context, in *pb.Note) (*pb.Note, error) {
    var row *sql.Row
    if in.Id == "" {
        row = db.QueryRow(`insert into notes ("userId", title, content) values ($1, $2, $3) returning *`, in.UserId, in.Title, in.Content)
    } else {
        row = db.QueryRow(`update notes set title = $1, content = $2 where id = $3 and "userId" = $4 returning *`, in.Title, in.Content, in.Id, in.UserId)
    }
    note, err := mapNote(nil, row)
    if err != nil {
        log.Printf("mapNote: %v", err)
        return nil, err
    }
    return note, nil
}

func (s *server) DeleteNote(ctx context.Context, in *pb.NoteId) (*pb.Note, error) {
	row := db.QueryRow(`update notes set deleted = now() where id = $1 and "userId" = $2 returning *`, in.NoteId, in.UserId)
	note, err := mapNote(nil, row)
	if err != nil {
        log.Printf("mapNote: %v", err)
		return nil, err
	}
	return note, nil
}

