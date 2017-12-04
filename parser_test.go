package midi

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	for _, pathToMid := range pathsToMid {
		file, err := ioutil.ReadFile(pathToMid)
		if err != nil {
			t.Fatal(err)
		}
		_, err = NewParser(file).Parse(file)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestParser_parseHeader(t *testing.T) {
	for _, pathToMid := range pathsToMid {
		file, err := ioutil.ReadFile(pathToMid)
		if err != nil {
			t.Fatal(err)
		}
		formatType, numberOfTracks, timeDivision, err := NewParser(file).parseHeader()
		if err != nil {
			t.Fatal(err)
		}
		if formatType != 1 {
			t.Fatalf("expected: 1 actual: %v", formatType)
		}
		if numberOfTracks != 18 {
			t.Fatalf("expected: 18 actual: %v", numberOfTracks)
		}
		if timeDivision != 480 {
			t.Fatalf("expected: 480 actual: %v", timeDivision)
		}
	}
}

func TestParser_parseTracks(t *testing.T) {
	pathToMid := filepath.Join("testdata", "vegetable_valley.mid")
	file, err := ioutil.ReadFile(pathToMid)
	if err != nil {
		t.Fatal(err)
	}
	tracks, err := NewParser(file[14:]).parseTracks(18)
	if err != nil {
		t.Fatal(err)
	}
	if len(tracks) != 18 {
		t.Fatalf("number of tracks must be 18, but got %v", len(tracks))
	}
}

func TestParser_parseTrack(t *testing.T) {
	textEvent1 := []byte{0x00, 0xff, 0x01, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x31}
	textEvent2 := []byte{0x00, 0xff, 0x01, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32}
	endOfTrackEvent := []byte{0x00, 0xff, 0x2f, 0x00}

	stream := []byte{}
	stream = append(stream, textEvent1...)
	stream = append(stream, textEvent2...)
	stream = append(stream, endOfTrackEvent...)

	track, err := NewParser(stream).parseTrack()
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

func TestParser_parseEvent(t *testing.T) {
	stream := []byte{0x00, 0xff, 0x02, 0x12, 0x43, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x20, 0x28, 0x43, 0x29, 0x20, 0x32, 0x30, 0x31, 0x37}
	event, err := NewParser(stream).parseEvent()
	if err != nil {
		t.Fatal(err)
	}

	sizeOfEvent := len(event.Serialize())
	if sizeOfEvent != 22 {
		t.Fatalf("expected: size of event = 22, actual: size of event = %v", sizeOfEvent)
	}
	if event.DeltaTime().Quantity().Uint32() != 0 {
		t.Fatalf("expected: 0 actual: %v", event.DeltaTime().Quantity().Uint32())
	}
	switch event.(type) {
	case *CopyrightNoticeEvent:
		break
	default:
		t.Fatalf("type of event must be CopyrightNoticeEvent")
	}

	expectedText := stream[4:]
	actualText := event.(*CopyrightNoticeEvent).text

	if len(expectedText) != len(actualText) {
		t.Fatalf("expect: len(event.(*CopyrightNoticeEvent).text) = %v actual: len(event.(*CopyrightNoticeEvent).text) = %v", len(expectedText), len(actualText))
	}
	for i, v := range expectedText {
		if v != actualText[i] {
			t.Fatalf("expected: text[%v] = %v actual: text[%v] = %v", i, v, i, actualText[i])
		}
	}
}
