package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

/**
 * ValidateStruct validates a struct using the rules provided.
 * @param rules map[string]string
 * @param s T
 * @param data *T
 * @return error
 */
func ValidateStruct[T interface{}](rules map[string]string, s T, data *T) error {

	validate := validator.New()
	validate.RegisterStructValidationMapRules(rules, s)
	err := validate.Struct(data)

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ValidationError, len(ve))
		for i, fe := range ve {
			out[i] = ValidationError{fe.Field(), fe.Tag()}
		}
		errorJSON, err := json.Marshal(out)
		if err != nil {
			return fmt.Errorf("json.Marshal: %w", err)
		}
		return status.Errorf(codes.InvalidArgument, "%s", errorJSON)
	}
	return nil
}
