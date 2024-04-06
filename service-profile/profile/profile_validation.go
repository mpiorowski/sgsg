package profile

import (
	pb "service-profile/proto"
	"service-profile/system"
)

func validateProfile(profile *pb.Profile) []system.ValidationError {
	var errors []system.ValidationError
	if profile.Username == "" {
		errors = append(errors, system.ValidationError{Field: "Username", Tag: "required"})
	}
	// max length 100
	if len(profile.Username) > 100 {
		errors = append(errors, system.ValidationError{Field: "Username", Tag: "max100"})
	}
	if profile.About == "" {
		errors = append(errors, system.ValidationError{Field: "About", Tag: "required"})
	}
	// max length 1000
	if len(profile.About) > 1000 {
		errors = append(errors, system.ValidationError{Field: "About", Tag: "max1000"})
	}
    return errors
}
