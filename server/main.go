package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"sgsg/auth"
	"sgsg/system"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "sgsg/proto"
)

type server struct {
	pb.UnimplementedServiceServer
	storage system.Storage
}

func main() {
	// Set up the logger
	system.InitLogger()

	// Connect to the database
	storage, err := system.NewStorage()
	if err != nil {
		slog.Error("Error opening database", "system.NewStorage", err)
		panic(err)
	}
	slog.Info("Database connected")

	// Run migrations
	err = storage.Migrations()
	if err != nil {
		slog.Error("Error running migrations", "storage.Migrations", err)
		panic(err)
	}
	slog.Info("Migrations completed")

	// Run the gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", system.GRPC_PORT))
	if err != nil {
		slog.Error("Error listening on gRPC port", "net.Listen", err)
		panic(err)
	}
	var s *grpc.Server
	if system.TLS == "true" {
		certificate, err := tls.LoadX509KeyPair(system.CERT_PATH, system.KEY_PATH)
		if err != nil {
			slog.Error("Error loading TLS certificate", "tls.LoadX509KeyPair", err)
			panic(err)
		}
		s = grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&certificate)))
	} else {
		s = grpc.NewServer()
	}
	pb.RegisterServiceServer(s, &server{
		storage: storage,
	})
	go func() {
		slog.Info("gRPC server listening on", "port", system.GRPC_PORT)
		err = s.Serve(lis)
		if err != nil {
			slog.Error("Error serving gRPC", "s.Serve", err)
			panic(err)
		}
	}()

	// Run the HTTP server
    var authDb = auth.NewAuthDB(&storage)
	var auth = auth.NewAuthService(authDb)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		defer system.Perf("ping", time.Now())
		c.Response().Header().Set("Access-Control-Allow-Origin", system.CLIENT_URL)
		id := 0
		err := storage.Conn.QueryRow("SELECT 1").Scan(&id)
		if err != nil {
			slog.Error("Error pinging database", "storage.Conn.QueryRow", err)
			return c.String(http.StatusInternalServerError, "Error pinging database")
		}
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/oauth-login/:provider", func(c echo.Context) error {
		return auth.OauthLogin(c)
	})
	e.GET("/oauth-callback/:provider", func(c echo.Context) error {
		return auth.OauthCallback(c)
	})
	go func() {
		slog.Info("HTTP server listening on", "port", system.HTTP_PORT)
		if system.TLS == "true" {
			err = e.StartTLS(":"+system.HTTP_PORT, system.CERT_PATH, system.KEY_PATH)
		} else {
			err = e.Start(":" + system.HTTP_PORT)
		}
		if err != nil {
			slog.Error("Error serving HTTP", "e.Start", err)
			panic(err)
		}
	}()

	// Run the system tasks
	go system.StartTask(context.Background(), auth.CleanTokens, time.Hour*24, "auth.CleanTokens")

	select {}
}
