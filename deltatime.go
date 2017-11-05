package midi

import (
	"fmt"
)

type DeltaTime struct {
	value []byte
}

func (d *DeltaTime) serialize() []byte {
	return d.value
}

func parseDeltaTime(stream []byte) (*DeltaTime, error) {
	if len(stream) == 0 {
		return nil, fmt.Errorf("midi.parseDeltaTime: stream is empty")
	}

	var i int
	dt := &DeltaTime{}

	for {
		if i > 3 {
			return nil, fmt.Errorf("midi.parseDeltaTime: maximum size of delta time is 4 bytes")
		}
		if len(stream) < (i + 1) {
			return nil, fmt.Errorf("midi.parseDeltaTime: missing next byte (stream=%+v)", stream)
		}
		if stream[i] < 0x80 {
			break
		}
		i++
	}

	dt.value = make([]byte, i+1)
	copy(dt.value, stream)

	return dt, nil
}
