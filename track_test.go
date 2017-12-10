package midi

import (
	"testing"

	"github.com/moutend/go-midi/event"
)

func TestTrack_Serialize(t *testing.T) {
	event1, _ := event.NewTextEvent(nil, []byte("txt1"))
	event2, _ := event.NewTextEvent(nil, []byte("txt2"))
	event3, _ := event.NewTextEvent(nil, []byte("txt3"))
	track := &Track{
		Events: []event.Event{
			event1,
			event2,
			event3,
		},
	}

	expected := []byte{0x4d, 0x54, 0x72, 0x6B, 0x00, 0x00, 0x00, 0x18, 0x00, 0xff, 0x01, 0x04, 0x74, 0x78, 0x74, 0x31, 0x00, 0xff, 0x01, 0x04, 0x74, 0x78, 0x74, 0x32, 0x00, 0xff, 0x01, 0x04, 0x74, 0x78, 0x74, 0x33}
	actual := track.Serialize()

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
