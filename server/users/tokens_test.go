package users

import (
	"sgsg/db"
	"testing"
	"time"
)

var token = Token{
	UserId:       "test",
	Provider:     "google",
	AccessToken:  "test",
	RefreshToken: "test",
	TokenType:    "test",
	Expires:      time.Now().Add(time.Hour * 24 * 7),
}

func setup() {
	err := db.ConnectTest()
	if err != nil {
		panic(err)
	}
	err = db.Migrations()
	if err != nil {
		panic(err)
	}
}

func TestInsertToken(t *testing.T) {
    setup()
    newToken, err := insertToken(token)
    if err != nil {
        t.Error(err)
    }
    if newToken.Id == "" {
        t.Error("token.Id is empty")
    }
    if newToken.Created.IsZero() {
        t.Error("token.Created is zero")
    }
    if newToken.AccessToken != token.AccessToken {
        t.Error("token.AccessToken is not equal")
    }
    if newToken.RefreshToken != token.RefreshToken {
        t.Error("token.RefreshToken is not equal")
    }
}
