package main

import (
	"database/sql"

	pb "go-svelte-grpc/proto"
)

func mapUser(rows *sql.Rows, row *sql.Row) (*pb.User, error) {
	var user pb.User
	var err error
	if rows != nil {
		err = rows.Scan(
			&user.Id, &user.Created, &user.Updated, &user.Deleted, &user.Email, &user.Role, &user.ProviderId,
		)
	} else if row != nil {
		err = row.Scan(
			&user.Id, &user.Created, &user.Updated, &user.Deleted, &user.Email, &user.Role, &user.ProviderId,
		)
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
