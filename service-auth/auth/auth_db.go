package auth

import (
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

type AuthDB interface {
	selectTokenById(id string) (*Token, error)
	seleteTokenByState(state string) (*Token, error)
	insertToken(expires string, userId string, state string, verifier string) (*Token, error)
	updateToken(id string, expires string) error
    CleanTokens() error
	selectUserById(id string) (*pb.User, error)
    selectUserByTokenId(tokenId string) (*pb.User, error)
	selectUserByEmailAndSub(email string, sub string) (*pb.User, error)
	insertUser(email string, sub string, avatar string) (*pb.User, error)
	updateUser(id string) error
	updateSubscriptionId(userId string, subscriptionId string) error
	updateSubscriptionCheck(userId string, subscriptionCheck string) error
	updateSubscriptionEnd(userId string, subscriptionEnd string) error
}

type AuthDBImpl struct {
	*system.Storage
}

func NewAuthDB(s *system.Storage) AuthDB {
	return &AuthDBImpl{s}
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

func (db *AuthDBImpl) selectTokenById(id string) (*Token, error) {
	row := db.Conn.QueryRow("select * from tokens where id = ?", id)
	var token Token
	err := row.Scan(destT(&token)...)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (db *AuthDBImpl) seleteTokenByState(state string) (*Token, error) {
	row := db.Conn.QueryRow("select * from tokens where state = ?", state)
	var token Token
	err := row.Scan(destT(&token)...)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (db *AuthDBImpl) insertToken(expires string, userId string, state string, verifier string) (*Token, error) {
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

func (db *AuthDBImpl) updateToken(id string, expires string) error {
	_, err := db.Conn.Exec("update tokens set expires = ? where id = ?", expires, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *AuthDBImpl) CleanTokens() error {
    _, err := db.Conn.Exec("delete from tokens where expires < current_timestamp")
    if err != nil {
        return err
    }
    return nil
}

func (db *AuthDBImpl) selectUserById(id string) (*pb.User, error) {
	row := db.Conn.QueryRow("select * from users where id = ?", id)
	var user pb.User
	err := row.Scan(dest(&user)...)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *AuthDBImpl) selectUserByTokenId(tokenId string) (*pb.User, error) {
    row := db.Conn.QueryRow("select * from users where id = (select user_id from tokens where id = ?)", tokenId)
    var user pb.User
    err := row.Scan(dest(&user)...)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (db *AuthDBImpl) selectUserByEmailAndSub(email string, sub string) (*pb.User, error) {
	row := db.Conn.QueryRow("select * from users where email = ? and sub = ?", email, sub)
	var user pb.User
	err := row.Scan(dest(&user)...)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *AuthDBImpl) insertUser(email string, sub string, avatar string) (*pb.User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	row := db.Conn.QueryRow("insert into users (id, email, sub, role, avatar) values (?, ?, ?, ?, ?) returning *",
		id, email, sub, pb.Role_ROLE_USER, avatar)
	var user pb.User
	err = row.Scan(dest(&user)...)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *AuthDBImpl) updateUser(id string) error {
	_, err := db.Conn.Exec("update users set updated = current_timestamp where id = ? returning *", id)
	if err != nil {
		return err
	}
	return nil
}

func (db *AuthDBImpl) updateSubscriptionId(userId string, subscriptionId string) error {
	_, err := db.Conn.Exec("update users set subscription_id = ? where id = ?", subscriptionId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (db *AuthDBImpl) updateSubscriptionCheck(userId string, subscriptionCheck string) error {
	_, err := db.Conn.Exec("update users set subscription_check = ? where id = ?", subscriptionCheck, userId)
	if err != nil {
		return err
	}
	return nil
}

func (db *AuthDBImpl) updateSubscriptionEnd(userId string, subscriptionEnd string) error {
	_, err := db.Conn.Exec("update users set subscription_end = ? where id = ?", subscriptionEnd, userId)
	if err != nil {
		return err
	}
	return nil
}
