package midi

import (
	"fmt"
)

// Quantity represents variable length quantity in MIDI.
type Quantity struct {
	value []byte
}

// Uint32 returns value as uint32.
func (q *Quantity) Uint32() uint32 {
	var u32 uint32

	for i, b := range q.Value() {
		u := uint32(b) & 0x7f
		j := len(q.value) - i - 1
		shift := (uint32(j * 8)) - uint32(j)
		u = u << shift
		u32 += u
	}

	return u32
}

// Value returns value as byte slice.
func (q *Quantity) Value() []byte {
	if q.value == nil {
		q.value = make([]byte, 1)
	}

	return q.value
}

// SetUint32 reads value as uint32 and sets the value of Quantity.
func (q *Quantity) SetUint32(value uint32) error {
	return nil
}

// SetValue reads value as byte slice and sets the value of Quantity.
func (q *Quantity) SetValue(value []byte) error {
	if len(value) > 4 {
		return fmt.Errorf("midi: maximum length of byte slice is 4, but len(value) = %v", len(value))
	}
	q.value = value

	return nil
}

// Serialize serializes value of variable length quantity.
func (q *Quantity) Serialize() []byte {
	return q.Value()
}

func parseQuantity(stream []byte) (*Quantity, error) {
	if len(stream) == 0 {
		return nil, fmt.Errorf("midi.parseQuantity: stream is empty")
	}

	var i int
	q := &Quantity{}

	for {
		if i > 3 {
			return nil, fmt.Errorf("midi.parseQuantity: maximum size of variable quantity is 4 bytes")
		}
		if len(stream) < (i + 1) {
			return nil, fmt.Errorf("midi.parseQuantity: missing next byte (stream=%+v)", stream)
		}
		if stream[i] < 0x80 {
			break
		}
		i++
	}

	q.value = make([]byte, i+1)
	copy(q.value, stream)

	return q, nil
}
