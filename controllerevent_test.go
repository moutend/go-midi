package midi

import "testing"

func TestControllerEvent_DeltaTime(t *testing.T) {
	event := &ControllerEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestControllerEvent_String(t *testing.T) {
	event, err := NewControllerEvent(nil, 1, 12, 34)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&ControllerEvent{channel: 1, control: 12, value: 34}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestControllerEvent_Serialize(t *testing.T) {
	event, err := NewControllerEvent(nil, 0, 12, 34)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0xb0, 0x0c, 0x22}
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
