package midi

import "testing"

func TestNewDeltaTime(t *testing.T) {
	_, err := NewDeltaTime(0xffffffff)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewDeltaTime(0x0fffffff)
	if err != nil {
		t.Fatal(err)
	}
	if event.Quantity().Uint32() != 0x0fffffff {
		t.Fatalf("expected: 0x0fffffff actual: 0x%x", event.Quantity().Uint32())
	}
}
