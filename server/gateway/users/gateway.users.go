package users

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "go-svelte-grpc/proto"

	base "go-svelte-grpc/gateway/base"
)

func GetUsers(c *gin.Context) {

	user, err := base.Authorization(c)
	if err != nil {
		return
	}
	if user.GetRole() != pb.UserRole_ROLE_ADMIN.String() {
		log.Printf("Unauthorized user: %v", user.GetEmail())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Make a gRPC request.
    ctx, err, cancel := base.CreateContext(base.ENV, base.URI_USERS)
    if err != nil {
        base.GrpcError(c, err, "base.CreateContext")
        return
    }
    defer cancel()
	stream, err := base.UsersGrpcClient.GetUsers(ctx, &pb.Empty{})

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

	// Make a gRPC request.
    ctx, err, cancel := base.CreateContext(base.ENV, base.URI_USERS)
    if err != nil {
        base.GrpcError(c, err, "base.CreateContext")
        return
    }
    defer cancel()
	stream, err := base.UsersGrpcClient.CreateUser(ctx)
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

	_, err = base.Authorization(c)
	if err != nil {
		return
	}
	// Make a gRPC request.
    ctx, err, cancel := base.CreateContext(base.ENV, base.URI_USERS)
    if err != nil {
        base.GrpcError(c, err, "base.CreateContext")
        return
    }
    defer cancel()
	user, err := base.UsersGrpcClient.DeleteUser(ctx, &request)
	if err != nil {
		base.GrpcError(c, err, "service.DeleteUser")
		return
	}

	c.JSON(http.StatusOK, user)
}
