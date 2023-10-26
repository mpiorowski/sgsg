package users

import (
	"fmt"
	"sgsg/db"

	pb "sgsg/proto"

	"github.com/google/uuid"
)

func selectUserById(id string) (*pb.User, error) {
	row := db.Db.QueryRow("update users set updated = current_timestamp where id = $1 returning *", id)
	user, err := scanUser(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanUser: %w", err)
	}

	return user, nil
}

func selectUserByEmailAndSub(email string, sub string) (*pb.User, error) {
	row := db.Db.QueryRow("select * from users where email = $1 and sub = $2", email, sub)
	user, err := scanUser(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanUser: %w", err)
	}

	return user, nil
}

func insertUser(email string, sub string, avatar string) (*pb.User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("uuid.NewRandom: %w", err)
	}
	row := db.Db.QueryRow("insert into users (id, email, sub, role, avatar, subscription_id, subscription_end) values ($1, $2, $3, $4, $5, $6, $7) returning *",
		id, email, sub, pb.UserRole_ROLE_USER, avatar, "", nil)
	user, err := scanUser(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanUser: %w", err)
	}

	return user, nil
}
