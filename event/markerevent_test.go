package event

import "testing"

func TestMarkerEventDeltaTime(t *testing.T) {
	event := &MarkerEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("deltatime.DeltaTime() don't return nil")
	}
}

func TestMarkerEvent_String(t *testing.T) {
	event, err := NewMarkerEvent(nil, []byte("text"))
	if err != nil {
		t.Fatal(err)
	}

	expected := "&MarkerEvent{text: \"text\"}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestMarkerEvent_Serialize(t *testing.T) {
	event, err := NewMarkerEvent(nil, []byte("text"))
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0xff, 0x06, 0x04, 0x74, 0x65, 0x78, 0x74}
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

func TestMarkerEvent_SetText(t *testing.T) {
	event := &MarkerEvent{}

	err := event.SetText(bigdata)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetText(bigdata[1:])
	if err != nil {
		t.Fatal(err)
	}
}

func TestMarkerEvent_Text(t *testing.T) {
	event := &MarkerEvent{}

	expected := ""
	actual := string(event.Text())
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}

	event = &MarkerEvent{text: []byte("text")}

	expected = "text"
	actual = string(event.Text())

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewMarkerEvent(t *testing.T) {
	_, err := NewMarkerEvent(nil, bigdata)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewMarkerEvent(nil, []byte("text"))
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
