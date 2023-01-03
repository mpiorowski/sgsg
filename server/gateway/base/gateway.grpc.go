package base

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	grpcMetadata "google.golang.org/grpc/metadata"
	pb "go-svelte-grpc/grpc"
)

type Cache struct {
    token string
    expiry time.Time
}

var cacheByHost = map[string]Cache{}

var (
    UsersGrpcClient pb.UsersServiceClient
    NotesGrpcClient pb.NotesServiceClient
)

func Connect(env string, host string) (*grpc.ClientConn, error) {
	// Create a TLS credentials object with the certificate authority.
	host = fmt.Sprintf("%s:443", host)

	// For local development use insecure credentials
	if env != "production" {
		conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("grpc.Dial: %v", err)
			return nil, err
		}
		return conn, nil
	}

	var opts []grpc.DialOption
	if host != "" {
		opts = append(opts, grpc.WithAuthority(host))
	}
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Printf("x509.SystemCertPool: %v", err)
		return nil, err
	}
	cred := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	opts = append(opts, grpc.WithTransportCredentials(cred))

	// Dial the server.
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Printf("grpc.Dial: %v", err)
		return nil, err
	}
	return conn, nil
}

func CreateContext(env string, host string) (context.Context, error, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if env != "production" {
        return ctx, nil, cancel
	}

    token := cacheByHost[host].token
    expiry := cacheByHost[host].expiry

	if token == "" || time.Now().After(expiry) {
        log.Printf("Token not found in cache or expired")
		// Create an identity token.
		// With a global TokenSource tokens would be reused and auto-refreshed at need.
		// A given TokenSource is specific to the audience.
		tokenSource, err := idtoken.NewTokenSource(ctx, "https://"+host)
		if err != nil {
			log.Printf("idtoken.NewTokenSource: %v", err)
            return nil, err, cancel
		}
		tokenStruct, err := tokenSource.Token()
		if err != nil {
			log.Printf("tokenSource.Token: %v", err)
            return nil, err, cancel
		}
		token = tokenStruct.AccessToken
		// Cache token
        cacheByHost[host] = Cache{
            token: token,
            expiry: time.Now().Add(time.Hour),
        }
	} else {
        log.Printf("Token found in cache")
    }

	// Add token to gRPC Request.
	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
    return ctx, nil, cancel
}
