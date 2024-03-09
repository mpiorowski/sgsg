package system

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

type Claims struct {
	Id string `json:"id"`
}

func ExtractToken(ctx context.Context) (Claims, error) {
	claims := Claims{}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return claims, fmt.Errorf("Missing context metadata")
	}

	token := md.Get("x-authorization")
	if len(token) == 0 {
		return claims, fmt.Errorf("Missing authorization header")
	}

	// Validate the token
	tokenParts := strings.SplitN(token[0], " ", 2)
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return claims, fmt.Errorf("Invalid authorization header")
	}

	// Get the public public key from file
	publicKey, err := os.ReadFile("./public.key")
	if err != nil {
		return claims, fmt.Errorf("Invalid public key")
	}

	// Decode the token
	tokenString := tokenParts[1]
	jwtClaims := jwt.MapClaims{}
	t, err := jwt.ParseWithClaims(tokenString, jwtClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return jwt.ParseRSAPublicKeyFromPEM(publicKey)
	})
	if err != nil {
		return claims, fmt.Errorf("jwt.ParseWithClaims: %w", err)
	}
	if !t.Valid {
		return claims, fmt.Errorf("Invalid token")
	}
	if _, ok := jwtClaims["id"]; !ok {
		return claims, fmt.Errorf("Missing id in token")
	}

	claims.Id = jwtClaims["id"].(string)
	return claims, nil
}
