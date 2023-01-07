package main

import (
	"context"
	"database/sql"
	"io"
	"log"

	pb "go-svelte-grpc/proto"

	utils "github.com/mpiorowski/golang"
)

func (s *server) GetNotes(in *pb.UserId, stream pb.NotesService_GetNotesServer) error {
	log.Println("GetNotes")

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

	// Steam user for each note.
	streamUser, err := service.GetUsersByIds(ctx)
	if err != nil {
		log.Printf("GetUsersByIds: %v", err)
		return err
	}

	notesArray := []*pb.Note{}

	for rows.Next() {
		note, err := mapNote(rows, nil)
		if err != nil {
			log.Printf("mapNote: %v", err)
			return err
		}
		notesArray = append(notesArray, note)

		err = streamUser.Send(&pb.UserId{UserId: note.UserId})
		user, err := streamUser.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("streamUser.Recv: %v", err)
			return err
		}
		note.User = user
		stream.Send(note)
	}

	// This is an asynchronus stream, but doesn't fit here, cos it will send all notes connected to user each time
	// Left here for future reference
	// waitc := make(chan struct{})
	// go func() {
	// 	for {
	// 		user, err := streamUser.Recv()
	// 		if err == io.EOF {
	// 			// read done.
	// 			close(waitc)
	// 			return
	// 		}
	// 		if err != nil {
	// 			log.Printf("streamUser.Recv: %v", err)
	// 			return
	// 		}
	// 		for _, note := range notesArray {
	// 			if note.UserId == user.Id {
	// 				note.User = user
	// 				stream.Send(note)
	// 			}
	// 		}
	// 	}
	// }()

	err = streamUser.CloseSend()
	// <-waitc
	if err != nil {
		log.Printf("streamUser.CloseSend: %v", err)
		return err
	}

	if rows.Err() != nil {
		log.Printf("rows.Err: %v", err)
		return err
	}
	return nil
}

func (s *server) CreateNote(ctx context.Context, in *pb.Note) (*pb.Note, error) {
	log.Println("CreateNote")

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
	log.Println("DeleteNote")

	row := db.QueryRow(`update notes set deleted = now() where id = $1 and "userId" = $2 returning *`, in.NoteId, in.UserId)
	note, err := mapNote(nil, row)
	if err != nil {
		log.Printf("mapNote: %v", err)
		return nil, err
	}
	return note, nil
}
