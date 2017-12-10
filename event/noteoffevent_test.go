package event

import (
	"testing"

	"github.com/moutend/go-midi/constant"
)

func TestNoteOffEventDeltaTime(t *testing.T) {
	event := &NoteOffEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("deltatime.DeltaTime() don't return nil")
	}
}

func TestNoteOffEvent_String(t *testing.T) {
	event, err := NewNoteOffEvent(nil, 1, constant.C3, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&NoteOffEvent{channel: 1, note: C3, velocity: 50}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNoteOffEvent_Serialize(t *testing.T) {
	event, err := NewNoteOffEvent(nil, 0, constant.C3, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x80, 0x3c, 0x32}
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
func TestNoteOffEvent_SetChannel(t *testing.T) {
	event := &NoteOffEvent{}

	err := event.SetChannel(0x10)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetChannel(0x0f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoteOffEvent_Channel(t *testing.T) {
	event := &NoteOffEvent{channel: 1}

	expected := uint8(1)
	actual := event.Channel()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNoteOffEvent_SetNote(t *testing.T) {
	event := &NoteOffEvent{}

	err := event.SetNote(0x80)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetNote(0x7f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoteOffEvent_Note(t *testing.T) {
	event := &NoteOffEvent{note: constant.C3}

	expected := constant.C3
	actual := event.Note()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNoteOffEvent_SetVelocity(t *testing.T) {
	event := &NoteOffEvent{}

	err := event.SetVelocity(0x80)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetVelocity(0x7f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoteOffEvent_Velocity(t *testing.T) {
	event := &NoteOffEvent{velocity: 1}

	expected := uint8(1)
	actual := event.Velocity()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewNoteOffEvent(t *testing.T) {
	_, err := NewNoteOffEvent(nil, 255, 127, 127)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	_, err = NewNoteOffEvent(nil, 15, 255, 127)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	_, err = NewNoteOffEvent(nil, 15, 127, 255)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewNoteOffEvent(nil, 15, 127, 127)
	if err != nil {
		t.Fatal(err)
	}
	if event.channel != 15 {
		t.Fatalf("expected: 15 actual: %v", event.channel)
	}
	if event.note != 127 {
		t.Fatalf("expected: 127 actual: %v", event.note)
	}
	if event.velocity != 127 {
		t.Fatalf("expected: 127 actual: %v", event.velocity)
	}

}
