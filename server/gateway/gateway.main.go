package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

	"go-svelte-grpc/gateway/base"
	"go-svelte-grpc/gateway/notes"
	"go-svelte-grpc/gateway/users"

	pb "go-svelte-grpc/proto"
)

func init() {

    // Init Firebase Admin SDK
    var err error
    var app *firebase.App
	if base.ENV == "production" {
        app, err = firebase.NewApp(context.Background(), nil)
	} else {
		opt := option.WithCredentialsFile("../../serviceAccount.json")
		app, err = firebase.NewApp(context.Background(), nil, opt)
	}
    if err != nil {
        log.Fatalf("Error initializing firebase app: %v", err)
    }
    base.Client, err = app.Auth(base.Ctx)
    if err != nil {
        log.Fatalf("Error initializing firebase auth: %v", err)
    }

    // Init gRPC client
    usersConn, err := base.Connect(base.ENV, base.URI_USERS)
    if err != nil {
        log.Fatalf("Error connecting to users gRPC server: %v", err)
    }
    base.UsersGrpcClient = pb.NewUsersServiceClient(usersConn)

    notesConn, err := base.Connect(base.ENV, base.URI_NOTES)
    if err != nil {
        log.Fatalf("Error connecting to notes gRPC server: %v", err)
    }
    base.NotesGrpcClient = pb.NewNotesServiceClient(notesConn)

}


func main() {
	if base.ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{base.DOMAIN}
	config.AllowCredentials = true
    config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))

	router.GET("/auth", base.Auth)
	router.POST("/login", base.Login)

	router.GET("/files/:targetId", base.GetFiles)
	router.POST("/files", base.CreateFile)
	router.DELETE("/files/:targetId/:fileId", base.DeleteFile)

	router.GET("/users", users.GetUsers)
	router.POST("/users", users.CreateUser)
	router.DELETE("/users", users.DeleteUser)

    router.GET("/notes", notes.GetNotes)
    router.POST("/notes", notes.CreateNote)
    router.DELETE("/notes/:noteId", notes.DeleteNote)

	if err := router.Run(fmt.Sprintf("0.0.0.0:%v", base.PORT)); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Printf("Server listening at: %v", base.PORT)
}
