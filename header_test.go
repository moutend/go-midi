package midi

import (
	"bytes"
	"encoding/binary"
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
	data := h.Serialize()
	// Header Length of the header is always 14 bytes.
	if len(data) != 14 {
		t.Fatalf("expected: 14 actual: %v", len(data))
	}

	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, []byte("MThd"))
	binary.Write(buf, binary.BigEndian, uint32(6))
	binary.Write(buf, binary.BigEndian, uint16(1))
	binary.Write(buf, binary.BigEndian, uint16(18))
	binary.Write(buf, binary.BigEndian, uint16(480))

	for i, b := range buf.Bytes() {
		if data[i] != b {
			t.Fatalf("expected: data[%v] = %v actual: data[%v] = %v", i, b, i, data[i])
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
