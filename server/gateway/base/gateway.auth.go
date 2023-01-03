package base

import (
	"errors"
	pb "go-svelte-grpc/grpc"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) (*pb.User, error) {

	token, _, err := GetFirebaseToken(c)
	if err != nil || token == nil {
		log.Printf("GetFirebaseToken: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, err
	}
	var email string
	if token.Claims["email"] != nil {
		email = token.Claims["email"].(string)
	} else {
		log.Printf("token.Claims[email] is empty")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, errors.New("token.Claims[email] is empty")
	}

	// Make a gRPC request.
    ctx, err, cancel := CreateContext(ENV, URI_USERS)
    if err != nil {
        GrpcError(c, err, "base.CreateContext")
        return nil, err
    }
    defer cancel()
	user, err := UsersGrpcClient.Auth(ctx, &pb.AuthRequest{
		Email:      email,
		ProviderId: token.UID,
	})

	if err != nil || user.GetId() == "" {
		log.Printf("service.AuthUser: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, err
	}

	return user, nil
}

func Auth(c *gin.Context) {
	session, err := Authorization(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, session)
}

type IdToken struct {
	IdToken string `json:"idToken"`
}

func Login(c *gin.Context) {

	var idToken IdToken
	err := c.BindJSON(&idToken)
	if err != nil {
		log.Printf("c.BindJSON: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid idToken"})
		return
	}

	if err != nil {
		log.Printf("ConnectToFirebase: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Set session expiration to 5 days.
	expiresIn := time.Hour * 24 * 5
	cookie, err := Client.SessionCookie(c, idToken.IdToken, expiresIn)
	if err != nil {
		log.Printf("client.SessionCookie: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid idToken"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cookie": cookie})
}
