package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/jackc/pgx/v5/stdlib"
	migrate "github.com/rubenv/sql-migrate"
	"google.golang.org/grpc"

	utils "github.com/mpiorowski/golang"
	pb "go-svelte-grpc/server/grpc"
)

var db *sql.DB

type server struct {
	pb.UnimplementedFilesServiceServer
}

var (
	PORT    = utils.MustGetenv("PORT")
	ENV     = utils.MustGetenv("ENV")
	BUCKET  = utils.MustGetenv("BUCKET")
	DB_USER = utils.MustGetenv("DB_USER")
	DB_PASS = utils.MustGetenv("DB_PASS")
	DB_HOST = utils.MustGetenv("DB_HOST")
	DB_NAME = utils.MustGetenv("DB_NAME")
)

func init() {
	// Db connection
	var err error
	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s",
		DB_USER, DB_PASS, DB_NAME, DB_HOST)
	if db, err = sql.Open("pgx", dbURI); err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected to database")

	// Migrations
	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Migrations failed: %v", err)
	}
	log.Printf("Applied %d migrations", n)
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", PORT))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFilesServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
