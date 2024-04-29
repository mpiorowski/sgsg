package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"service-auth/service"
	"service-auth/store"
	"service-auth/system"
	"time"

	"google.golang.org/grpc"

	pb "service-auth/proto"
)

type server struct {
	pb.UnimplementedAuthServiceServer
	storage system.Storage
}

func main() {
	// Set up the logger
	system.InitLogger()

	// Connect to the database
	storage, err, clean := system.NewStorage()
	defer clean()
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
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{
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
	mux := http.NewServeMux()
	mux.HandleFunc("/oauth-login/{provider}", func(w http.ResponseWriter, r *http.Request) {
		service.OauthLogin(storage, w, r)
	})
	mux.HandleFunc("/oauth-callback/{provider}", func(w http.ResponseWriter, r *http.Request) {
		service.OauthCallback(storage, w, r)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer system.Perf("ping", time.Now())
		w.Header().Set("Access-Control-Allow-Origin", system.CLIENT_URL)
		id := 0
		err := storage.Conn.QueryRow("SELECT 1").Scan(&id)
		if err != nil {
			slog.Error("Error pinging database", "storage.Conn.QueryRow", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("Hello, World!"))
		if err != nil {
			slog.Error("Error writing response", "w.Write", err)
		}
	})

	go func() {
		slog.Info("HTTP server listening on", "port", system.HTTP_PORT)
		err = http.ListenAndServe(":"+system.HTTP_PORT, mux)
		if err != nil {
			slog.Error("Error serving HTTP", "http.ListenAndServe", err)
			panic(err)
		}
	}()

	// Run the system tasks
	var authDB = store.NewAuthDB(&storage)
	go system.StartTask(context.Background(), authDB.CleanTokens, time.Hour*24, "auth.CleanTokens")

	select {}
}
