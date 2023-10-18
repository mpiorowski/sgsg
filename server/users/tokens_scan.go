package users

import "database/sql"

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
