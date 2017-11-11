package midi

import "fmt"

// SetTempoEvent corresponds to set tempo event.
type SetTempoEvent struct {
	deltaTime *DeltaTime
	tempo     uint32
}

// DeltaTime returns delta time of set tempo event.
func (e *SetTempoEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of tempo event.
func (e *SetTempoEvent) String() string {
	return fmt.Sprintf("&SetTempoEvent{tempo: %v}", e.tempo)
}

// Serialize serializes set tempo event.
func (e *SetTempoEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, SetTempo)
	bs = append(bs, 0x03)
	bs = append(bs, byte(e.tempo>>16))
	bs = append(bs, byte((0xff00&e.tempo)>>8))
	bs = append(bs, byte(e.tempo&0xff))

	return bs
}

// SetSetTempo sets text.
func (e *SetTempoEvent) SetTempo(tempo uint32) error {
	if tempo > 0x7fffff {
		return fmt.Errorf("midi: maximum tempo is 0x7fffff")
	}
	e.tempo = tempo

	return nil
}

// Tempo returns tempo.
func (e *SetTempoEvent) Tempo() uint32 {
	return e.tempo
}

// NewSetTempoEvent returns SetTempoEvent with the given parameter.
func NewSetTempoEvent(deltaTime *DeltaTime, tempo uint32) (*SetTempoEvent, error) {
	var err error

	event := &SetTempoEvent{}
	event.deltaTime = deltaTime

	err = event.SetTempo(tempo)
	if err != nil {
		return nil, err
	}
	return event, nil
}
