package utils

import "os"

func mustHaveEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Environment variable not set: " + key)
	}
	return value
}

var (
	ENV                  = mustHaveEnv("ENV")
	HTTP_PORT            = mustHaveEnv("HTTP_PORT")
	GRPC_PORT            = mustHaveEnv("GRPC_PORT")
	CLIENT_URL           = mustHaveEnv("CLIENT_URL")
	SERVER_HTTP          = mustHaveEnv("SERVER_HTTP")
	COOKIE_DOMAIN        = mustHaveEnv("COOKIE_DOMAIN")
	CERT_PATH            = mustHaveEnv("CERT_PATH")
	KEY_PATH             = mustHaveEnv("KEY_PATH")
	GOOGLE_CLIENT_ID     = mustHaveEnv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET = mustHaveEnv("GOOGLE_CLIENT_SECRET")
	GITHUB_CLIENT_ID     = mustHaveEnv("GITHUB_CLIENT_ID")
	GITHUB_CLIENT_SECRET = mustHaveEnv("GITHUB_CLIENT_SECRET")
)
