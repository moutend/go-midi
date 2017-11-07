package midi

import "testing"

func TestEndOfTrackEvent_Serialize(t *testing.T) {
	event := &EndOfTrackEvent{}
	expected := []byte{0x00, 0xff, 0x2f, 0x00}
	actual := event.Serialize()

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v bytes actual: %v bytes", len(expected), len(actual))
	}
}
