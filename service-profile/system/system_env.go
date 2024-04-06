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
	GRPC_PORT            = mustHaveEnv("GRPC_PORT")
	JWT_SECRET           = mustHaveEnv("JWT_SECRET")
	TURSO_URL            = mustHaveEnv("TURSO_URL")
)
