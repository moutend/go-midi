package midi

import "testing"

func TestSetTempoEvent_DeltaTime(t *testing.T) {
	event := &SetTempoEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestSetTempoEvent_String(t *testing.T) {
	event, err := NewSetTempoEvent(nil, 123456)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&SetTempoEvent{tempo: 123456}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestSetTempoEvent_Serialize(t *testing.T) {
	event, err := NewSetTempoEvent(nil, 0x123456)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0xff, 0x51, 0x03, 0x12, 0x34, 0x56}
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

func TestSetTempoEvent_SetTempo(t *testing.T) {
	event := &SetTempoEvent{}

	err := event.SetTempo(0x80ffff)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetTempo(0x7fffff)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetTempoEvent_Tempo(t *testing.T) {
	event := &SetTempoEvent{tempo: 0x123456}

	expected := uint32(0x123456)
	actual := event.Tempo()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewSetTempoEvent(t *testing.T) {
	event, err := NewSetTempoEvent(nil, 0x123456)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint32(0x123456)
	actual := event.tempo

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}
