package base

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"

	"firebase.google.com/go/v4/auth"
	utils "github.com/mpiorowski/golang"
	pb "go-svelte-grpc/grpc"
)

var (
	PORT       = utils.MustGetenv("PORT")
	ENV        = utils.MustGetenv("ENV")
	DOMAIN     = utils.MustGetenv("DOMAIN")
	PROJECT_ID = utils.MustGetenv("PROJECT_ID")
	URI_USERS  = utils.MustGetenv("URI_USERS")
	URI_NOTES  = utils.MustGetenv("URI_NOTES")
	URI_FILES  = utils.MustGetenv("URI_FILES")
)

var Client *auth.Client
var Ctx = context.Background()

func GrpcError(c *gin.Context, err error, message string) {
	log.Printf(message+": %v", err)
	s, ok := status.FromError(err)
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": s.Message(), "code": s.Code().String()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong", "code": "BAD_REQUEST"})
}

func GatewayError(c *gin.Context, err string, message string) {
	log.Printf(message+": %v", err)
	c.JSON(http.StatusBadRequest, gin.H{"error": err, "code": "BAD_REQUEST"})
}

func GetFirebaseToken(c *gin.Context) (*auth.Token, *auth.Client, error) {
	sessionCookie, err := c.Cookie("sessionCookie")
	if err != nil {
		return nil, nil, errors.New("sessionCookie is empty")
	}
	if err != nil {
		return nil, nil, err
	}
	token, err := Client.VerifySessionCookie(Ctx, sessionCookie)
	if err != nil {
		return nil, nil, err
	}
	return token, Client, nil
}

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

	// Connect to gRPC server.
	conn, err, ctx, cancel := utils.Connect(ENV, URI_USERS)
	if err != nil {
		log.Printf("utils.Connect: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	// Make a gRPC request.
	service := pb.NewUsersServiceClient(conn)
	user, err := service.Auth(ctx, &pb.AuthRequest{
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
