package midi

import "testing"

func TestParseQuantity(t *testing.T) {
	invalidStreams := [][]byte{
		[]byte{},
		[]byte{0x80},
		[]byte{0x80, 0x80},
		[]byte{0x80, 0x80, 0x80},
		[]byte{0x80, 0x80, 0x80, 0x80},
		[]byte{0x80, 0x80, 0x80, 0x80, 0x80},
	}
	for _, stream := range invalidStreams {
		if _, err := parseQuantity(stream); err == nil {
			t.Fatalf("err should not be nil (stream=%+v)", stream)
		}
	}

	validStreams := []struct {
		value    []byte
		expected []byte
	}{
		{
			value:    []byte{0x00},
			expected: []byte{0x00},
		},
		{
			value:    []byte{0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			expected: []byte{0x00},
		},
		{
			value:    []byte{0x7f},
			expected: []byte{0x7f},
		},
		{
			value:    []byte{0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			expected: []byte{0x7f},
		},
		{
			value:    []byte{0x80, 0x7f},
			expected: []byte{0x80, 0x7f},
		},
		{
			value:    []byte{0x80, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			expected: []byte{0x80, 0x7f},
		},
		{
			value:    []byte{0x80, 0x80, 0x7f},
			expected: []byte{0x80, 0x80, 0x7f},
		},
		{
			value:    []byte{0x80, 0x80, 0x7f, 0xff, 0xff, 0xff, 0xff, 0xff},
			expected: []byte{0x80, 0x80, 0x7f},
		},
		{
			value:    []byte{0x80, 0x80, 0x80, 0x7f},
			expected: []byte{0x80, 0x80, 0x80, 0x7f},
		},
		{
			value:    []byte{0x80, 0x80, 0x80, 0x7f, 0xff, 0xff, 0xff, 0xff},
			expected: []byte{0x80, 0x80, 0x80, 0x7f},
		},
	}
	for _, stream := range validStreams {
		q, err := parseQuantity(stream.value)
		if err != nil {
			t.Fatal(err)
		}
		expected := stream.expected
		actual := q.value
		if len(expected) != len(actual) {
			t.Fatalf("expected:%+v actual: %+v", expected, actual)
		}
		for i, a := range actual {
			if a != expected[i] {
				t.Fatalf("expected: stream[%v] = %v actual: stream[%v] = %v", i, expected[i], i, a)
			}
		}
	}
}

func TestQuantity_Uint32(t *testing.T) {
	var u32 uint32
	q := &Quantity{}

	q.value = []byte{0xff, 0xff, 0xff, 0x7f}
	u32 = q.Uint32()
	if u32 != 0x0fffffff {
		t.Fatalf("expected: 0x0fffffff actual: 0x%x", u32)
	}

	q.value = []byte{0xff, 0xff, 0x7f}
	u32 = q.Uint32()
	if u32 != 0x1fffff {
		t.Fatalf("expected: 0x1fffff actual: 0x%x", u32)
	}

	q.value = []byte{0xff, 0x7f}
	u32 = q.Uint32()
	if u32 != 0x3fff {
		t.Fatalf("expected: 0x3fff actual: 0x%x", u32)
	}

	q.value = []byte{0x7f}
	u32 = q.Uint32()
	if u32 != 0x7f {
		t.Fatalf("expected: 0x7f actual: 0x%x", u32)
	}
}
