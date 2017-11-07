package midi

import "testing"

func TestNoteOnEvent_Serialize(t *testing.T) {
	event, err := NewNoteOnEvent(nil, 0, 18, 52)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0x90, 0x12, 0x34}
	actual := event.Serialize()

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v bytes actual: %v bytes", len(expected), len(actual))
	}
}
