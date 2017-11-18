package midi

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestHeader_Serialize(t *testing.T) {
	td := &TimeDivision{
		value: 480,
	}
	h := &Header{
		formatType:   1,
		tracks:       18,
		timeDivision: td,
	}

	expected := []byte{0x4d, 0x54, 0x68, 0x64, 0x00, 0x00, 0x00, 0x06, 0x00, 0x01, 0x00, 0x12, 0x01, 0xe0}
	actual := h.Serialize()

	if len(actual) != 14 {
		t.Fatalf("expected: 14 actual: %v", len(actual))
	}
	for i, e := range expected {
		a := actual[i]
		if e != a {
			t.Fatalf("expected[%v] = %v actual[%v] = %v", i, e, i, a)
		}
	}
}

func TestParseHeader(t *testing.T) {
	pathToMid := filepath.Join("testdata", "vegetable_valley.mid")
	file, err := ioutil.ReadFile(pathToMid)
	if err != nil {
		t.Fatal(err)
	}
	header, err := parseHeader(file)
	if err != nil {
		t.Fatal(err)
	}
	data := header.Serialize()

	// Header Length of the header is always 14 bytes.
	if len(data) != 14 {
		t.Fatalf("expected: 14 actual: %v", len(data))
	}

	for i, b := range data {
		if file[i] != b {
			t.Fatalf("expected: data[%v] = %v actual: data[%v] = %v", i, file[i], i, b)
		}
	}
}
