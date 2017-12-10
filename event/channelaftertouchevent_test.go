package event

import "testing"

func TestChannelAfterTouchEventDeltaTime(t *testing.T) {
	event := &ChannelAfterTouchEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("deltatime.DeltaTime() don't return nil")
	}
}

func TestChannelAfterTouchEvent_String(t *testing.T) {
	event, err := NewChannelAfterTouchEvent(nil, 1, 123)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&ChannelAfterTouchEvent{channel: 1, velocity: 123}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestChannelAfterTouchEvent_Serialize(t *testing.T) {
	event, err := NewChannelAfterTouchEvent(nil, 1, 123)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0xd1, 0x7b}
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

func TestChannelAfterTouchEvent_SetChannel(t *testing.T) {
	event := &ChannelAfterTouchEvent{}

	err := event.SetChannel(0x10)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetChannel(0x0f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChannelAfterTouchEvent_Channel(t *testing.T) {
	event := &ChannelAfterTouchEvent{channel: 1}

	expected := uint8(1)
	actual := event.Channel()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestChannelAfterTouchEvent_SetVelocity(t *testing.T) {
	event := &ChannelAfterTouchEvent{}

	err := event.SetVelocity(0x80)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetVelocity(0x7f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChannelAfterTouchEvent_Velocity(t *testing.T) {
	event := &ChannelAfterTouchEvent{velocity: 1}

	expected := uint8(1)
	actual := event.Velocity()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewChannelAfterTouchEvent(t *testing.T) {
	_, err := NewChannelAfterTouchEvent(nil, 255, 127)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	_, err = NewChannelAfterTouchEvent(nil, 15, 255)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewChannelAfterTouchEvent(nil, 15, 127)
	if err != nil {
		t.Fatal(err)
	}
	if event.channel != 15 {
		t.Fatalf("expected: 15 actual: %v", event.channel)
	}
	if event.velocity != 127 {
		t.Fatalf("expected: 127 actual: %v", event.velocity)
	}
}
