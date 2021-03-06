package event

import "testing"

func TestAlienEvent_DeltaTime(t *testing.T) {
	event := &AlienEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("deltatime.DeltaTime() don't return nil")
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

	expected := []byte{0xff, 0xaa, 0x03, 0x12, 0x34, 0x56}
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

func TestAlienEvent_SetMetaEventType(t *testing.T) {
	event := &AlienEvent{}
	event.SetMetaEventType(0x12)

	expected := uint8(0x12)
	actual := event.metaEventType

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestAlienEvent_MetaEventType(t *testing.T) {
	event := &AlienEvent{metaEventType: 0x12}

	expected := uint8(0x12)
	actual := event.MetaEventType()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
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
	event := &AlienEvent{}

	expected := []byte{}
	actual := event.Data()

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v actual: %v", len(expected), len(actual))
	}

	event = &AlienEvent{data: []byte{0x12, 0x34, 0x56}}

	expected = []byte{0x12, 0x34, 0x56}
	actual = event.Data()

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
	event, err := NewAlienEvent(nil, 0x12, bigdata)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err = NewAlienEvent(nil, 0x12, []byte{0x12, 0x34, 0x56})
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
