package event

import (
	"testing"

	"github.com/moutend/go-midi/constant"
)

func TestNoteAfterTouchEventDeltaTime(t *testing.T) {
	event := &NoteAfterTouchEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("deltatime.DeltaTime() don't return nil")
	}
}

func TestNoteAfterTouchEvent_String(t *testing.T) {
	event, err := NewNoteAfterTouchEvent(nil, 1, constant.C3, 50)
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
	event, err := NewNoteAfterTouchEvent(nil, 0, constant.C3, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0xa0, 0x3c, 0x32}
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

func TestNoteAfterTouchEvent_SetChannel(t *testing.T) {
	event := &NoteAfterTouchEvent{}

	err := event.SetChannel(0x10)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetChannel(0x0f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoteAfterTouchEvent_Channel(t *testing.T) {
	event := &NoteAfterTouchEvent{channel: 1}

	expected := uint8(1)
	actual := event.Channel()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNoteAfterTouchEvent_SetNote(t *testing.T) {
	event := &NoteAfterTouchEvent{}

	err := event.SetNote(0x80)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetNote(0x7f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoteAfterTouchEvent_Note(t *testing.T) {
	event := &NoteAfterTouchEvent{note: constant.C3}

	expected := constant.C3
	actual := event.Note()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNoteAfterTouchEvent_SetVelocity(t *testing.T) {
	event := &NoteAfterTouchEvent{}

	err := event.SetVelocity(0x80)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetVelocity(0x7f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoteAfterTouchEvent_Velocity(t *testing.T) {
	event := &NoteAfterTouchEvent{velocity: 1}

	expected := uint8(1)
	actual := event.Velocity()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewNoteAfterTouchEvent(t *testing.T) {
	_, err := NewNoteAfterTouchEvent(nil, 255, 127, 127)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	_, err = NewNoteAfterTouchEvent(nil, 15, 255, 127)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	_, err = NewNoteAfterTouchEvent(nil, 15, 127, 255)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewNoteAfterTouchEvent(nil, 15, 127, 127)
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
