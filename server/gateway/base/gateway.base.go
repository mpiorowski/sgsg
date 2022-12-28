package base

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/grpc/status"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	utils "github.com/mpiorowski/golang"
	pb "github.com/mpiorowski/go-svelte-grpc/server/grpc"
)

var (
	PORT       = utils.MustGetenv("PORT")
	ENV        = utils.MustGetenv("ENV")
	PROJECT_ID = utils.MustGetenv("PROJECT_ID")
	DOMAIN     = utils.MustGetenv("DOMAIN")
	URI_FILES  = utils.MustGetenv("URI_FILES")
	URI_USER   = utils.MustGetenv("URI_USER")
	URI_BILLS  = utils.MustGetenv("URI_BILLS")
	URI_FOOD   = utils.MustGetenv("URI_FOOD")
	URI_ENERGY = utils.MustGetenv("URI_ENERGY")
)

var ctx = context.Background()

func GrpcError(c *gin.Context, err error, message string) {
	log.Printf(message+": %v", err)
	s, ok := status.FromError(err)
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": s.Message(), "code": s.Code().String()})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func GatewayError(c *gin.Context, err string, message string) {
    log.Printf(message+": %v", err)
	c.JSON(http.StatusBadRequest, gin.H{"error": err, "code": "BAD_REQUEST"})
}

func ConnectToFirebase() (*auth.Client, error) {
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetFirebaseToken(c *gin.Context) (*auth.Token, *auth.Client, error) {
	idToken := c.GetHeader("Authorization")
	idToken = idToken[7:]

	if idToken == "" {
		return nil, nil, errors.New("idToken is empty")
	}
	client, err := ConnectToFirebase()
	if err != nil {
		return nil, nil, err
	}
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, nil, err
	}
	return token, client, nil
}

func Authorization(c *gin.Context) (*pb.Session, error) {

	token, _, err := GetFirebaseToken(c)
	if err != nil {
		log.Printf("GetFirebaseToken: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, err
	}

	// Connect to gRPC server.
	conn, err, ctx, cancel := utils.Connect(ENV, URI_USER)
	if err != nil {
		log.Printf("utils.Connect: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, err
	}
	defer conn.Close()
	defer cancel()

	// Make a gRPC request.
	service := pb.NewUsersServiceClient(conn)
	user, err := service.AuthUser(ctx, &pb.AuthUserRequest{
		Uid: token.UID,
	})

	if err != nil || user.GetUser().GetId() == "" {
		log.Printf("service.AuthUser: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil, err
	}

	var session pb.Session
	session.UserId = user.GetUser().GetId()
	session.Email = user.GetUser().GetEmail()
	session.Role = user.GetUser().GetRole()

	return &session, nil
}
