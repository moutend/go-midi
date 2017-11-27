package midi

import "testing"

func TestMIDIPortPrefixEvent_DeltaTime(t *testing.T) {
	event := &MIDIPortPrefixEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestMIDIPortPrefixEvent_String(t *testing.T) {
	event, err := NewMIDIPortPrefixEvent(nil, 1)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&MIDIPortPrefixEvent{port: 1}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestMIDIPortPrefixEvent_Serialize(t *testing.T) {
	event, err := NewMIDIPortPrefixEvent(nil, 12)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0xff, 0x20, 0x01, 0x0c}
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
func TestMIDIPortPrefixEvent_SetPort(t *testing.T) {
	event := &MIDIPortPrefixEvent{}

	err := event.SetPort(0x10)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetPort(0x0f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMIDIPortPrefixEvent_Port(t *testing.T) {
	event := &MIDIPortPrefixEvent{port: 1}

	expected := uint8(1)
	actual := event.Port()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewMIDIPortPrefixEvent(t *testing.T) {
	_, err := NewMIDIPortPrefixEvent(nil, 255)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewMIDIPortPrefixEvent(nil, 15)
	if err != nil {
		t.Fatal(err)
	}
	if event.port != 15 {
		t.Fatalf("expected: 15 actual: %v", event.port)
	}
}
