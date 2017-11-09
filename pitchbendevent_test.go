package midi

import "testing"

func TestPitchBendEvent_DeltaTime(t *testing.T) {
	event := &PitchBendEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestPitchBendEvent_String(t *testing.T) {
	event, err := NewPitchBendEvent(nil, 1, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&PitchBendEvent{channel: 1, pitch: 50}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestPitchBendEvent_Serialize(t *testing.T) {
	event, err := NewPitchBendEvent(nil, 1, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0xe1, 0x32, 0x00}
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
