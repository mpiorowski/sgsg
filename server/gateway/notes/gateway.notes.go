package notes

import (
	"io"
	"net/http"

	pb "go-svelte-grpc/grpc"

	"github.com/gin-gonic/gin"
	utils "github.com/mpiorowski/golang"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "go-svelte-grpc/gateway/base"
)

func GetNotes(c *gin.Context) {

	user, err := base.Authorization(c)
	if err != nil {
		return
	}

	// Connect to gRPC server.
	conn, err, ctx, cancel := utils.Connect(base.ENV, base.URI_NOTES)
	if err != nil {
		base.GrpcError(c, err, "utils.Connect")
		return
	}
	defer conn.Close()
	defer cancel()

	// Make a gRPC request.
	service := pb.NewNotesServiceClient(conn)
	stream, err := service.GetNotes(ctx, &pb.UserId{
        UserId: user.Id,
    })

	if err != nil {
		base.GrpcError(c, err, "service.GetNotes")
		return
	}

	var notes []*pb.Note
	for {
		note, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			base.GrpcError(c, err, "stream.Recv")
			return
		}
		notes = append(notes, note)
	}

	c.JSON(http.StatusOK, notes)
}

func CreateNote(c *gin.Context) {

	var request pb.Note
	err := c.BindJSON(&request)
	if err != nil {
		base.GrpcError(c, err, "c.BindJSON")
		return
	}

    user, err := base.Authorization(c)
	if err != nil {
		return
	}

	// Connect to gRPC server.
	conn, err, ctx, cancel := utils.Connect(base.ENV, base.URI_NOTES)
	if err != nil {
		base.GrpcError(c, err, "utils.Connect")
		return
	}
	defer conn.Close()
	defer cancel()

	// Make a gRPC request.
	service := pb.NewNotesServiceClient(conn)

    request.UserId = user.Id
    note, err := service.CreateNote(ctx, &request)
    if err != nil {
        base.GrpcError(c, err, "service.CreateNote")
        return
    }

    c.JSON(http.StatusOK, note)
}

func DeleteNote(c *gin.Context) {

    noteId := c.Param("noteId")
    if noteId == "" {
        base.GrpcError(c, status.Error(codes.InvalidArgument, "noteId is empty"), "c.Param")
        return
    }

    user, err := base.Authorization(c)
	if err != nil {
		return
	}

	// Connect to gRPC server.
	conn, err, ctx, cancel := utils.Connect(base.ENV, base.URI_NOTES)
	if err != nil {
		base.GrpcError(c, err, "utils.Connect")
		return
	}
	defer conn.Close()
	defer cancel()

	// Make a gRPC request.
	service := pb.NewNotesServiceClient(conn)

	note, err := service.DeleteNote(ctx, &pb.NoteId{
        NoteId: noteId,
        UserId: user.Id,
    })
	if err != nil {
		base.GrpcError(c, err, "service.DeleteUser")
		return
	}

	c.JSON(http.StatusOK, note)
}
