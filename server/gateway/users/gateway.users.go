package users

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/mpiorowski/golang"
	pb "github.com/mpiorowski/go-svelte-grpc/server/grpc"

    base "github.com/mpiorowski/go-svelte-grpc/server/gateway/base"
)

func GetUsers(c *gin.Context) {

	user, err := base.Authorization(c)
	if err != nil {
		return
	}
    if user.GetRole() != pb.UserRole_ROLE_ADMIN.String() {
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
