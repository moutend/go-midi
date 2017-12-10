package event

import "testing"

func TestSequencerSpecificEventDeltaTime(t *testing.T) {
	event := &SequencerSpecificEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("deltatime.DeltaTime() don't return nil")
	}
}

func TestSequencerSpecificEvent_String(t *testing.T) {
	event, err := NewSequencerSpecificEvent(nil, []byte{0x12, 0x34, 0x56})
	if err != nil {
		t.Fatal(err)
	}

	expected := "&SequencerSpecificEvent{data: 3 bytes}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestSequencerSpecificEvent_Serialize(t *testing.T) {
	event, err := NewSequencerSpecificEvent(nil, []byte{0x12, 0x34, 0x56})
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0xff, 0x7f, 0x03, 0x12, 0x34, 0x56}
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

func TestSequencerSpecificEvent_SetData(t *testing.T) {
	event := &SequencerSpecificEvent{}
	err := event.SetData(bigdata)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetData(bigdata[1:])
	if err != nil {
		t.Fatal(err)
	}
}

func TestSequencerSpecificEvent_Data(t *testing.T) {
	event := &SequencerSpecificEvent{data: []byte{0x12, 0x34, 0x56}}

	expected := []byte{0x12, 0x34, 0x56}
	actual := event.Data()

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

func TestNewSequencerSpecificEvent(t *testing.T) {
	event, err := NewSequencerSpecificEvent(nil, []byte{0x12, 0x34, 0x56})
	if err != nil {
		t.Fatal(err)
	}
	expected := []byte{0x12, 0x34, 0x56}
	actual := event.data

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
