package base

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go/v4/auth"
	utils "github.com/mpiorowski/golang"
	pb "github.com/mpiorowski/go-svelte-grpc/server/grpc"

)

func Auth(c *gin.Context) {
	session, err := Authorization(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, session)
}

/**
* Create magic link using firebase, only in production
* @param {string} Email
 */
func CreateMagicLink(c *gin.Context) {

	userEmail := c.Param("email")
	if userEmail == "" {
		GatewayError(c, "Email is required", "c.Param")
		return
	}

	client, err := ConnectToFirebase()
	if err != nil {
		log.Printf("base.ConnectToFirebaseAuth: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// TODO - make it generic
	actionCodeSettings := &firebase.ActionCodeSettings{
		URL: "https://www." + DOMAIN + "/auth",
	}

	link, err := client.EmailSignInLink(ctx, userEmail, actionCodeSettings)
	if err != nil {
		log.Printf("client.EmailSignInLink: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// TODO - make template variable an enum
	var email utils.Email = utils.Email{
		To:       userEmail,
		Html:     []string{link},
		Template: "FIREBASE_MAGIC_LINK",
	}
	out, err := json.Marshal(email)
	if err != nil {
		log.Printf("json.Marshal: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Send email only in production
	if ENV == "production" {
		err = utils.PublishPubSub(PROJECT_ID, "email", string(out))
		if err != nil {
			log.Printf("utils.PublishPubSub: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": "email send"})
}

/**
* Login user, return user id using firebase uid, create user if not exists
 */
// TODO - clean up LoginRequest type
func Login(c *gin.Context) {
	token, client, err := GetFirebaseToken(c)
	if err != nil {
		log.Printf("base.ConnectToFirebase: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := client.GetUser(ctx, token.UID)
	if err != nil {
		log.Printf("client.GetUser: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Connect to gRPC server.
	conn, err, ctx, cancel := utils.Connect(ENV, URI_USER)
	if err != nil {
		GrpcError(c, err, "utils.Connect")
		return
	}
	defer conn.Close()
	defer cancel()

	// Make a gRPC request
	service := pb.NewUsersServiceClient(conn)
	response, err := service.Login(ctx, &pb.LoginRequest{
		Email: user.Email,
		Code:  token.UID,
	})

	if err != nil {
		GrpcError(c, err, "service.Login")
		return
	}
	c.JSON(http.StatusOK, response)
}
