package event

import "testing"

func TestProgramChangeEventDeltaTime(t *testing.T) {
	event := &ProgramChangeEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("deltatime.DeltaTime() don't return nil")
	}
}

func TestProgramChangeEvent_String(t *testing.T) {
	event, err := NewProgramChangeEvent(nil, 1, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&ProgramChangeEvent{channel: 1, program: SynthStrings1}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestProgramChangeEvent_Serialize(t *testing.T) {
	event, err := NewProgramChangeEvent(nil, 0, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0xc0, 0x32}
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

func TestProgramChangeEvent_SetChannel(t *testing.T) {
	event := &ProgramChangeEvent{}

	err := event.SetChannel(0x10)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetChannel(0x0f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProgramChangeEvent_Channel(t *testing.T) {
	event := &ProgramChangeEvent{channel: 1}

	expected := uint8(1)
	actual := event.Channel()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestProgramChangeEvent_SetProgram(t *testing.T) {
	event := &ProgramChangeEvent{}

	err := event.SetProgram(0x80)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetProgram(0x7f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProgramChangeEvent_Program(t *testing.T) {
	event := &ProgramChangeEvent{program: 1}

	expected := uint8(1)
	actual := uint8(event.Program())

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewProgramChangeEvent(t *testing.T) {
	_, err := NewProgramChangeEvent(nil, 255, 127)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	_, err = NewProgramChangeEvent(nil, 15, 255)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewProgramChangeEvent(nil, 15, 127)
	if err != nil {
		t.Fatal(err)
	}
	if event.channel != 15 {
		t.Fatalf("expected: 15 actual: %v", event.channel)
	}
	if event.program != 127 {
		t.Fatalf("expected: 127 actual: %v", event.program)
	}
}
