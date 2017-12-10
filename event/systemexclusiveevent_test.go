package event

import "testing"

func TestSystemExclusiveEventDeltaTime(t *testing.T) {
	event := &SystemExclusiveEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("deltatime.DeltaTime() don't return nil")
	}
}

func TestSystemExclusiveEvent_String(t *testing.T) {
	event, err := NewSystemExclusiveEvent(nil, []byte{0x01, 0x02, 0x03, 0x04})
	if err != nil {
		t.Fatal(err)
	}

	expected := "&SystemExclusiveEvent{data: 4 bytes}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestSystemExclusiveEvent_Serialize(t *testing.T) {
	event, err := NewSystemExclusiveEvent(nil, []byte{0x11, 0x12, 0x13, 0x14})
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0xf0, 0x04, 0x11, 0x12, 0x13, 0x14}
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

func TestSystemExclusiveEvent_SetData(t *testing.T) {
	event := &SystemExclusiveEvent{}

	err := event.SetData(bigdata)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetData(bigdata[1:])
	if err != nil {
		t.Fatal(err)
	}
}

func TestSystemExclusiveEvent_Data(t *testing.T) {
	event := &SystemExclusiveEvent{}
	if event.Data() == nil {
		t.Fatalf("Data() must return empty slice")
	}

	event.data = []byte{0x11, 0x12, 0x13, 0x14}
	expected := []byte{0x11, 0x12, 0x13, 0x14}
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

func TestNewSystemExclusiveEvent(t *testing.T) {
	event, err := NewSystemExclusiveEvent(nil, []byte{0x11, 0x12, 0x13, 0x14})
	if err != nil {
		t.Fatal(err)
	}
	expected := []byte{0x11, 0x12, 0x13, 0x14}
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
