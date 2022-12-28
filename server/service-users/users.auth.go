package main

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/mpiorowski/go-svelte-grpc/server/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserIds struct {
	Uid    string `json:"uid"`
	UserId string `json:"userId"`
}

/**
* Authorize user, user must exsit in database
*/
func (s *server) AuthUser(c context.Context, in *pb.AuthUserRequest) (*pb.AuthUserResponse, error) {
	var userIds UserIds
	var user *pb.User
	var row *sql.Row
	var err error

	row = db.QueryRow(`select uid, "userId" from uids where uid = $1`, in.Uid)
	err = row.Scan(&userIds.Uid, &userIds.UserId)
	if err != nil {
		log.Printf("row.Scan: %v", err)
		return nil, err
	}

	row = db.QueryRow(`update users set "lastLogin" = now() where id = $1 returning *`, userIds.UserId)
	user, err = mapUser(nil, row)
	if err != nil {
		log.Printf("mapUser: %v", err)
		return nil, err
	}
	return &pb.AuthUserResponse{User: user}, nil
}

/**
* Add / update user after successful firebase authentication
*/
// TODO - change code to uid
func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.Session, error) {

	rules := map[string]string{
		"Email": "required,max=100,email",
		"Code":  "required,max=100",
	}
	validate.RegisterStructValidationMapRules(rules, pb.LoginRequest{})
	err := validate.Struct(in)
	if err != nil {
		log.Printf("validate.Struct: %v", err)
		return nil, status.Error(codes.InvalidArgument, "Invalid email or code")
	}

	row := db.QueryRow(`update users set "lastLogin" = now() where $1 = any(email) returning *`, in.Email)
	user, err := mapUser(nil, row)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("mapUser: %v", err)
		return nil, err
	}
	if err == sql.ErrNoRows {
        // TODO - dynamic user role creation
		row = db.QueryRow(`insert into users (email, role) values (array[$1], $2) returning *`, in.Email, "ROLE_USER")
		user, err = mapUser(nil, row)
		if err != nil {
			log.Printf("db.Query: %v", err)
			return nil, err
		}
	}

	_, err = db.Exec(`
        insert into uids (uid, "userId") values ($1, $2) 
        on conflict (uid) do update set "userId" = $2
    `, in.Code, user.Id)
	if err != nil {
		log.Printf("db.Exec: %v", err)
		return nil, err
	}

	return &pb.Session{Role: user.Role, UserId: user.Id, Email: user.Email}, nil
}
