package users

import (
	"fmt"
	"sgsg/db"
	"sgsg/utils"
	"time"
)

type Token struct {
	Id           string     `json:"id"`
	Created      time.Time  `json:"created"`
	Updated      time.Time  `json:"updated"`
	Deleted      *time.Time `json:"deleted"`
	UserId       string     `json:"user_id"`
	Provider     string     `json:"provider"`
	AccessToken  string     `json:"access_token"`
	RefreshToken string     `json:"refresh_token"`
	TokenType    string     `json:"token_type"`
	Expires      time.Time  `json:"expires"`
}

func getToken(id string) (*Token, error) {
	row := db.Db.QueryRow("select * from tokens where id = $1", id)
	token, err := scanToken(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanToken: %w", err)
	}
	return token, nil
}

func createToken(t Token) (*Token, error) {
	// delete old tokens
	_, err := db.Db.Exec("delete from tokens where user_id = $1", t.UserId)
	if err != nil {
		return nil, fmt.Errorf("Db.Exec: %w", err)
	}

	// create new token
	id, err := utils.GenerateRandomString(32)
	if err != nil {
		return nil, fmt.Errorf("GenerateRandomString: %w", err)
	}
	row := db.Db.QueryRow("insert into tokens (id, user_id, provider, access_token, refresh_token, token_type, expires) values ($1, $2, $3, $4, $5, $6, $7) returning *",
		id, t.UserId, t.Provider, t.AccessToken, t.RefreshToken, t.TokenType, t.Expires)
	token, err := scanToken(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanToken: %w", err)
	}
	return token, nil
}
