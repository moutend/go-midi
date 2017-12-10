package event

import "testing"

func TestTimeSignatureEventDeltaTime(t *testing.T) {
	event := &TimeSignatureEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("deltatime.DeltaTime() don't return nil")
	}
}

func TestTimeSignatureEvent_String(t *testing.T) {
	event, err := NewTimeSignatureEvent(nil, 123, 123, 123, 123)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&TimeSignatureEvent{numerator: 123, denominator: 123, metronomePulse: 123, quarterNote: 123}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestTimeSignatureEvent_Serialize(t *testing.T) {
	event, err := NewTimeSignatureEvent(nil, 1, 2, 3, 4)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0xff, 0x58, 0x04, 0x01, 0x02, 0x03, 0x04}
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

func TestTimeSignatureEvent_SetNumerator(t *testing.T) {
	event := &TimeSignatureEvent{}
	err := event.SetNumerator(0xff)
	if err != nil {
		t.Fatalf("err must be always nil")
	}
}

func TestTimeSignatureEvent_Numerator(t *testing.T) {
	event := &TimeSignatureEvent{numerator: 123}

	expected := uint8(123)
	actual := event.Numerator()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestTimeSignatureEvent_SetDenominator(t *testing.T) {
	event := &TimeSignatureEvent{}
	err := event.SetDenominator(0xff)
	if err != nil {
		t.Fatalf("err must be always nil")
	}
}

func TestTimeSignatureEvent_Denominator(t *testing.T) {
	event := &TimeSignatureEvent{denominator: 123}

	expected := uint8(123)
	actual := event.Denominator()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestTimeSignatureEvent_SetMetronomePulse(t *testing.T) {
	event := &TimeSignatureEvent{}
	err := event.SetMetronomePulse(0xff)
	if err != nil {
		t.Fatalf("err must be always nil")
	}
}

func TestTimeSignatureEvent_MetronomePulse(t *testing.T) {
	event := &TimeSignatureEvent{metronomePulse: 123}

	expected := uint8(123)
	actual := event.MetronomePulse()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestTimeSignatureEvent_SetQuarterNote(t *testing.T) {
	event := &TimeSignatureEvent{}
	err := event.SetQuarterNote(0xff)
	if err != nil {
		t.Fatalf("err must be always nil")
	}
}

func TestTimeSignatureEvent_QuarterNote(t *testing.T) {
	event := &TimeSignatureEvent{quarterNote: 123}

	expected := uint8(123)
	actual := event.QuarterNote()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewTimeSignatureEvent(t *testing.T) {
	event, err := NewTimeSignatureEvent(nil, 1, 2, 3, 4)
	if err != nil {
		t.Fatal(err)
	}
	if event.numerator != 1 {
		t.Fatalf("expected: 1 actual: %v", event.numerator)
	}
	if event.denominator != 2 {
		t.Fatalf("expected: 2 actual: %v", event.denominator)
	}
	if event.metronomePulse != 3 {
		t.Fatalf("expected: 3 actual: %v", event.metronomePulse)
	}
	if event.quarterNote != 4 {
		t.Fatalf("expected: 4 actual: %v", event.quarterNote)
	}
}
