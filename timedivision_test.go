package midi

import "testing"

func TestSetBPM(t *testing.T) {
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
