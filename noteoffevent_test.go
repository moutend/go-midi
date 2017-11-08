package midi

import "testing"

func TestNoteOffEvent_DeltaTime(t *testing.T) {
	event := &NoteOffEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestNoteOffEvent_Serialize(t *testing.T) {
	event, err := NewNoteOffEvent(nil, 0, C3, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0x80, 0x3c, 0x32}
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
