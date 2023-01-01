package users

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "go-svelte-grpc/server/grpc"
	utils "github.com/mpiorowski/golang"

	base "go-svelte-grpc/server/gateway/base"
)

func GetUsers(c *gin.Context) {

	user, err := base.Authorization(c)
	if err != nil {
		return
	}
	// TODO - change to admin
	if user.GetRole() != pb.UserRole_ROLE_USER.String() {
		log.Printf("Unauthorized user: %v", user.GetEmail())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Connect to gRPC server.
	conn, err, ctx, cancel := utils.Connect(base.ENV, base.URI_USER)
	if err != nil {
		base.GrpcError(c, err, "utils.Connect")
		return
	}
	defer conn.Close()
	defer cancel()

	// Make a gRPC request.
	service := pb.NewUsersServiceClient(conn)
	stream, err := service.GetUsers(ctx, &pb.Empty{})

	if err != nil {
		base.GrpcError(c, err, "service.GetUsers")
		return
	}

	var users []*pb.User
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			base.GrpcError(c, err, "stream.Recv")
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {

	var request pb.User

	err := c.BindJSON(&request)
	if err != nil {
		base.GrpcError(c, err, "c.BindJSON")
		return
	}

	// Connect to gRPC server.
	conn, err, ctx, cancel := utils.Connect(base.ENV, base.URI_USER)
	if err != nil {
		base.GrpcError(c, err, "utils.Connect")
		return
	}
	defer conn.Close()
	defer cancel()

	// Make a gRPC request.
	service := pb.NewUsersServiceClient(conn)
	stream, err := service.CreateUser(ctx)
	if err != nil {
		base.GrpcError(c, err, "service.CreateUser")
		return
	}

	err = stream.Send(&request)
	if err != nil {
		base.GrpcError(c, err, "stream.Send")
		return
	}
	err = stream.CloseSend()
	if err != nil {
		base.GrpcError(c, err, "stream.CloseSend")
		return
	}

	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			base.GrpcError(c, err, "stream.Recv")
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {

	var request pb.User
	err := c.BindJSON(&request)
	if err != nil {
		base.GrpcError(c, err, "c.BindJSON")
		return
	}

	// Connect to gRPC server.
	conn, err, ctx, cancel := utils.Connect(base.ENV, base.URI_USER)
	if err != nil {
		base.GrpcError(c, err, "utils.Connect")
		return
	}
	defer conn.Close()
	defer cancel()

	// Make a gRPC request.
	service := pb.NewUsersServiceClient(conn)

	user, err := service.DeleteUser(ctx, &request)
	if err != nil {
		base.GrpcError(c, err, "service.DeleteUser")
		return
	}

	c.JSON(http.StatusOK, user)
}
