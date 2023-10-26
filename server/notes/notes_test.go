package notes

import (
	"sgsg/db"
	pb "sgsg/proto"
	"testing"
)

var initNote = pb.Note{
    UserId: "test",
	Title:   "Test note",
	Content: "Test note content",
}

func setup() {
    err := db.ConnectTest()
    if err != nil {
        panic(err)
    }
    err = db.Migrations()
    if err != nil {
        panic(err)
    }
}

func TestInsertNote(t *testing.T) {
	setup()
	note, err := insertNote(&initNote)
	if err != nil {
		t.Errorf("insertNote error: %v", err)
	}
	if note.Title != initNote.Title {
		t.Errorf("insertNote error: title not equal")
	}
	if note.Content != initNote.Content {
		t.Errorf("insertNote error: content not equal")
	}
}

func TestUpdateNoteTitle(t *testing.T) {
    setup()
    note, err := insertNote(&initNote)
    if err != nil {
        t.Errorf("insertNote error: %v", err)
    }
    newNote := pb.Note{
        Id: note.Id,
        UserId: note.UserId,
        Title: "New title",
        Content: note.Content,
    }
    note, err = updateNote(&newNote)
    if err != nil {
        t.Errorf("updateNoteTitle error: %v", err)
    }
    if note.Title != "New title" {
        t.Errorf("updateNoteTitle error: title not equal")
    }
}

func TestErrorUpdateNote(t *testing.T) {
    setup()
    newNote := pb.Note{
        Id: "not_exist",
    }
    _, err := updateNote(&newNote)
    if err == nil {
        t.Errorf("updateNoteTitle error: %v", err)
    }
}

