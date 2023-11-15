package notes

import (
	"sgsg/db"
	pb "sgsg/proto"
	"strings"
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

func clearNotes() {
	_, err := db.Db.Exec("delete from notes")
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	err := db.ConnectTest()
	if err != nil {
		panic(err)
	}
	err = db.Migrations()
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestInsertNote(t *testing.T) {
	// Test case 1: Insert a valid note
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

	// Test case 2: Insert a second valid note
	note, err = insertNote(&notes[1])
	if err != nil {
		t.Errorf("insertNote error: %v", err)
	}
	if note.Title != notes[1].Title {
		t.Errorf("insertNote error: title not equal")
	}
}

func TestUpdateNote(t *testing.T) {
	// Test case 1: Update a valid note
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

	// Test case 2: Update a note that does not exist
	newNote = pb.Note{
		Id: "not_exist",
	}
	_, err = updateNote(&newNote)
	if err == nil {
		t.Errorf("updateNoteTitle error: %v", err)
	}
}

func TestDeleteNoteById(t *testing.T) {
	// Test case 1: Delete a note
	note, err := insertNote(&notes[0])
	if err != nil {
		t.Errorf("insertNote error: %v", err)
	}
	err = deleteNoteById(note.Id)
	if err != nil {
		t.Errorf("deleteNoteById error: %v", err)
	}
	note, err = selectNoteById(note.Id, note.UserId)
	if note.Id != "" || err != nil {
		t.Errorf("selectNoteById error: %v", err)
	}

	// Test case 2: Delete a note that does not exist
	err = deleteNoteById("not_exist")
	if err == nil {
		t.Errorf("deleteNoteById error: %v", err)
	}
}
func TestSelectNoteyId(t *testing.T) {
	clearNotes()
	// Test case 1: Select a note by id
	newNote, _ := insertNote(&notes[2])
	note, err := selectNoteById(newNote.Id, newNote.UserId)
	if err != nil {
		t.Errorf("selectNoteId error: %v", err)
	}
	if note.Id != newNote.Id {
		t.Errorf("selectNoteId error: id not equal")
	}

    // Test case 2: Select a note by id that does not exist
    note, err = selectNoteById("not_exist", "not_exist")
    if err != nil && note.Id != "" {
        t.Errorf("selectNoteId error: %v", err)
    }
}


func TestSelectNotes(t *testing.T) {
	clearNotes()
	// Test case 1: Select notes using stream
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

func TestNoteValidation(t *testing.T) {
	notes[0].Title = ""
	notes[0].Content = ""
	_, err := CreateNote(&notes[0])
	containsTitle := strings.Contains(err.Error(), "Title") && strings.Contains(err.Error(), "required")
	containsContent := strings.Contains(err.Error(), "Content") && strings.Contains(err.Error(), "required")
	if !containsTitle || !containsContent {
		t.Errorf("validation error: %v", err)
	}

	// gen 101 chars
	notes[0].Title = strings.Repeat("a", 101)
	notes[0].Content = strings.Repeat("a", 1001)
	_, err = CreateNote(&notes[0])
	containsTitle = strings.Contains(err.Error(), "Title") && strings.Contains(err.Error(), "max")
	containsContent = strings.Contains(err.Error(), "Content") && strings.Contains(err.Error(), "max")
	if !containsTitle || !containsContent {
		t.Errorf("validation error: %v", err)
	}
}

func TestConcurrency(t *testing.T) {
	newNotes := make([]*pb.Note, 0)
	// Test case 1: Insert a note concurrently
	notesChanel := make(chan *pb.Note)
	gooroutines := 10
	for i := 0; i < gooroutines; i++ {
		go func() {
			newNote, err := insertNote(&notes[0])
			if err != nil {
				t.Errorf("insertNote error: %v", err)
			}
			notesChanel <- newNote
		}()
	}

	for i := 0; i < gooroutines; i++ {
		n := <-notesChanel
		newNotes = append(newNotes, n)
	}

	// Test case 2: Update a note concurrently
	done := make(chan bool)
	for i := 0; i < gooroutines; i++ {
		go func(i int) {
			note := newNotes[i]
			note.Title = "New title"
			_, err := updateNote(note)
			if err != nil {
				t.Errorf("updateNoteTitle error: %v", err)
			}
			done <- true
		}(i)
	}

	for i := 0; i < gooroutines; i++ {
		<-done
	}

	// Test case 4: Select note by id concurrently
	done = make(chan bool)
	for i := 0; i < gooroutines; i++ {
		go func(i int) {
			note := newNotes[i]
			_, err := selectNoteById(note.Id, note.UserId)
			if err != nil {
				t.Errorf("selectNotes error: %v", err)
			}
			done <- true
		}(i)
	}

	for i := 0; i < gooroutines; i++ {
		<-done
	}

	// Test case 3: Delete a note concurrently
	done = make(chan bool)
	for i := 0; i < gooroutines; i++ {
		go func(i int) {
			note := newNotes[i]
			err := deleteNoteById(note.Id)
			if err != nil {
				t.Errorf("deleteNoteById error: %v", err)
			}
			done <- true
		}(i)
	}

	for i := 0; i < gooroutines; i++ {
		<-done
	}
}
