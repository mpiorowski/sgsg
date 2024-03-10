package notes

import (
	pb "sgsg/proto"
	"sgsg/system"
)

func validateNote(in *pb.Note) []system.ValidationError {
	var errors []system.ValidationError
	if in.UserId == "" {
		errors = append(errors, system.ValidationError{Field: "UserId", Tag: "required"})
	}
	if in.Title == "" {
		errors = append(errors, system.ValidationError{Field: "Title", Tag: "required"})
	}
	if len(in.Title) > 100 {
		errors = append(errors, system.ValidationError{Field: "Title", Tag: "max100"})
	}
	if in.Content == "" {
		errors = append(errors, system.ValidationError{Field: "Content", Tag: "required"})
	}
	if len(in.Content) > 1000 {
		errors = append(errors, system.ValidationError{Field: "Content", Tag: "max1000"})
	}
	return errors
}
