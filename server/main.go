package main

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"sgsg/db"
	"sgsg/users"
	"sgsg/utils"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "sgsg/proto"
)

type server struct {
	pb.UnimplementedServiceServer
}

func main() {
    utils.SetupEnv()

	err := db.Connect()
	if err != nil {
		slog.Error("Error opening database", "db.Connect", err)
		panic(err)
	}
	slog.Info("Database connected")

	err = db.Migrations()
	if err != nil {
		slog.Error("Error running migrations", "db.Migrations", err)
		panic(err)
	}
	slog.Info("Migrations completed")

	// Run the gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", utils.GRPC_PORT))
	if err != nil {
		slog.Error("Error listening on gRPC port", "net.Listen", err)
		panic(err)
	}

	var s *grpc.Server
	if utils.ENV == "production" {
		certificate, err := tls.LoadX509KeyPair(utils.CERT_PATH, utils.KEY_PATH)
		if err != nil {
			slog.Error("Error loading TLS certificate", "tls.LoadX509KeyPair", err)
			panic(err)
		}
		s = grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&certificate)))
	} else {
		s = grpc.NewServer()
	}
	pb.RegisterServiceServer(s, &server{})
	go func() {
		slog.Info("gRPC server listening on", "port", utils.GRPC_PORT)
		err = s.Serve(lis)
		if err != nil {
			slog.Error("Error serving gRPC", "s.Serve", err)
			panic(err)
		}
	}()

	// Run the HTTP server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		id := 0
		err := db.Db.QueryRow("SELECT 1").Scan(&id)
		if err != nil {
			slog.Error("Error pinging database", "Db.QueryRow", err)
			return c.String(http.StatusInternalServerError, "Error pinging database")
		}
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/oauth-login/:provider", func(c echo.Context) error {
		return users.OauthLogin(c, &users.OAuthConfigImpl{})
	})
	e.GET("/oauth-callback/:provider", func(c echo.Context) error {
		return users.OauthCallback(c, &users.OAuthConfigImpl{})
	})

	go func() {
		slog.Info("HTTP server listening on", "port", utils.HTTP_PORT)
		if utils.ENV == "production" {
			err = e.StartTLS(":"+utils.HTTP_PORT, utils.CERT_PATH, utils.KEY_PATH)
		} else {
			err = e.Start(":" + utils.HTTP_PORT)
		}
		if err != nil {
			slog.Error("Error serving HTTP", "e.Start", err)
			panic(err)
		}
	}()

	select {}
}
