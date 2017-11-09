package midi

import "testing"

func TestChannelAfterTouchEvent_DeltaTime(t *testing.T) {
	event := &ChannelAfterTouchEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestChannelAfterTouchEvent_String(t *testing.T) {
	event, err := NewChannelAfterTouchEvent(nil, 1, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&ChannelAfterTouchEvent{channel: 1, velocity: 50}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestChannelAfterTouchEvent_Serialize(t *testing.T) {
	event, err := NewChannelAfterTouchEvent(nil, 0, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0xd0, 0x32}
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
