package midi

import "testing"

func TestAlienEvent_DeltaTime(t *testing.T) {
	event := &AlienEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestAlienEvent_String(t *testing.T) {
	event, err := NewAlienEvent(nil, 0xaa, []byte{0x12, 0x34, 0x56})
	if err != nil {
		t.Fatal(err)
	}

	expected := "&AlienEvent{metaEventType: 0xaa, data: 3 bytes}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestAlienEvent_Serialize(t *testing.T) {
	event, err := NewAlienEvent(nil, 0xaa, []byte{0x12, 0x34, 0x56})
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0xff, 0xaa, 0x03, 0x12, 0x34, 0x56}
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

func TestAlienEvent_SetData(t *testing.T) {
	event := &AlienEvent{}
	err := event.SetData(bigdata)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetData(bigdata[1:])
	if err != nil {
		t.Fatal(err)
	}
}

func TestAlienEvent_Data(t *testing.T) {
	event := &AlienEvent{data: []byte{0x12, 0x34, 0x56}}

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

func TestNewAlienEvent(t *testing.T) {
	event, err := NewAlienEvent(nil, 0xaa, []byte{0x12, 0x34, 0x56})
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
