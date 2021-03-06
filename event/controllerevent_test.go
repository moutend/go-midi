package event

import (
	"testing"
)

func TestControllerEventDeltaTime(t *testing.T) {
	event := &ControllerEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("deltatime.DeltaTime() don't return nil")
	}
}

func TestControllerEvent_String(t *testing.T) {
	event, err := NewControllerEvent(nil, 1, 12, 34)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&ControllerEvent{channel: 1, control: EffectControl1, value: 34}"
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

	expected := []byte{0xb0, 0x0c, 0x22}
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

func TestControllerEvent_SetChannel(t *testing.T) {
	event := &ControllerEvent{}

	err := event.SetChannel(0x10)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetChannel(0x0f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestControllerEvent_Channel(t *testing.T) {
	event := &ControllerEvent{channel: 1}

	expected := uint8(1)
	actual := event.Channel()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestControllerEvent_SetControl(t *testing.T) {
	event := &ControllerEvent{}

	err := event.SetControl(0x80)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetControl(0x7f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestControllerEvent_Control(t *testing.T) {
	event := &ControllerEvent{control: 1}

	expected := uint8(1)
	actual := uint8(event.Control())

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestControllerEvent_SetValue(t *testing.T) {
	event := &ControllerEvent{}

	err := event.SetValue(0x80)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetValue(0x7f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestControllerEvent_Value(t *testing.T) {
	event := &ControllerEvent{value: 1}

	expected := uint8(1)
	actual := event.Value()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewControllerEvent(t *testing.T) {
	_, err := NewControllerEvent(nil, 255, 127, 127)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	_, err = NewControllerEvent(nil, 15, 255, 127)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	_, err = NewControllerEvent(nil, 15, 127, 255)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewControllerEvent(nil, 15, 127, 127)
	if err != nil {
		t.Fatal(err)
	}
	if event.channel != 15 {
		t.Fatalf("expected: 15 actual: %v", event.channel)
	}
	if event.control != 127 {
		t.Fatalf("expected: 127 actual: %v", event.control)
	}
}
