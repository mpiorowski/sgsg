package base

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	pb "github.com/mpiorowski/go-svelte-grpc/server/grpc"
	utils "github.com/mpiorowski/golang"
)

func Auth(c *gin.Context) {
	session, err := Authorization(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, session)
}

/**
* Login user, return user id using firebase uid, create user if not exists
 */
func Login(c *gin.Context) {
	token, client, err := GetFirebaseToken(c)
	if err != nil {
		log.Printf("GetFirebaseToken: %v", err)
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
		ProviderId: token.UID,
		Email:      user.Email,
	})

	if err != nil {
		GrpcError(c, err, "service.Login")
		return
	}
	c.JSON(http.StatusOK, response)
}
