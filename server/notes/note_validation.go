package notes

import (
	"fmt"
	pb "sgsg/proto"
	"sgsg/utils"
)

func validateNote(in *pb.Note) error {
	rules := map[string]string{
		"UserId":  "required,uuid",
		"Title":   "required,max=100",
		"Content": "required,max=1000",
	}
	err := utils.ValidateStruct[pb.Note](rules, pb.Note{}, in)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	return nil
}
