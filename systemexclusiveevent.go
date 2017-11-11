package midi

import "fmt"

// SystemExclusiveEvent corresponds to system exclusive meta event.
type SystemExclusiveEvent struct {
	deltaTime *DeltaTime
	data      []byte
}

// DeltaTime returns delta time of system exclusive event.
func (e *SystemExclusiveEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of system exclusive event.
func (e *SystemExclusiveEvent) String() string {
	return fmt.Sprintf("&SystemExclusiveEvent{data: %v bytes}", len(e.Data()))
}

// Serialize serializes system exclusive event.
func (e *SystemExclusiveEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, SystemExclusive)
	q := &Quantity{}
	q.SetUint32(uint32(len(e.Data())))
	bs = append(bs, q.Serialize()...)
	bs = append(bs, e.Data()...)

	return bs
}

// SetData sets data.
func (e *SystemExclusiveEvent) SetData(data []byte) error {
	if len(data) > 0xfffffff {
		return fmt.Errorf("midi: maximum size of data is 256 MB")
	}
	e.data = data

	return nil
}

// Data returns data.
func (e *SystemExclusiveEvent) Data() []byte {
	if e.data == nil {
		e.data = []byte{}
	}

	return e.data
}

// NewSystemExclusiveEvent returns SystemExclusiveEvent with the given parameter.
func NewSystemExclusiveEvent(deltaTime *DeltaTime, data []byte) (*SystemExclusiveEvent, error) {
	var err error

	event := &SystemExclusiveEvent{}
	event.deltaTime = deltaTime

	err = event.SetData(data)
	if err != nil {
		return nil, err
	}
	return event, nil
}
