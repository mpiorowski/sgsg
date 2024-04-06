package system

import (
	"os"
	"strings"
)

func isRunningTest() bool {
	for _, arg := range os.Args {
		if strings.HasSuffix(arg, ".test") {
			return true
		}
	}
	return false
}

func mustHaveEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		if isRunningTest() {
			return "test"
		}
		panic("Missing environment variable: " + key)
	}
	return value
}

var (
	LOG_LEVEL            = mustHaveEnv("LOG_LEVEL")
	HTTP_PORT            = mustHaveEnv("HTTP_PORT")
	GRPC_PORT            = mustHaveEnv("GRPC_PORT")
	JWT_SECRET           = mustHaveEnv("JWT_SECRET")
	TURSO_URL            = mustHaveEnv("TURSO_URL")
	CLIENT_URL           = mustHaveEnv("CLIENT_URL")
	SERVER_HTTP          = mustHaveEnv("SERVER_HTTP")
	STRIPE_API_KEY       = mustHaveEnv("STRIPE_API_KEY")
	STRIPE_PRICE_ID      = mustHaveEnv("STRIPE_PRICE_ID")
	GOOGLE_CLIENT_ID     = mustHaveEnv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET = mustHaveEnv("GOOGLE_CLIENT_SECRET")
	GITHUB_CLIENT_ID     = mustHaveEnv("GITHUB_CLIENT_ID")
	GITHUB_CLIENT_SECRET = mustHaveEnv("GITHUB_CLIENT_SECRET")
)
