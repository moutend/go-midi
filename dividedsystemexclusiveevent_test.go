package midi

import "testing"

func TestDividedSystemExclusiveEvent_DeltaTime(t *testing.T) {
	event := &DividedSystemExclusiveEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestDividedSystemExclusiveEvent_String(t *testing.T) {
	event, err := NewDividedSystemExclusiveEvent(nil, []byte{0x01, 0x02, 0x03, 0x04})
	if err != nil {
		t.Fatal(err)
	}

	expected := "&DividedSystemExclusiveEvent{data: 4 bytes}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestDividedSystemExclusiveEvent_Serialize(t *testing.T) {
	event, err := NewDividedSystemExclusiveEvent(nil, []byte{0x11, 0x12, 0x13, 0x14})
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0xf7, 0x04, 0x11, 0x12, 0x13, 0x14}
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

func TestDividedSystemExclusiveEvent_SetData(t *testing.T) {
	event := &DividedSystemExclusiveEvent{}

	err := event.SetData(bigdata)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetData(bigdata[1:])
	if err != nil {
		t.Fatal(err)
	}
}

func TestDividedSystemExclusiveEvent_Data(t *testing.T) {
	event := &DividedSystemExclusiveEvent{}
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

func TestNewDividedSystemExclusiveEvent(t *testing.T) {
	event, err := NewDividedSystemExclusiveEvent(nil, []byte{0x11, 0x12, 0x13, 0x14})
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
