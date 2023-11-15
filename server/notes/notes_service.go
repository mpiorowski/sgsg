package notes

import (
	"fmt"
	pb "sgsg/proto"
	"sgsg/utils"
)

func GetNotesStream(stream pb.Service_GetNotesServer, userId string) error {
	notes, err := selectNotesStream(userId)
	if err != nil {
		return fmt.Errorf("selectNotesStream: %w", err)
	}
	defer notes.Close()

	for notes.Next() {
		note, err := scanNote(notes, nil)
		if err != nil {
			return fmt.Errorf("scanNote: %w", err)
		}
		err = stream.Send(note)
		if err != nil {
			return fmt.Errorf("stream.Send: %w", err)
		}
	}

	err = notes.Err()
	if err != nil {
		return fmt.Errorf("notes.Err: %w", err)
	}
	return nil
}

func GetNoteById(id string, userId string) (*pb.Note, error) {
    note, err := selectNoteById(id, userId)
    if err != nil {
        return nil, fmt.Errorf("selectNoteById: %w", err)
    }
    return note, nil
}

func CreateNote(in *pb.Note) (*pb.Note, error) {
	rules := map[string]string{
		"UserId":  "required,uuid",
		"Title":   "required,max=100",
		"Content": "required,max=1000",
	}
	err := utils.ValidateStruct[pb.Note](rules, pb.Note{}, in)
	if err != nil {
		return nil, err
	}

	var note *pb.Note
	if in.Id == "" {
		note, err = insertNote(in)
	} else {
		note, err = updateNote(in)
	}
	if err != nil {
		return nil, fmt.Errorf("createNote: %w", err)
	}
	return note, nil
}

func DeleteNote(id string) error {
	err := deleteNoteById(id)
	if err != nil {
		return fmt.Errorf("deleteNoteById: %w", err)
	}
	return nil
}
