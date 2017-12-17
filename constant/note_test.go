package constant

import "testing"

func TestParseNote(t *testing.T) {
	expected1 := C1
	actual1, err := ParseNote("C1")
	if err != nil {
		t.Fatal(err)
	}
	if expected1 != actual1 {
		t.Fatalf("expected: %v actual: %v", expected1, actual1)
	}

	expected2 := Db1
	actual2, err := ParseNote("Db1")
	if err != nil {
		t.Fatal(err)
	}
	if expected2 != actual2 {
		t.Fatalf("expected: %v actual: %v", expected2, actual2)
	}
	expected3 := Cminus2
	actual3, err := ParseNote("C-2")
	if err != nil {
		t.Fatal(err)
	}
	if expected3 != actual3 {
		t.Fatalf("expected: %v actual: %v", expected3, actual3)
	}

}
