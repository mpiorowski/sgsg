package notes

import (
	"io"
	"net/http"

	pb "go-svelte-grpc/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "go-svelte-grpc/gateway/base"
)

func GetNotes(c *gin.Context) {

	user, err := base.Authorization(c)
	if err != nil {
		return
	}

	// Make a gRPC request.
    ctx, err, cancel := base.CreateContext(base.ENV, base.URI_NOTES)
    if err != nil {
        base.GrpcError(c, err, "base.CreateContext")
        return
    }
    defer cancel()
	stream, err := base.NotesGrpcClient.GetNotes(ctx, &pb.UserId{
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

	// Make a gRPC request.
    ctx, err, cancel := base.CreateContext(base.ENV, base.URI_NOTES)
    if err != nil {
        base.GrpcError(c, err, "base.CreateContext")
        return
    }
    defer cancel()
    request.UserId = user.Id
    note, err := base.NotesGrpcClient.CreateNote(ctx, &request)
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

	// Make a gRPC request.
    ctx, err, cancel := base.CreateContext(base.ENV, base.URI_NOTES)
    if err != nil {
        base.GrpcError(c, err, "base.CreateContext")
        return
    }
    defer cancel()
	note, err := base.NotesGrpcClient.DeleteNote(ctx, &pb.NoteId{
        NoteId: noteId,
        UserId: user.Id,
    })
	if err != nil {
		base.GrpcError(c, err, "service.DeleteUser")
		return
	}

	c.JSON(http.StatusOK, note)
}
