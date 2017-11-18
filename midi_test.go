package midi

import (
	"io/ioutil"
	"testing"
)

func TestMIDI_Serialize(t *testing.T) {
	for _, pathToMid := range pathsToMid {
		file, err := ioutil.ReadFile(pathToMid)
		if err != nil {
			t.Fatal(err)
		}

		m, err := Parse(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := file
		actual := m.Serialize()

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
}

func TestMIDI_TimeDivision(t *testing.T) {
	m := &MIDI{}
	td := m.TimeDivision()
	if td == nil {
		t.Fatalf("TimeDivision must not return nil")
	}
}

func TestParse(t *testing.T) {
	for _, pathToMid := range pathsToMid {
		file, err := ioutil.ReadFile(pathToMid)
		if err != nil {
			t.Fatal(err)
		}
		_, err = Parse(file)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestParseHeader(t *testing.T) {
	for _, pathToMid := range pathsToMid {
		file, err := ioutil.ReadFile(pathToMid)
		if err != nil {
			t.Fatal(err)
		}
		formatType, numberOfTracks, timeDivision, err := parseHeader(file)
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
