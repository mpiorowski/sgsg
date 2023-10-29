package users

import (
	"testing"
	"time"
)

var tokens = []Token{
	{
		UserId:       "test",
		Provider:     "google",
		AccessToken:  "test1",
		RefreshToken: "test1",
		TokenType:    "test1",
		Expires:      time.Now().Add(time.Hour * 24 * 7),
	},
	{
		UserId:       "test",
		Provider:     "github",
		AccessToken:  "test2",
		RefreshToken: "test2",
		TokenType:    "test2",
		Expires:      time.Now().Add(time.Hour * 24 * 7),
	},
}


func TestInsertToken(t *testing.T) {
	newToken, err := insertToken(tokens[0])
	if err != nil {
		t.Error(err)
	}
	if newToken.Id == "" {
		t.Error("token.Id is empty")
	}
	if newToken.Created.IsZero() {
		t.Error("token.Created is zero")
	}
	if newToken.AccessToken != tokens[0].AccessToken {
		t.Error("token.AccessToken is not equal")
	}
}

func TestSelectToken(t *testing.T) {
	newToken, err := insertToken(tokens[0])
	if err != nil {
		t.Error(err)
	}
	selectedToken, err := selectTokenById(newToken.Id)
	if err != nil {
		t.Error(err)
	}
	if selectedToken.Id != newToken.Id {
		t.Error("token.Id is not equal")
	}
	if selectedToken.Created != newToken.Created {
		t.Error("token.Created is not equal")
	}
	if selectedToken.AccessToken != newToken.AccessToken {
		t.Error("token.AccessToken is not equal")
	}
}

func TestDeleteTokensByUserId(t *testing.T) {
	token1, _ := insertToken(tokens[0])
    token2, _ := insertToken(tokens[1])
    err := deleteTokensByUserId(token1.UserId)
	if err != nil {
		t.Error(err)
	}
	_, err = selectTokenById(token1.Id)
    if err == nil {
        t.Error("token1 is not deleted")
    }
    _, err = selectTokenById(token2.Id)
    if err == nil {
        t.Error("token2 is not deleted")
    }

}
