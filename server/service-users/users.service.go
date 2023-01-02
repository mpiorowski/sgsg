package main

import (
	"context"
	"io"
	"log"

	pb "go-svelte-grpc/grpc"
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

func (s *server) GetUser(ctx context.Context, in *pb.UserId) (*pb.User, error) {
	row := db.QueryRow(`select * from users where id = $1`, in.UserId)
	user, err := mapUser(nil, row)
	if err != nil {
		log.Printf("mapUser: %v", err)
		return nil, err
	}
	return user, nil
}

func (s *server) DeleteUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	row := db.QueryRow(`update users set deleted = now() where id = $1 and "providerId" = $2 and email = $3 returning *`, in.Id, in.ProviderId, in.Email)
	user, err := mapUser(nil, row)
	if err != nil {
		log.Printf("db.Exec: %v", err)
		return nil, err
	}
	return user, nil
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
