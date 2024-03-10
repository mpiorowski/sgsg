package notes

import (
	pb "sgsg/proto"
	"sgsg/system"
	"strings"
	"sync"
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

func setup() NoteDB {
	storage := system.NewMemoryStorage()
	err := storage.Migrations()
	if err != nil {
		panic(err)
	}
	_, err = storage.Conn.Exec("delete from notes")
	if err != nil {
		panic(err)
	}
	return NewNoteDB(&storage)
}

func TestInsertNote(t *testing.T) {
	noteDB := setup()
	// Test case 1: Insert a valid note
	note, err := noteDB.InsertNote(&notes[0])
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
	note, err = noteDB.InsertNote(&notes[1])
	if err != nil {
		t.Errorf("insertNote error: %v", err)
	}
	if note.Title != notes[1].Title {
		t.Errorf("insertNote error: title not equal")
	}
}

func TestUpdateNote(t *testing.T) {
	noteDB := setup()
	// Test case 1: Update a valid note
	note, err := noteDB.InsertNote(&notes[0])
	if err != nil {
		t.Errorf("insertNote error: %v", err)
	}
	newNote := pb.Note{
		Id:      note.Id,
		UserId:  note.UserId,
		Title:   "New title",
		Content: note.Content,
	}
	note, err = noteDB.UpdateNote(&newNote)
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
	_, err = noteDB.UpdateNote(&newNote)
	if err == nil {
		t.Errorf("updateNoteTitle error: %v", err)
	}
}

func TestDeleteNoteById(t *testing.T) {
	noteDB := setup()
	// Test case 1: Delete a note
	note, err := noteDB.InsertNote(&notes[0])
	if err != nil {
		t.Errorf("insertNote error: %v", err)
	}
	err = noteDB.DeleteNoteByID(note.Id)
	if err != nil {
		t.Errorf("deleteNoteById error: %v", err)
	}
	note, err = noteDB.SelectNoteByID(note.Id, note.UserId)
	if note.Id != "" || err != nil {
		t.Errorf("selectNoteById error: %v", err)
	}

	// Test case 2: Delete a note that does not exist
	err = noteDB.DeleteNoteByID("not_exist")
	if err == nil {
		t.Errorf("deleteNoteById error: %v", err)
	}
}
func TestSelectNoteyId(t *testing.T) {
	noteDB := setup()
	// Test case 1: Select a note by id
	newNote, _ := noteDB.InsertNote(&notes[0])
	note, err := noteDB.SelectNoteByID(newNote.Id, newNote.UserId)
	if err != nil {
		t.Errorf("selectNoteId error: %v", err)
	}
	if note.Id != newNote.Id {
		t.Errorf("selectNoteId error: id not equal")
	}

	// Test case 2: Select a note by id that does not exist
	note, err = noteDB.SelectNoteByID("not_exist", "not_exist")
	if err != nil && note.Id != "" {
		t.Errorf("selectNoteId error: %v", err)
	}
}

func TestSelectNotes(t *testing.T) {
	noteDB := setup()
	// Test case 1: Select notes using stream
	_, _ = noteDB.InsertNote(&notes[0])
	_, _ = noteDB.InsertNote(&notes[1])
	_, _ = noteDB.InsertNote(&notes[2])

	notesCh := make(chan *pb.Note)
	errCh := make(chan error, 1)
	defer close(errCh)
	go noteDB.SelectNotesByUserID(notesCh, errCh, "test")

	count := 0
	for note := range notesCh {
		if note.UserId != notes[count].UserId {
			t.Errorf("selectNotes error: user_id not equal")
		}
		if note.Title != notes[count].Title {
			t.Errorf("selectNotes error: title not equal")
		}
		if note.Content != notes[count].Content {
			t.Errorf("selectNotes error: content not equal")
		}
		count++
	}

	if len(errCh) > 0 {
		t.Errorf("selectNotes error: %v", <-errCh)
	}

	if count != 3 {
		t.Errorf("scanNote error: count not equal")
	}
}

func TestNoteValidation(t *testing.T) {
	notes[0].Title = ""
	notes[0].Content = ""
	err := validateNote(&notes[0])
	if len(err) != 2 {
		t.Errorf("validation error: %v", err)
	}

	// gen 101 chars
	notes[0].Title = strings.Repeat("a", 101)
	notes[0].Content = strings.Repeat("a", 1001)
	err = validateNote(&notes[0])
	if len(err) != 2 || err[0].Tag != "max100" || err[1].Tag != "max1000" {
		t.Errorf("validation error: %v", err)
	}
}

func TestConcurrency(t *testing.T) {
	noteDB := setup()

	// Test case 1: Insert a note concurrently
	var newNotes []*pb.Note
	notesChanel := make(chan *pb.Note)
	wg := sync.WaitGroup{}
	gooroutines := 10
	for i := 0; i < gooroutines; i++ {
		println("i", i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			newNote, err := noteDB.InsertNote(&notes[0])
			if err != nil {
				t.Errorf("insertNote error: %v", err)
				return
			}
			notesChanel <- newNote
		}()
	}
	go func() {
		wg.Wait()
		close(notesChanel)
	}()
	for note := range notesChanel {
		newNotes = append(newNotes, note)
	}
	if len(newNotes) != gooroutines {
		t.Errorf("insertNote error: %v", len(newNotes))
		return
	}

	// Test case 2: Update a note concurrently
	done := make(chan bool)
	for i := 0; i < gooroutines; i++ {
		go func(i int) {
			note := newNotes[0]
			note.Title = "New title"
			_, err := noteDB.UpdateNote(note)
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
			_, err := noteDB.SelectNoteByID(note.Id, note.UserId)
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
			err := noteDB.DeleteNoteByID(note.Id)
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
