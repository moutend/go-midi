package midi

import "testing"

func TestTimeDivision_SetBPM(t *testing.T) {
	td := &TimeDivision{}

	if err := td.SetBPM(480); err != nil {
		t.Fatal(err)
	}
	if err := td.SetBPM(0x8000); err == nil {
		t.Fatalf("err should not be nil")
	}
	if err := td.SetBPM(0x8001); err == nil {
		t.Fatalf("err should not be nil")
	}
}

func TestTimeDivision_Serialize(t *testing.T) {
	td := &TimeDivision{value: 480}
	expected := []byte{0x01, 0xe0}
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
}
