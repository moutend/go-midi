package midi

import (
	"fmt"
)

type Quantity struct {
	value []byte
}

func (q *Quantity) Int() int {
	return int(q.Uint32())
}

func (q *Quantity) Uint32() uint32 {
	var u32 uint32

	for i, b := range q.value {
		u := uint32(b) & 0x7f
		j := len(q.value) - i - 1
		shift := (uint32(j * 8)) - uint32(j)
		u = u << shift
		u32 += u
	}

	return u32
}

func (q *Quantity) Value() int {
	return int(q.value[0])
}

func (q *Quantity) SetRawValue(value []byte) {
	q.value = value
}

func (q *Quantity) serialize() []byte {
	return q.value
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
