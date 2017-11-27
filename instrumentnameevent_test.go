package midi

import "testing"

func TestInstrumentNameEvent_DeltaTime(t *testing.T) {
	event := &InstrumentNameEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestInstrumentNameEvent_String(t *testing.T) {
	event, err := NewInstrumentNameEvent(nil, []byte("text"))
	if err != nil {
		t.Fatal(err)
	}

	expected := "&InstrumentNameEvent{text: \"text\"}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestInstrumentNameEvent_Serialize(t *testing.T) {
	event, err := NewInstrumentNameEvent(nil, []byte("text"))
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0xff, 0x04, 0x04, 0x74, 0x65, 0x78, 0x74}
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

func TestInstrumentNameEvent_SetText(t *testing.T) {
	event := &InstrumentNameEvent{}

	err := event.SetText(bigdata)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetText(bigdata[1:])
	if err != nil {
		t.Fatal(err)
	}
}

func TestInstrumentNameEvent_Text(t *testing.T) {
	event := &InstrumentNameEvent{}

	expected := ""
	actual := string(event.Text())
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}

	event = &InstrumentNameEvent{text: []byte("text")}

	expected = "text"
	actual = string(event.Text())

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewInstrumentNameEvent(t *testing.T) {
	_, err := NewInstrumentNameEvent(nil, bigdata)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewInstrumentNameEvent(nil, []byte("text"))
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
