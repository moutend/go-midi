package midi

import "fmt"

// SequencerSpecificEvent corresponds to sequencer specific event.
type SequencerSpecificEvent struct {
	deltaTime *DeltaTime
	data      []byte
}

// DeltaTime returns delta time.
func (e *SequencerSpecificEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of sequencer specific event.
func (e *SequencerSpecificEvent) String() string {
	return fmt.Sprintf("&SequencerSpecificEvent{data: %v bytes}", len(e.Data()))
}

// Serialize serializes sequencer specific event.
func (e *SequencerSpecificEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, SequencerSpecific)

	q := &Quantity{}
	q.SetUint32(uint32(len(e.Data())))
	bs = append(bs, q.Value()...)
	bs = append(bs, e.Data()...)

	return bs
}

// SetData sets data.
func (e *SequencerSpecificEvent) SetData(data []byte) error {
	if len(data) > 0xfffffff {
		return fmt.Errorf("midi: maximum length of data is 256 MB")
	}
	e.data = data

	return nil
}

// Data returns data.
func (e *SequencerSpecificEvent) Data() []byte {
	if e.data == nil {
		e.data = []byte{}
	}
	return e.data
}

// NewSequencerSpecificEvent returns SequencerSpecificEvent with the given parameter.
func NewSequencerSpecificEvent(deltaTime *DeltaTime, data []byte) (*SequencerSpecificEvent, error) {
	var err error

	event := &SequencerSpecificEvent{}
	event.deltaTime = deltaTime

	err = event.SetData(data)
	if err != nil {
		return nil, err
	}
	return event, nil
}
