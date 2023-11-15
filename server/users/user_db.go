package users

import (
	"database/sql"
	"fmt"
	"sgsg/db"

	pb "sgsg/proto"

	"github.com/google/uuid"
)

func scanUser(rows *sql.Rows, row *sql.Row) (*pb.User, error) {
	user := pb.User{}
	if rows != nil {
		err := rows.Scan(&user.Id, &user.Created, &user.Updated, &user.Deleted, &user.Email, &user.Role, &user.Sub, &user.Avatar, &user.SubscriptionId, &user.SubscriptionEnd)
		if err != nil {
			return nil, err
		}
	} else {
		err := row.Scan(&user.Id, &user.Created, &user.Updated, &user.Deleted, &user.Email, &user.Role, &user.Sub, &user.Avatar, &user.SubscriptionId, &user.SubscriptionEnd)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

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
	row := db.Db.QueryRow("insert into users (id, email, sub, role, avatar) values ($1, $2, $3, $4, $5) returning *",
		id, email, sub, pb.UserRole_ROLE_USER, avatar)
	user, err := scanUser(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanUser: %w", err)
	}

	return user, nil
}
