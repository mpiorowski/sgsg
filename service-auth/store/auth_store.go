package store

import (
	"context"
	"service-auth/system"

	pb "service-auth/proto"

	"github.com/google/uuid"
)

type Token struct {
	Id       string
	Expires  string
	UserId   string
	State    string
	Verifier string
}

type AuthDB struct {
	*system.Storage
}

func NewAuthDB(s *system.Storage) AuthDB {
	return AuthDB{s}
}

func destT(token *Token) []interface{} {
	return []interface{}{
		&token.Id,
		&token.Expires,
		&token.UserId,
		&token.State,
		&token.Verifier,
	}
}

func dest(user *pb.User) []interface{} {
	return []interface{}{
		&user.Id,
		&user.Created,
		&user.Updated,
		&user.Deleted,
		&user.Email,
		&user.Role,
		&user.Sub,
		&user.Avatar,
		&user.SubscriptionId,
		&user.SubscriptionEnd,
		&user.SubscriptionCheck,
	}
}

func (db AuthDB) SelectTokenById(
	ctx context.Context,
	id string,
) (*Token, error) {
	row := db.Conn.QueryRowContext(ctx, "select * from tokens where id = ?", id)
	var token Token
	err := row.Scan(destT(&token)...)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (db AuthDB) SeleteTokenByState(
	ctx context.Context,
	state string,
) (*Token, error) {
	row := db.Conn.QueryRowContext(ctx, "select * from tokens where state = ?", state)
	var token Token
	err := row.Scan(destT(&token)...)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (db AuthDB) InsertToken(
	ctx context.Context,
	expires string,
	userId string,
	state string,
	verifier string,
) (*Token, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	row := db.Conn.QueryRow("insert into tokens (id, expires, user_id, state, verifier) values (?, ?, ?, ?, ?) returning *",
		id, expires, userId, state, verifier)
	var token Token
	err = row.Scan(destT(&token)...)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (db AuthDB) UpdateToken(ctx context.Context, id string, expires string) error {
	_, err := db.Conn.Exec("update tokens set expires = ? where id = ?", expires, id)
	if err != nil {
		return err
	}
	return nil
}

func (db AuthDB) CleanTokens(
	ctx context.Context,
) error {
	_, err := db.Conn.ExecContext(ctx, "delete from tokens where expires < current_timestamp")
	if err != nil {
		return err
	}
	return nil
}

func (db AuthDB) SelectUserById(
	ctx context.Context,
	id string,
) (*pb.User, error) {
	row := db.Conn.QueryRowContext(ctx, "select * from users where id = ?", id)
	var user pb.User
	err := row.Scan(dest(&user)...)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db AuthDB) SelectUserByTokenId(
	ctx context.Context,
	tokenId string,
) (*pb.User, error) {
	row := db.Conn.QueryRowContext(ctx, "select * from users where id = (select user_id from tokens where id = ?)", tokenId)
	var user pb.User
	err := row.Scan(dest(&user)...)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db AuthDB) SelectUserByEmail(
	ctx context.Context,
	email string,
) (*pb.User, error) {
	row := db.Conn.QueryRowContext(ctx, "select * from users where email = ?", email)
	var user pb.User
	err := row.Scan(dest(&user)...)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db AuthDB) InsertUser(
	ctx context.Context,
	email string,
	sub string,
	avatar string,
) (*pb.User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	row := db.Conn.QueryRowContext(ctx, "insert into users (id, email, sub, role, avatar) values (?, ?, ?, ?, ?) returning *",
		id, email, sub, pb.Role_ROLE_USER, avatar)
	var user pb.User
	err = row.Scan(dest(&user)...)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db AuthDB) UpdateUserSub(
	ctx context.Context,
	id string,
	sub string,
) error {
	_, err := db.Conn.ExecContext(ctx, "update users set sub = ? where id = ?", sub, id)
	if err != nil {
		return err
	}
	return nil
}

func (db AuthDB) UpdateUserActivity(ctx context.Context, id string) error {
	_, err := db.Conn.ExecContext(ctx, "update users set updated = current_timestamp where id = ? returning *", id)
	if err != nil {
		return err
	}
	return nil
}

func (db AuthDB) UpdateSubscriptionId(
	ctx context.Context,
	userId string,
	subscriptionId string,
) error {
	_, err := db.Conn.ExecContext(ctx, "update users set subscription_id = ? where id = ?", subscriptionId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (db AuthDB) UpdateSubscriptionCheck(
	ctx context.Context,
	userId string,
	subscriptionCheck string,
) error {
	_, err := db.Conn.ExecContext(ctx, "update users set subscription_check = ? where id = ?", subscriptionCheck, userId)
	if err != nil {
		return err
	}
	return nil
}

func (db AuthDB) UpdateSubscriptionEnd(
	ctx context.Context,
	userId string,
	subscriptionEnd string,
) error {
	_, err := db.Conn.ExecContext(ctx, "update users set subscription_end = ? where id = ?", subscriptionEnd, userId)
	if err != nil {
		return err
	}
	return nil
}
