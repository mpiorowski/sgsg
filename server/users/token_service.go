package users

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

type Claims struct {
	TokenId string
}

func extractToken(ctx context.Context) (Claims, error) {
    claim := Claims{}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return claim, fmt.Errorf("Missing context metadata")
	}

	token := md.Get("x-authorization")
	if len(token) == 0 {
		return claim, fmt.Errorf("Missing authorization header")
	}

	// Validate the token
	tokenParts := strings.SplitN(token[0], " ", 2)
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return claim, fmt.Errorf("Invalid authorization header")
	}

	// Get the public public key from file
	publicKey, err := os.ReadFile("./public.key")
	if err != nil {
		return claim, fmt.Errorf("Invalid public key")
	}

	// Decode the token
	tokenString := tokenParts[1]
	claims, err := decodeToken(tokenString, publicKey)
	if err != nil {
		return claim, fmt.Errorf("Invalid token: %w", err)
	}
	return *claims, nil
}

func decodeToken(tokenString string, publicKey []byte) (*Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return jwt.ParseRSAPublicKeyFromPEM(publicKey)
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	// check if tokenId is in the claims
	if _, ok := claims["tokenId"]; !ok {
		return nil, fmt.Errorf("Invalid token")
	}

	return &Claims{
		TokenId: claims["tokenId"].(string),
	}, nil
}
