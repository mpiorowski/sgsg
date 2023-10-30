package profiles

import (
	"fmt"
	pb "sgsg/proto"
	"sgsg/utils"
)

func GetProfileByUserId(userId string) (*pb.Profile, error) {
	profile, err := selectProfileByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("selectProfileByUserId: %w", err)
	}
	return profile, nil
}

func CreateProfile(in *pb.Profile) (*pb.Profile, error) {
	rules := map[string]string{
		"UserId":   "required,uuid",
		"Username": "required,max=100",
		"About":    "required,max=1000",
		"ResumeId":   "max=1000",
		"CoverUrl":    "max=1000",
	}
	err := utils.ValidateStruct[pb.Profile](rules, pb.Profile{}, in)
	if err != nil {
		return nil, err
	}

	var profile *pb.Profile
	if in.Id == "" {
		profile, err = insertProfile(in)
	} else {
		profile, err = updateProfile(in)
	}
	if err != nil {
		return nil, fmt.Errorf("createProfile: %w", err)
	}
	return profile, nil
}

func DeleteProfile(id string) error {
	err := deleteProfileById(id)
	if err != nil {
		return fmt.Errorf("deleteProfileById: %w", err)
	}
	return nil
}
