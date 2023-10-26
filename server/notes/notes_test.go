package notes

import (
	"sgsg/db"
	pb "sgsg/proto"
	"testing"
)

var initNote = pb.Note{
	Title:   "Test note",
	Content: "Test note content",
}

func Clear() {
    _, err := db.Db.Exec("DELETE FROM notes")
    if err != nil {
        panic(err)
    }
}

func TestInsertNote(t *testing.T) {
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
    Clear()
}
