package main

import (
	"io"
	"log"

	pb "github.com/mpiorowski/go-svelte-grpc/server/grpc"
)

func (s *server) GetUsers(in *pb.Empty, stream pb.UsersService_GetUsersServer) error {
	rows, err := db.Query(`select * from users`)
	if err != nil {
		log.Printf("db.Query: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := mapUser(rows, nil)
		if err != nil {
			log.Printf("mapUser: %v", err)
			return err
		}
		err = stream.Send(user)
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

func (s *server) CreateUser(stream pb.UsersService_CreateUserServer) error {
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("stream.Recv: %v", err)
			return err
		}
		rows := db.QueryRow(`insert into users (email, role) values ($1, $2) returning *`, user.Email, user.Role)
		user, err = mapUser(nil, rows)
		if err != nil {
			log.Printf("mapUser: %v", err)
			return err
		}

		err = stream.Send(user)
		if err != nil {
			log.Printf("stream.Send: %v", err)
			return err
		}
	}
}
