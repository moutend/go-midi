package midi

import "testing"

func TestNoteAfterTouchEvent_DeltaTime(t *testing.T) {
	event := &NoteAfterTouchEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestNoteAfterTouchEvent_String(t *testing.T) {
	event, err := NewNoteAfterTouchEvent(nil, 1, C3, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&NoteAfterTouchEvent{channel: 1, note: C3, velocity: 50}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNoteAfterTouchEvent_Serialize(t *testing.T) {
	event, err := NewNoteAfterTouchEvent(nil, 0, C3, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0xa0, 0x3c, 0x32}
	actual := event.Serialize()

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v bytes actual: %v bytes", len(expected), len(actual))
	}
	for i, e := range expected {
		a := actual[i]
		if e != a {
			t.Fatalf("expected[%v] = 0x%x actual[%v] = 0x%x", i, e, i, a)
		}
	}
}
