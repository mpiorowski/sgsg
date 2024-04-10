package system

import (
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

func CreateErrors (errors []ValidationError) error {
	errorJSON, err := json.Marshal(errors)
	if err != nil {
		return status.Errorf(codes.Internal, "json.Marshal: %v", err)
	}
	return status.Errorf(codes.InvalidArgument, "%s", errorJSON)
}
