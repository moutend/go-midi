package event

import "testing"

func TestCopyrightNoticeEventDeltaTime(t *testing.T) {
	event := &CopyrightNoticeEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("deltatime.DeltaTime() don't return nil")
	}
}

func TestCopyrightNoticeEvent_String(t *testing.T) {
	event, err := NewCopyrightNoticeEvent(nil, []byte("text"))
	if err != nil {
		t.Fatal(err)
	}

	expected := "&CopyrightNoticeEvent{text: \"text\"}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestCopyrightNoticeEvent_Serialize(t *testing.T) {
	event, err := NewCopyrightNoticeEvent(nil, []byte("text"))
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0xff, 0x02, 0x04, 0x74, 0x65, 0x78, 0x74}
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

func TestCopyrightNoticeEvent_SetText(t *testing.T) {
	event := &CopyrightNoticeEvent{}

	err := event.SetText(bigdata)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetText(bigdata[1:])
	if err != nil {
		t.Fatal(err)
	}
}

func TestCopyrightNoticeEvent_Text(t *testing.T) {
	event := &CopyrightNoticeEvent{}

	expected := ""
	actual := string(event.Text())

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}

	event = &CopyrightNoticeEvent{text: []byte("text")}

	expected = "text"
	actual = string(event.Text())

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewCopyrightNoticeEvent(t *testing.T) {
	_, err := NewCopyrightNoticeEvent(nil, bigdata)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewCopyrightNoticeEvent(nil, []byte("text"))
	if err != nil {
		t.Fatal(err)
	}
	expected := []byte("text")
	actual := event.text

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
