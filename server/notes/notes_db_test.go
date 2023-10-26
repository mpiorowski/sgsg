package notes

import (
	"sgsg/db"
	pb "sgsg/proto"
	"testing"
)

var notes = []pb.Note{
	{
		UserId:  "test",
		Title:   "Test note 1",
		Content: "Test note content 1",
	},
	{
		UserId:  "test",
		Title:   "Test note 2",
		Content: "Test note content 2",
	},
	{
		UserId:  "test",
		Title:   "Test note 3",
		Content: "Test note content 3",
	},
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
	note, err := insertNote(&notes[0])
	if err != nil {
		t.Errorf("insertNote error: %v", err)
	}
	if note.Title != notes[0].Title {
		t.Errorf("insertNote error: title not equal")
	}
	if note.Content != notes[0].Content {
		t.Errorf("insertNote error: content not equal")
	}
}

func TestUpdateNoteTitle(t *testing.T) {
	setup()
	note, err := insertNote(&notes[0])
	if err != nil {
		t.Errorf("insertNote error: %v", err)
	}
	newNote := pb.Note{
		Id:      note.Id,
		UserId:  note.UserId,
		Title:   "New title",
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

func TestDeleteNoteById(t *testing.T) {
    setup()
    note, err := insertNote(&notes[0])
    if err != nil {
        t.Errorf("insertNote error: %v", err)
    }
    err = deleteNoteById(note.Id)
    if err != nil {
        t.Errorf("deleteNoteById error: %v", err)
    }
    _, err = selectNoteById(note.Id, note.UserId)
    if err == nil {
        t.Errorf("selectNoteById error: %v", err)
    }
}

func TestSelectNotes(t *testing.T) {
	setup()
	_, _ = insertNote(&notes[0])
	_, _ = insertNote(&notes[1])
	_, _ = insertNote(&notes[2])
	notesStream, err := selectNotesStream(notes[0].UserId)
	if err != nil {
		t.Errorf("selectNotes error: %v", err)
	}

	count := 0
	for notesStream.Next() {
		note, err := scanNote(notesStream, nil)
		if err != nil {
			t.Errorf("scanNote error: %v", err)
		}
		if note.UserId != notes[count].UserId {
			t.Errorf("scanNote error: user_id not equal")
		}
		if note.Title != notes[count].Title {
			t.Errorf("scanNote error: title not equal")
		}
		if note.Content != notes[count].Content {
			t.Errorf("scanNote error: content not equal")
		}
        count++
	}
    if count != 3 {
        t.Errorf("scanNote error: count not equal")
    }
}

func TestSelectNoteyId(t *testing.T) {
    setup()
    newNote, _ := insertNote(&notes[2])
    note, err := selectNoteById(newNote.Id, newNote.UserId)
    if err != nil {
        t.Errorf("selectNoteId error: %v", err)
    }
    if note.Id != newNote.Id {
        t.Errorf("selectNoteId error: id not equal")
    }
}
