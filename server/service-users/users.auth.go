package main

import (
	"context"
	"database/sql"
	"log"

	pb "go-svelte-grpc/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/**
* Check if user exists, if not create new user
 */
func (s *server) Auth(ctx context.Context, in *pb.AuthRequest) (*pb.User, error) {

	rules := map[string]string{
		"Email":      "required,max=100,email",
		"ProviderId": "required,max=100",
	}
	validate.RegisterStructValidationMapRules(rules, pb.AuthRequest{})
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
