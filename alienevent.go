package midi

import "fmt"

// Alien represents unknown meta event.
type AlienEvent struct {
	deltaTime     *DeltaTime
	metaEventType uint8
	data          []byte
}

// DeltaTime returns delta time.
func (e *AlienEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of alien event.
func (e *AlienEvent) String() string {
	return fmt.Sprintf("&AlienEvent{metaEventType: 0x%x, data: %v bytes}", e.metaEventType, len(e.Data()))
}

// Serialize serializes alien event.
func (e *AlienEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, e.metaEventType)

	q := &Quantity{}
	q.SetUint32(uint32(len(e.Data())))
	bs = append(bs, q.Value()...)
	bs = append(bs, e.Data()...)

	return bs
}

// SetMetaEventType sets meta event type.
func (e *AlienEvent) SetMetaEventType(metaEventType uint8) error {
	e.metaEventType = metaEventType

	return nil
}

// MetaEventType returns meta event type.
func (e *AlienEvent) MetaEventType() uint8 {
	return e.metaEventType
}

// SetData sets data.
func (e *AlienEvent) SetData(data []byte) error {
	if len(data) > 0xfffffff {
		return fmt.Errorf("midi: maximum length of data is 256 MB")
	}
	e.data = data

	return nil
}

// Data returns data.
func (e *AlienEvent) Data() []byte {
	if e.data == nil {
		e.data = []byte{}
	}
	return e.data
}

// NewAlienEvent returns AlienEvent with the given parameter.
func NewAlienEvent(deltaTime *DeltaTime, metaEventType uint8, data []byte) (*AlienEvent, error) {
	var err error

	event := &AlienEvent{}
	event.deltaTime = deltaTime

	err = event.SetMetaEventType(metaEventType)
	if err != nil {
		return nil, err
	}
	err = event.SetData(data)
	if err != nil {
		return nil, err
	}
	return event, nil
}
