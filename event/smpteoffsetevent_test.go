package event

import "testing"

func TestSMPTEOffsetEventDeltaTime(t *testing.T) {
	event := &SMPTEOffsetEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("deltatime.DeltaTime() don't return nil")
	}
}

func TestSMPTEOffsetEvent_String(t *testing.T) {
	event, err := NewSMPTEOffsetEvent(nil, 23, 59, 59, 30, 99)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&SMPTEOffsetEvent{hour: 23, minute: 59, second: 59, frame: 30, subFrame: 99}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestSMPTEOffsetEvent_Serialize(t *testing.T) {
	event, err := NewSMPTEOffsetEvent(nil, 23, 59, 59, 30, 99)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0xff, 0x54, 0x05, 0x17, 0x3b, 0x3b, 0x1e, 0x63}
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

func TestSMPTEOffsetEvent_SetHour(t *testing.T) {
	event := &SMPTEOffsetEvent{}

	err := event.SetHour(24)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetHour(23)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSMPTEOffsetEvent_Hour(t *testing.T) {
	event := &SMPTEOffsetEvent{hour: 23}

	expected := uint8(23)
	actual := event.Hour()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestSMPTEOffsetEvent_SetMinute(t *testing.T) {
	event := &SMPTEOffsetEvent{}

	err := event.SetMinute(60)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetMinute(59)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSMPTEOffsetEvent_Minute(t *testing.T) {
	event := &SMPTEOffsetEvent{minute: 59}

	expected := uint8(59)
	actual := event.Minute()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestSMPTEOffsetEvent_SetSecond(t *testing.T) {
	event := &SMPTEOffsetEvent{}

	err := event.SetSecond(60)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetSecond(59)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSMPTEOffsetEvent_Second(t *testing.T) {
	event := &SMPTEOffsetEvent{second: 59}

	expected := uint8(59)
	actual := event.Second()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestSMPTEOffsetEvent_SetFrame(t *testing.T) {
	event := &SMPTEOffsetEvent{}

	err := event.SetFrame(31)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetFrame(30)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSMPTEOffsetEvent_Frame(t *testing.T) {
	event := &SMPTEOffsetEvent{frame: 30}

	expected := uint8(30)
	actual := event.Frame()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestSMPTEOffsetEvent_SetSubFrame(t *testing.T) {
	event := &SMPTEOffsetEvent{}

	err := event.SetSubFrame(100)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetSubFrame(99)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSMPTEOffsetEvent_SubFrame(t *testing.T) {
	event := &SMPTEOffsetEvent{subFrame: 99}

	expected := uint8(99)
	actual := event.SubFrame()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewSMPTEOffsetEvent(t *testing.T) {
	event, err := NewSMPTEOffsetEvent(nil, 23, 59, 59, 30, 99)
	if err != nil {
		t.Fatal(err)
	}
	if event.hour != 23 {
		t.Fatalf("expected: 23 actual: %v", event.hour)
	}
	if event.minute != 59 {
		t.Fatalf("expected: 59 actual: %v", event.minute)
	}
	if event.second != 59 {
		t.Fatalf("expected: 59 actual: %v", event.second)
	}
	if event.frame != 30 {
		t.Fatalf("expected: 30 actual: %v", event.frame)
	}
	if event.subFrame != 99 {
		t.Fatalf("expected: 99 actual: %v", event.subFrame)
	}
}
