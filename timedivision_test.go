package midi

import "testing"

func TestTimeDivision_String(t *testing.T) {
	td := &TimeDivision{}

	expected := "&TimeDivision{bpm: 120}"
	actual := td.String()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}

	td = &TimeDivision{value: 32767}

	expected = "&TimeDivision{bpm: 32767}"
	actual = td.String()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}

	td = &TimeDivision{value: 0x980a}

	expected = "&TimeDivision{frames: 24, ticks: 10}"
	actual = td.String()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestTimeDivision_Serialize(t *testing.T) {
	td := &TimeDivision{}

	expected := []byte{0x00, 0x78}
	actual := td.Serialize()

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v actual: %v", len(expected), len(actual))
	}
	for i, e := range expected {
		a := actual[i]
		if e != a {
			t.Fatalf("expected[%v] = %v actual[%v] = %v", i, e, i, a)
		}
	}

	td = &TimeDivision{value: 480}

	expected = []byte{0x01, 0xe0}
	actual = td.Serialize()

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v actual: %v", len(expected), len(actual))
	}
	for i, e := range expected {
		a := actual[i]
		if e != a {
			t.Fatalf("expected[%v] = %v actual[%v] = %v", i, e, i, a)
		}
	}
}
func TestTimeDivision_SetBPM(t *testing.T) {
	td := &TimeDivision{}

	if err := td.SetBPM(0x8000); err == nil {
		t.Fatalf("err should not be nil")
	}
	if err := td.SetBPM(480); err != nil {
		t.Fatal(err)
	}

	expected := uint16(480)
	actual := td.value

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestTimeDivision_BPM(t *testing.T) {
	td := &TimeDivision{value: 32768}

	_, err := td.BPM()
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	td = &TimeDivision{}

	expected := uint16(120)
	actual, _ := td.BPM()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}

	td = &TimeDivision{value: 32767}

	expected = uint16(32767)
	actual, _ = td.BPM()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestTimeDivision_SetFPS(t *testing.T) {
	td := &TimeDivision{}
	td.SetFPS(24, 10)

	expected := uint16(0x980a)
	actual := td.value

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestTimeDivision_FPS(t *testing.T) {
	td := &TimeDivision{}
	_, _, err := td.FPS()
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	td = &TimeDivision{value: 0x980a}

	expectedFrames, expectedTicks := uint16(24), uint16(10)
	actualFrames, actualTicks, _ := td.FPS()

	if expectedFrames != actualFrames {
		t.Fatalf("expected: %v actual: %v", expectedFrames, actualFrames)
	}
	if expectedTicks != actualTicks {
		t.Fatalf("expected: %v actual: %v", expectedTicks, actualTicks)
	}
}
