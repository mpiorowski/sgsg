package base

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/mpiorowski/golang"
	pb "go-svelte-grpc/grpc"
)

func GetFiles(c *gin.Context) {

	targetId := c.Param("targetId")
	if targetId == "" {
		GatewayError(c, "targetId is required", "c.Param")
		return
	}

	_, err := Authorization(c)
	if err != nil {
		return
	}

	// Connect to the gRPC server
	conn, err, ctx, cancel := utils.Connect(ENV, URI_FILES)
	if err != nil {
		GrpcError(c, err, "utils.Connect")
		return
	}
	defer conn.Close()
	defer cancel()

	// Make a gRPC request.
	service := pb.NewFilesServiceClient(conn)
	stream, err := service.GetFiles(ctx, &pb.TargetId{TargetId: targetId})
	if err != nil {
		GrpcError(c, err, "service.GetFiles")
		return
	}
	var files []*pb.File
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			GrpcError(c, err, "stream.Recv")
			return
		}
		files = append(files, r)
	}

	c.JSON(http.StatusOK, files)
}

func CreateFile(c *gin.Context) {

	// TODO - multi files
	_, err := Authorization(c)
	if err != nil {
		return
	}

	// Parse the multipart form in the request
	var body pb.File
	form, err := c.MultipartForm()
	if err != nil {
        GatewayError(c, "Invalid form", "c.MultipartForm")
		return
	}
	request := form.Value["request"][0]
	err = json.Unmarshal([]byte(request), &body)
	if err != nil {
        GatewayError(c, "Invalid request", "json.Unmarshal")
		return
	}

	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
        GatewayError(c, "Invalid file", "c.FormFile")
		return
	}
	open, err := file.Open()
	if err != nil {
        GatewayError(c, "Invalid file", "file.Open")
		return
	}
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, open)
	if err != nil {
        GatewayError(c, "Invalid file", "io.Copy")
		return
	}

	// Connect to the gRPC server
	conn, err, ctx, cancel := utils.Connect(ENV, URI_FILES)
	if err != nil {
        GrpcError(c, err, "utils.Connect")
		return
	}
	defer conn.Close()
	defer cancel()

	// Make a gRPC request.
	service := pb.NewFilesServiceClient(conn)

	r, err := service.CreateFile(ctx, &pb.File{
		Id:       body.Id,
		TargetId: body.TargetId,
		Name:     body.Name,
		Type:     body.Type,
		Data:     buf.Bytes(),
	})
	if err != nil {
        GrpcError(c, err, "service.CreateFile")
		return
	}
	c.JSON(http.StatusOK, r)
}

func DeleteFile(c *gin.Context) {
	fileId := c.Param("fileId")
	targetId := c.Param("targetId")
	if fileId == "" || targetId == "" {
        GatewayError(c, "fileId or targetId is required", "c.Param")
		return
	}

	_, err := Authorization(c)
	if err != nil {
		return
	}

    // Connect to the gRPC server
	conn, err, ctx, cancel := utils.Connect(ENV, URI_FILES)
	if err != nil {
        GrpcError(c, err, "utils.Connect")
		return
	}
	defer conn.Close()
	defer cancel()

	// Make a gRPC request.
	service := pb.NewFilesServiceClient(conn)
	response, err := service.DeleteFile(ctx, &pb.FileId{FileId: fileId, TargetId: targetId})

	if err != nil {
        GrpcError(c, err, "service.DeleteFile")
		return
	}
	c.JSON(http.StatusOK, response)
}
