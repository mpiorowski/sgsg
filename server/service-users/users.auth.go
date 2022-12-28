package main

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/mpiorowski/go-svelte-grpc/server/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/**
* Authorize user
 */
func (s *server) AuthUser(c context.Context, in *pb.AuthRequest) (*pb.User, error) {
	var user *pb.User

	row := db.QueryRow(`select * from users where "providerId" = $1 and deleted is null`, in.ProviderId)
    user, err := mapUser(nil, row)
    if err != nil {
        log.Printf("mapUser: %v", err)
        return nil, status.Errorf(codes.NotFound, "User not found")
    }

	return user, nil
}

/**
* Check if user exists, if not create new user
 */
func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.User, error) {

	rules := map[string]string{
		"Email": "required,max=100,email",
		"Id":    "required,max=100",
	}
	validate.RegisterStructValidationMapRules(rules, pb.LoginRequest{})
	err := validate.Struct(in)
	if err != nil {
		log.Printf("validate.Struct: %v", err)
		return nil, status.Error(codes.InvalidArgument, "Invalid email or code")
	}

	row := db.QueryRow(`select * from users where email = $1 and "providerId" = $2`, in.Email, in.ProviderId)
	user, err := mapUser(nil, row)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("mapUser: %v", err)
		return nil, err
	}

    if user.GetDeleted() != "" {
        return nil, status.Error(codes.Unauthenticated, "User is deleted")
    }

	if err == sql.ErrNoRows {
		// TODO - dynamic role assign
		row = db.QueryRow(`insert into users (email, role, "providerId") values ($1, $2, $3) returning *`, in.Email, pb.UserRole_ROLE_USER.String(), in.ProviderId)
		user, err = mapUser(nil, row)
		if err != nil {
			log.Printf("mapUser: %v", err)
			return nil, err
		}
	}

	return user, nil
}
