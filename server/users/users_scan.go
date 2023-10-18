package users

import (
	"database/sql"
	pb "sgsg/proto"
)

func scanUser(rows *sql.Rows, row *sql.Row) (*pb.User, error) {
	user := pb.User{}
	if rows != nil {
		err := rows.Scan(&user.Id, &user.Created, &user.Updated, &user.Deleted, &user.Email, &user.Sub, &user.Role, &user.Avatar, &user.SubscriptionId, &user.SubscriptionEnd)
		if err != nil {
			return nil, err
		}
	} else {
		err := row.Scan(&user.Id, &user.Created, &user.Updated, &user.Deleted, &user.Email, &user.Sub, &user.Role, &user.Avatar, &user.SubscriptionId, &user.SubscriptionEnd)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
