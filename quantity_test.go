package midi

import "testing"

func TestQuantity_SetUint32(t *testing.T) {
	q := &Quantity{}

	var err error
	var expected, actual []byte

	err = q.SetUint32(0xffffffff)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = q.SetUint32(0xfffffff)
	if err != nil {
		t.Fatal(err)
	}

	expected = []byte{0xff, 0xff, 0xff, 0x7f}
	actual = q.value

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v bytes actual: %v bytes", len(expected), len(actual))
	}
	for i, e := range expected {
		a := actual[i]
		if e != a {
			t.Fatalf("expected[%v] = 0x%x actual[%v] = 0x%x", i, e, i, a)
		}
	}

	err = q.SetUint32(0x3fffff)
	if err != nil {
		t.Fatal(err)
	}

	expected = []byte{0x81, 0xff, 0xff, 0x7f}
	actual = q.value

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v bytes actual: %v bytes", len(expected), len(actual))
	}
	for i, e := range expected {
		a := actual[i]
		if e != a {
			t.Fatalf("expected[%v] = 0x%x actual[%v] = 0x%x", i, e, i, a)
		}
	}

	err = q.SetUint32(0x1fffff)
	if err != nil {
		t.Fatal(err)
	}

	expected = []byte{0xff, 0xff, 0x7f}
	actual = q.value

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v bytes actual: %v bytes", len(expected), len(actual))
	}
	for i, e := range expected {
		a := actual[i]
		if e != a {
			t.Fatalf("expected[%v] = 0x%x actual[%v] = 0x%x", i, e, i, a)
		}
	}

	err = q.SetUint32(0x7fff)
	if err != nil {
		t.Fatal(err)
	}

	expected = []byte{0x81, 0xff, 0x7f}
	actual = q.value

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v bytes actual: %v bytes", len(expected), len(actual))
	}
	for i, e := range expected {
		a := actual[i]
		if e != a {
			t.Fatalf("expected[%v] = 0x%x actual[%v] = 0x%x", i, e, i, a)
		}
	}

	err = q.SetUint32(0x3fff)
	if err != nil {
		t.Fatal(err)
	}

	expected = []byte{0xff, 0x7f}
	actual = q.value

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v bytes actual: %v bytes", len(expected), len(actual))
	}
	for i, e := range expected {
		a := actual[i]
		if e != a {
			t.Fatalf("expected[%v] = 0x%x actual[%v] = 0x%x", i, e, i, a)
		}
	}

	err = q.SetUint32(0xff)
	if err != nil {
		t.Fatal(err)
	}

	expected = []byte{0x81, 0x7f}
	actual = q.value

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v bytes actual: %v bytes", len(expected), len(actual))
	}
	for i, e := range expected {
		a := actual[i]
		if e != a {
			t.Fatalf("expected[%v] = 0x%x actual[%v] = 0x%x", i, e, i, a)
		}
	}

	err = q.SetUint32(0x7f)
	if err != nil {
		t.Fatal(err)
	}

	expected = []byte{0x7f}
	actual = q.value

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

func TestQuantity_Uint32(t *testing.T) {
	var expected, actual uint32

	q := &Quantity{}

	q.value = []byte{0xff, 0xff, 0xff, 0x7f}
	expected = uint32(0xfffffff)
	actual = q.Uint32()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}

	q.value = []byte{0x81, 0xff, 0xff, 0x7f}
	expected = uint32(0x3fffff)
	actual = q.Uint32()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}

	q.value = []byte{0xff, 0xff, 0x7f}
	expected = uint32(0x1fffff)
	actual = q.Uint32()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}

	q.value = []byte{0x81, 0xff, 0x7f}
	expected = uint32(0x7fff)
	actual = q.Uint32()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}

	q.value = []byte{0xff, 0x7f}
	expected = uint32(0x3fff)
	actual = q.Uint32()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}

	q.value = []byte{0x81, 0x7f}
	expected = uint32(0xff)
	actual = q.Uint32()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}

	q.value = []byte{0x7f}
	expected = uint32(0x7f)
	actual = q.Uint32()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestQuantity_SetValue(t *testing.T) {
	q := &Quantity{}

	err := q.SetValue([]byte{0x12, 0x34, 0x56, 0x78, 0x90})
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = q.SetValue([]byte{0x12, 0x34, 0x56, 0x78})
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x12, 0x34, 0x56, 0x78}
	actual := q.value

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

func TestQuantity_Value(t *testing.T) {
	q := &Quantity{}

	expected := []byte{0x0}
	actual := q.Value()

	if actual == nil {
		t.Fatalf("value must not be nil")
	}
	if len(expected) != len(actual) {
		t.Fatalf("expected: %v bytes actual: %v bytes", len(expected), len(actual))
	}
	if expected[0] != actual[0] {
		t.Fatalf("expected[%v] = 0x%x actual[%v] = 0x%x", 0, expected[0], 0, actual[0])
	}

	q = &Quantity{[]byte{0x12, 0x34, 0x56, 0x78}}

	expected = []byte{0x12, 0x34, 0x56, 0x78}
	actual = q.Value()

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
