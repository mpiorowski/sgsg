package users

import (
	"database/sql"
	"fmt"
	"sgsg/db"

	pb "sgsg/proto"

	"github.com/google/uuid"
)

func getUserById(id string) (*pb.User, error) {
    row := db.Db.QueryRow("update users set updated = current_timestamp where id = $1 returning *", id)
    user, err := scanUser(nil, row)
    if err != nil {
        return nil, fmt.Errorf("scanUser: %w", err)
    }

    return user, nil
}

func createUser(email string, sub string, avatar string) (*pb.User, error) {

	// check if user exists
	row := db.Db.QueryRow("select * from users where email = $1 and sub = $2", email, sub)
	user, err := scanUser(nil, row)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("scanUser: %w", err)
	}
	if user != nil {
		return user, nil
	}

	// create new user
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("uuid.NewRandom: %w", err)
	}
	// apiKey, err := utils.GenerateRandomUuids()
	// apiKey = strings.Replace(apiKey, "-", "", -1)
	// apiKey = strings.Join([]string{"sk_live_", apiKey}, "")
	// if err != nil {
	// 	return nil, err
	// }
	row = db.Db.QueryRow("insert into users (id, email, sub, role, avatar, subscription_id, subscription_end) values ($1, $2, $3, $4, $5, $6, $7) returning *",
		id, email, sub, pb.UserRole_ROLE_USER, avatar, "", nil)
	user, err = scanUser(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanUser: %w", err)
	}

	return user, nil
}
