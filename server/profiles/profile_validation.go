package profiles

import (
	"fmt"
	pb "sgsg/proto"
	"sgsg/utils"
)

func validateProfile(in *pb.Profile) error {
	rules := map[string]string{
		"UserId":   "required,uuid",
		"Username": "required,max=100",
		"About":    "required,max=1000",
		"ResumeId": "max=1000",
		"CoverId":  "max=1000",
		"CoverUrl": "max=1000",
	}
	err := utils.ValidateStruct[pb.Profile](rules, pb.Profile{}, in)
	if err != nil {
		return fmt.Errorf("validateProfile error: %w", err)
	}
	return nil
}
