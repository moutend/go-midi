package midi

import "testing"

func TestMIDIChannelPrefixEvent_DeltaTime(t *testing.T) {
	event := &MIDIChannelPrefixEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestMIDIChannelPrefixEvent_String(t *testing.T) {
	event, err := NewMIDIChannelPrefixEvent(nil, 1)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&MIDIChannelPrefixEvent{channel: 1}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestMIDIChannelPrefixEvent_Serialize(t *testing.T) {
	event, err := NewMIDIChannelPrefixEvent(nil, 12)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0xff, 0x21, 0x01, 0x0c}
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
func TestMIDIChannelPrefixEvent_SetChannel(t *testing.T) {
	event := &MIDIChannelPrefixEvent{}

	err := event.SetChannel(0x10)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetChannel(0x0f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMIDIChannelPrefixEvent_Channel(t *testing.T) {
	event := &MIDIChannelPrefixEvent{channel: 1}

	expected := uint8(1)
	actual := event.Channel()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewMIDIChannelPrefixEvent(t *testing.T) {
	_, err := NewMIDIChannelPrefixEvent(nil, 255)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewMIDIChannelPrefixEvent(nil, 15)
	if err != nil {
		t.Fatal(err)
	}
	if event.channel != 15 {
		t.Fatalf("expected: 15 actual: %v", event.channel)
	}
}
