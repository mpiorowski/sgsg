package notes

import (
	"strings"
	"testing"
)

func TestCreateNote(t *testing.T) {
	setup()
	_, err := CreateNote(&notes[0])
	if err == nil {
		t.Errorf("updateNoteTitle error: %v", err)
	}
}

func TestValidation(t *testing.T) {
    setup()
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
