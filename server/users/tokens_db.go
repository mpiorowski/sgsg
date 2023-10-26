package users

import (
	"database/sql"
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

func scanToken(rows *sql.Rows, row *sql.Row) (*Token, error) {
	token := &Token{}
	if rows != nil {
		err := rows.Scan(&token.Id, &token.Created, &token.Updated, &token.Deleted, &token.UserId, &token.Provider, &token.AccessToken, &token.RefreshToken, &token.TokenType, &token.Expires)
		if err != nil {
			return nil, err
		}
	} else {
		err := row.Scan(&token.Id, &token.Created, &token.Updated, &token.Deleted, &token.UserId, &token.Provider, &token.AccessToken, &token.RefreshToken, &token.TokenType, &token.Expires)
		if err != nil {
			return nil, err
		}
	}
	return token, nil
}

func selectToken(id string) (*Token, error) {
	row := db.Db.QueryRow("select * from tokens where id = $1", id)
	token, err := scanToken(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanToken: %w", err)
	}
	return token, nil
}

func insertToken(t Token) (*Token, error) {
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

func deleteTokensByUserId(userId string) error {
	_, err := db.Db.Exec("delete from tokens where user_id = $1", userId)
	if err != nil {
		return fmt.Errorf("Db.Exec: %w", err)
	}
	return nil
}
