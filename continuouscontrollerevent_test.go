package midi

import "testing"

func TestContinuousControllerEvent_DeltaTime(t *testing.T) {
	event := &ContinuousControllerEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestContinuousControllerEvent_String(t *testing.T) {
	event, err := NewContinuousControllerEvent(nil, 12, 34)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&ContinuousControllerEvent{control: 12, value: 34}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestContinuousControllerEvent_Serialize(t *testing.T) {
	event, err := NewContinuousControllerEvent(nil, 12, 34)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0x0c, 0x22}
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

func TestContinuousControllerEvent_SetControl(t *testing.T) {
	event := &ContinuousControllerEvent{}

	err := event.SetControl(0x80)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetControl(0x7f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestContinuousControllerEvent_Control(t *testing.T) {
	event := &ContinuousControllerEvent{control: 1}

	expected := uint8(1)
	actual := event.Control()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestContinuousControllerEvent_SetValue(t *testing.T) {
	event := &ContinuousControllerEvent{}

	err := event.SetValue(0x80)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetValue(0x7f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestContinuousControllerEvent_Value(t *testing.T) {
	event := &ContinuousControllerEvent{value: 1}

	expected := uint8(1)
	actual := event.Value()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewContinuousControllerEvent(t *testing.T) {
	_, err := NewContinuousControllerEvent(nil, 127, 255)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	_, err = NewContinuousControllerEvent(nil, 255, 127)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewContinuousControllerEvent(nil, 127, 127)
	if err != nil {
		t.Fatal(err)
	}
	if event.control != 127 {
		t.Fatalf("expected: 127 actual: %v", event.control)
	}
}
