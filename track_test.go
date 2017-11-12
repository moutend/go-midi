package midi

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestTrack_Serialize(t *testing.T) {
	event1, _ := NewTextEvent(nil, []byte("txt1"))
	event2, _ := NewTextEvent(nil, []byte("txt2"))
	event3, _ := NewTextEvent(nil, []byte("txt3"))
	track := &Track{
		Events: []Event{
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

func TestParseTrack(t *testing.T) {
	textEvent1 := []byte{0x00, 0xff, 0x01, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x31}
	textEvent2 := []byte{0x00, 0xff, 0x01, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32}
	endOfTrackEvent := []byte{0x00, 0xff, 0x2f, 0x00}

	stream := []byte{}
	stream = append(stream, textEvent1...)
	stream = append(stream, textEvent2...)
	stream = append(stream, endOfTrackEvent...)

	track, err := parseTrack(stream)
	if err != nil {
		t.Fatal(err)
	}
	if len(track.Events) != 3 {
		t.Fatalf("number of events must be 3")
	}
	for i, event := range track.Events {
		switch i {
		case 0:
			expectedText := "text event1"
			actualText := string(event.(*TextEvent).Text())
			if expectedText != actualText {
				t.Fatalf("expected: %v actual: %v", expectedText, actualText)
			}
		case 1:
			expectedText := "text event2"
			actualText := string(event.(*TextEvent).Text())
			if expectedText != actualText {
				t.Fatalf("expected: %v actual: %v", expectedText, actualText)
			}
		case 2:
			switch event.(type) {
			case *EndOfTrackEvent:
				break
			default:
				t.Fatalf("type of event must be EndOfTrackEvent")
			}
		}
	}
}

func TestParseTracks(t *testing.T) {
	pathToMid := filepath.Join("testdata", "vegetable_valley.mid")
	file, err := ioutil.ReadFile(pathToMid)
	if err != nil {
		t.Fatal(err)
	}
	tracks, err := parseTracks(file[14:], 18)
	if err != nil {
		t.Fatal(err)
	}
	if len(tracks) != 18 {
		t.Fatalf("number of tracks must be 18, but got %v", len(tracks))
	}
}
