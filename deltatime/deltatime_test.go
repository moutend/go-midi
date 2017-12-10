package deltatime

import "testing"

func TestNew(t *testing.T) {
	_, err := New(0xffffffff)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	deltaTime, err := New(0x0fffffff)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint32(0x0fffffff)
	actual := deltaTime.Quantity().Uint32()

	if expected != actual {
		t.Fatalf("expected: 0x%x actual: 0x%x", expected, actual)
	}
}

func TestParse(t *testing.T) {
	invalidStreams := [][]byte{
		[]byte{},
		[]byte{0x80},
		[]byte{0x80, 0x80},
		[]byte{0x80, 0x80, 0x80},
		[]byte{0x80, 0x80, 0x80, 0x80},
		[]byte{0x80, 0x80, 0x80, 0x80, 0x80},
	}
	for _, stream := range invalidStreams {
		if _, err := Parse(stream); err == nil {
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
		deltaTime, err := Parse(stream.value)
		if err != nil {
			t.Fatal(err)
		}
		expected := stream.expected
		actual := deltaTime.Quantity().Value()
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
