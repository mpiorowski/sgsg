package base

import (
	"errors"
	pb "go-svelte-grpc/grpc"
	"log"
	"net/http"
	"time"

	"github.com/dvsekhvalnov/jose2go"
	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) (*pb.User, error) {

	cookie, err := c.Cookie("sessionCookie")
	if err != nil {
		log.Printf("c.Cookie: %v", err)
		return nil, errors.New("Unauthorized")
	}
	log.Printf("token: %v", cookie)

	var tokenString = cookie
	var secret = []byte(SECRET)

	payload, headers, err := jose.Decode(tokenString, secret)
	if err != nil {
		log.Printf("jose.Decode: %v", err)
		return nil, errors.New("Unauthorized")
	}

	if err == nil {
		//go use token
		log.Printf("\npayload = %v\n", payload)

		//and/or use headers
		log.Printf("\nheaders = %v\n", headers)
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	return nil, errors.New("Unauthorized")

	var email = ""
	var proviedId = ""

	// Make a gRPC request.
	ctx, err, cancel := CreateContext(ENV, URI_USERS)
	if err != nil {
		GrpcError(c, err, "base.CreateContext")
		return nil, err
	}
	defer cancel()
	user, err := UsersGrpcClient.Auth(ctx, &pb.AuthRequest{
		Email:      email,
		ProviderId: proviedId,
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
